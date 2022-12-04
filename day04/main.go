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
	lines := strings.Split(input, "\n")

	var sum int

	for _, line := range lines {
		left := strings.Split(line, ",")[0]
		right := strings.Split(line, ",")[1]
		leftMin, _ := strconv.Atoi(strings.Split(left, "-")[0])
		leftMax, _ := strconv.Atoi(strings.Split(left, "-")[1])
		rightMin, _ := strconv.Atoi(strings.Split(right, "-")[0])
		rightMax, _ := strconv.Atoi(strings.Split(right, "-")[1])

		if (leftMin <= rightMin && rightMax <= leftMax) || (rightMin <= leftMin && leftMax <= rightMax) {
			sum++
		}
	}

	return sum
}

func part2() int {
	input := utils.ReadFile("input.txt")
	lines := strings.Split(input, "\n")

	var sum int

	for _, line := range lines {
		left := strings.Split(line, ",")[0]
		right := strings.Split(line, ",")[1]
		leftMin, _ := strconv.Atoi(strings.Split(left, "-")[0])
		leftMax, _ := strconv.Atoi(strings.Split(left, "-")[1])
		rightMin, _ := strconv.Atoi(strings.Split(right, "-")[0])
		rightMax, _ := strconv.Atoi(strings.Split(right, "-")[1])

		if (leftMin <= rightMin && rightMin <= leftMax) ||
			(leftMin <= rightMax && rightMax <= leftMax) ||
			(rightMin <= leftMin && leftMin <= rightMax) ||
			(rightMin <= leftMax && leftMax <= rightMax) {
			sum++
		}
	}

	return sum
}
