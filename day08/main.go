package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/justingodden/advent-of-code-2022/utils"
)

func constructGrid(input string) [][]int {
	lines := strings.Split(input, "\n")
	grid := make([][]int, len(lines))

	for j := 0; j < len(lines); j++ {
		row := make([]int, len(lines[0]))

		for i := 0; i < len(lines[0]); i++ {
			n, _ := strconv.Atoi(string(lines[j][i]))
			row[i] = n
		}
		grid[j] = row
	}

	return grid
}

func part1() int {
	input := utils.ReadFile("input.txt")
	grid := constructGrid(input)
	var numVisible int
	var pos int

	for j := 0; j < len(grid); j++ {
	outer:
		for i := 0; i < len(grid[0]); i++ {

			// Left
			pos = i
			leftVis := true
			for pos > 0 {
				if grid[j][i] <= grid[j][pos-1] {
					leftVis = false
					break
				}
				pos--
			}

			if leftVis {
				numVisible++
				continue outer
			}

			// Right
			pos = i
			rightVis := true
			for pos < len(grid[0])-1 {
				if grid[j][i] <= grid[j][pos+1] {
					rightVis = false
					break
				}
				pos++
			}

			if rightVis {
				numVisible++
				continue outer
			}

			// Up
			pos = j
			upVis := true
			for pos > 0 {
				if grid[j][i] <= grid[pos-1][i] {
					upVis = false
					break
				}
				pos--
			}

			if upVis {
				numVisible++
				continue outer
			}

			// Down
			pos = j
			downVis := true
			for pos < len(grid)-1 {
				if grid[j][i] <= grid[pos+1][i] {
					downVis = false
					break
				}
				pos++
			}

			if downVis {
				numVisible++
				continue outer
			}
		}
	}

	return numVisible
}

func part2() int {
	input := utils.ReadFile("input.txt")
	grid := constructGrid(input)
	scores := make([]int, len(grid)*len(grid[0]))
	var pos int

	for j := 0; j < len(grid); j++ {
		for i := 0; i < len(grid[0]); i++ {

			// Left
			pos = i
			leftVis := 0
			for pos > 0 {
				leftVis++
				if grid[j][i] <= grid[j][pos-1] {
					break
				}
				pos--
			}

			// Right
			pos = i
			rightVis := 0
			for pos < len(grid[0])-1 {
				rightVis++
				if grid[j][i] <= grid[j][pos+1] {
					break
				}
				pos++
			}

			// Up
			pos = j
			upVis := 0
			for pos > 0 {
				upVis++
				if grid[j][i] <= grid[pos-1][i] {
					break
				}
				pos--
			}

			// Down
			pos = j
			downVis := 0
			for pos < len(grid)-1 {
				downVis++
				if grid[j][i] <= grid[pos+1][i] {
					break
				}
				pos++
			}

			scores = append(scores, leftVis*rightVis*upVis*downVis)
		}
	}

	sort.Ints(scores)
	return scores[len(scores)-1]
}

func main() {
	fmt.Println(part1())
	fmt.Println(part2())
}
