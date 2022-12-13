package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/justingodden/advent-of-code-2022/utils"
)

func main() {
	fmt.Println(part1())
	fmt.Println(part2())
}

func part1() int {
	input := utils.ReadFile("input.txt")
	instructions := strings.Split(input, "\n")

	visited := make(map[[2]int]bool)
	headI, headJ, tailI, tailJ := 0, 0, 0, 0
	visited[[2]int{tailI, tailJ}] = true

	for _, instruction := range instructions {
		dir, num := parseInstruction(instruction)

		for i := 0; i < num; i++ {
			headI, headJ = moveHead(headI, headJ, dir)
			tailI, tailJ = moveTail(headI, headJ, tailI, tailJ)
			visited[[2]int{tailI, tailJ}] = true
		}
	}

	return len(visited)
}

func part2() int {
	input := utils.ReadFile("input.txt")
	instructions := strings.Split(input, "\n")

	visited := make(map[[2]int]bool)

	var (
		headI, headJ   int
		tailI1, tailJ1 int
		tailI2, tailJ2 int
		tailI3, tailJ3 int
		tailI4, tailJ4 int
		tailI5, tailJ5 int
		tailI6, tailJ6 int
		tailI7, tailJ7 int
		tailI8, tailJ8 int
		tailI9, tailJ9 int
	)

	visited[[2]int{tailI9, tailJ9}] = true

	for _, instruction := range instructions {
		dir, num := parseInstruction(instruction)

		for i := 0; i < num; i++ {
			headI, headJ = moveHead(headI, headJ, dir)
			tailI1, tailJ1 = moveTail(headI, headJ, tailI1, tailJ1)
			tailI2, tailJ2 = moveTail(tailI1, tailJ1, tailI2, tailJ2)
			tailI3, tailJ3 = moveTail(tailI2, tailJ2, tailI3, tailJ3)
			tailI4, tailJ4 = moveTail(tailI3, tailJ3, tailI4, tailJ4)
			tailI5, tailJ5 = moveTail(tailI4, tailJ4, tailI5, tailJ5)
			tailI6, tailJ6 = moveTail(tailI5, tailJ5, tailI6, tailJ6)
			tailI7, tailJ7 = moveTail(tailI6, tailJ6, tailI7, tailJ7)
			tailI8, tailJ8 = moveTail(tailI7, tailJ7, tailI8, tailJ8)
			tailI9, tailJ9 = moveTail(tailI8, tailJ8, tailI9, tailJ9)
			visited[[2]int{tailI9, tailJ9}] = true
		}
	}

	return len(visited)
}

func parseInstruction(instruction string) (string, int) {
	dir := strings.Split(instruction, " ")[0]
	num, _ := strconv.Atoi(strings.Split(instruction, " ")[1])
	return dir, num
}

func moveHead(i, j int, dir string) (int, int) {
	switch dir {
	case "L":
		return i - 1, j
	case "U":
		return i, j - 1
	case "R":
		return i + 1, j
	case "D":
		return i, j + 1
	default:
		return i, j
	}
}

func moveTail(headI, headJ, tailI, tailJ int) (int, int) {
	difI := headI - tailI
	difJ := headJ - tailJ

	switch {
	// diagonal
	case abs(difI) == 2 && abs(difJ) == 2:
		return tailI + (difI / 2), tailJ + (difJ / 2)

	// above or below
	case difI == 0 && abs(difJ) == 2:
		return tailI, tailJ + (difJ / 2)

	// left or right
	case abs(difI) == 2 && difJ == 0:
		return tailI + (difI / 2), tailJ

	case abs(difI) == 1 && abs(difJ) == 2:
		return tailI + difI, tailJ + (difJ / 2)

	case abs(difI) == 2 && abs(difJ) == 1:
		return tailI + (difI / 2), tailJ + difJ

	default:
		return tailI, tailJ
	}
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}
