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

func part1() int {
	input := utils.ReadFile("input.txt")
	allElves := strings.Split(input, "\n\n")

	var maxCals int

	for _, elf := range allElves {
		curElf := strings.Split(elf, "\n")
		var curElfSum int
		for _, str := range curElf {
			cals, err := strconv.Atoi(str)
			if err == nil {
				curElfSum += cals
			}
		}
		if curElfSum >= maxCals {
			maxCals = curElfSum
		}
	}
	return maxCals
}

func part2() int {
	input := utils.ReadFile("input.txt")
	allElves := strings.Split(input, "\n\n")

	var totalCals = []int{0}

	for _, elf := range allElves {
		curElf := strings.Split(elf, "\n")
		var curElfSum int
		for _, str := range curElf {
			cals, err := strconv.Atoi(str)
			if err == nil {
				curElfSum += cals
			}

			totalCals = append(totalCals, curElfSum)
		}
	}
	sort.Ints(totalCals)
	return totalCals[len(totalCals)-1] + totalCals[len(totalCals)-2] + totalCals[len(totalCals)-3]
}
