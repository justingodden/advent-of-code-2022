package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/justingodden/advent-of-code-2022/utils"
)

func main() {
	fmt.Println(part1())
	fmt.Println(part2())
}

type node struct {
	isDir    bool
	name     string
	size     int
	children map[string]*node
	parent   *node
}

func newDir(name string, parent *node) *node {
	return &node{
		isDir:    true,
		name:     name,
		children: map[string]*node{},
		parent:   parent,
	}
}

func newFile(name string, size int, parent *node) *node {
	return &node{
		isDir:  false,
		name:   name,
		size:   size,
		parent: parent,
	}
}

func (n *node) getSize() int {
	if !n.isDir {
		return n.size
	}

	if n.size > 0 {
		return n.size
	}

	for _, child := range n.children {
		n.size += child.getSize()
	}

	return n.size
}

func constuctFS(input string) *node {
	fs := newDir("/", nil)
	current := fs

	lines := strings.Split(input, "\n")
	for _, line := range lines[1:] {
		tokens := strings.Split(line, " ")
		switch tokens[0] {
		case "$":
			switch tokens[1] {
			case "cd":
				switch tokens[2] {
				case "..":
					current = current.parent
				default:
					current = current.children[tokens[2]]
				}
			case "ls":
				continue
			}
		case "dir":
			current.children[tokens[1]] = newDir(tokens[1], current)

		default:
			size, _ := strconv.Atoi(tokens[0])
			current.children[tokens[1]] = newFile(tokens[1], size, current)
		}
	}

	return fs
}

func traverse(node *node, lim int, total int) int {
	if node.children != nil {
		for _, child := range node.children {
			if child.isDir {
				total = traverse(child, lim, total)
			}
		}
	}
	if node.isDir {
		if node.size < lim {
			total += node.size
		}
	}

	return total
}

func getDirSizes(node *node, sizes []int) []int {
	sizes = append(sizes, node.size)

	for _, child := range node.children {
		if child.isDir {
			sizes = getDirSizes(child, sizes)
		}
	}

	return sizes
}

func part1() int {
	input := utils.ReadFile("input.txt")
	fs := constuctFS(input)
	_ = fs.getSize()
	return traverse(fs, 100_000, 0)
}

func part2() int {
	input := utils.ReadFile("input.txt")
	fs := constuctFS(input)
	_ = fs.getSize()

	unused := 70_000_000 - fs.size
	needed := 30_000_000 - unused

	sizes := make([]int, 0)
	sizes = getDirSizes(fs, sizes)

	sort.Ints(sizes)

	for _, i := range sizes {
		if i > needed {
			return i
		}
	}
	return 0
}
