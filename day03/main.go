package main

import (
	"fmt"
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
		l := len(line)

		first := line[0 : l/2]
		second := line[l/2:]

		existsLeft := map[rune]bool{}
		intersection := map[rune]bool{}

		for _, ch := range first {
			existsLeft[ch] = true
		}

		for _, ch := range second {
			intersection[ch] = existsLeft[ch]
		}

		for k, v := range intersection {
			if v {
				sum += priority(k)
			}
		}
	}

	return sum
}

func part2() int {
	input := utils.ReadFile("input.txt")
	lines := strings.Split(input, "\n")

	var sum int
	var existsFirst map[rune]bool
	var existsSecond map[rune]bool
	var intersection map[rune]bool

	for i, line := range lines {
		switch {
		case i%3 == 0:
			existsFirst = map[rune]bool{}
			for _, ch := range line {
				existsFirst[ch] = true
			}

		case i%3 == 1:
			existsSecond = map[rune]bool{}
			for _, ch := range line {
				existsSecond[ch] = true
			}

		case i%3 == 2:
			intersection = map[rune]bool{}
			for _, ch := range line {
				intersection[ch] = existsFirst[ch] && existsSecond[ch]
			}

			for k, v := range intersection {
				if v {
					sum += priority(k)
				}
			}
		}
	}

	return sum
}

func priority(ch rune) int {
	if ch > 90 {
		return int(ch) - 96
	}
	return int(ch) - 38
}
