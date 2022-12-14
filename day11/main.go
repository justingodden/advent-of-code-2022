package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/justingodden/advent-of-code-2022/utils"
)

type monkey struct {
	items        []int
	operation    func(x int) int
	divisor      int
	isDivisible  func(x int) bool
	toMonkey     map[bool]int
	numInspected int
}

func newMonkey(instruction string) monkey {
	lines := strings.Split(instruction, "\n")

	itemsStr := strings.Split(strings.Split(lines[1], ": ")[1], ", ")
	items := make([]int, len(itemsStr))
	for i, itmStr := range itemsStr {
		itm, _ := strconv.Atoi(itmStr)
		items[i] = itm
	}

	operationStr := strings.Split(strings.Split(lines[2], "= ")[1], " ")
	var operation func(x int) int
	switch {
	case operationStr[2] != "old":
		num, _ := strconv.Atoi(operationStr[2])
		switch operationStr[1] {
		case "+":
			operation = func(x int) int { return x + num }
		case "*":
			operation = func(x int) int { return x * num }
		}

	default:
		switch operationStr[1] {
		case "+":
			operation = func(x int) int { return x + x }
		case "*":
			operation = func(x int) int { return x * x }
		}
	}

	testNum, _ := strconv.Atoi(strings.Split(lines[3], "by ")[1])
	isDivisible := func(x int) bool { return x%testNum == 0 }

	t, _ := strconv.Atoi(strings.Split(lines[4], "monkey ")[1])
	f, _ := strconv.Atoi(strings.Split(lines[5], "monkey ")[1])
	toMonkey := map[bool]int{
		true:  t,
		false: f,
	}

	return monkey{
		items:       items,
		operation:   operation,
		isDivisible: isDivisible,
		divisor:     testNum,
		toMonkey:    toMonkey,
	}
}

func main() {
	fmt.Println(part1())
	fmt.Println(part2())
}

func part1() int {
	input := utils.ReadFile("input.txt")
	instructions := strings.Split(input, "\n\n")

	monkies := make([]monkey, len(instructions))

	for i, instruction := range instructions {
		monkies[i] = newMonkey(instruction)
	}

	for round := 1; round <= 20; round++ {
		for i, mnky := range monkies {
			for _, itm := range mnky.items {
				itm = mnky.operation(itm)
				itm = (itm / 3)
				n := mnky.toMonkey[mnky.isDivisible(itm)]
				monkies[n].items = append(monkies[n].items, itm)
				monkies[i].numInspected++
			}
			monkies[i].items = []int{}
		}
	}

	inspected := []int{}
	for _, mnky := range monkies {
		inspected = append(inspected, mnky.numInspected)
	}
	sort.Ints(inspected)

	return inspected[len(inspected)-1] * inspected[len(inspected)-2]
}

func part2() int {
	input := utils.ReadFile("input.txt")
	instructions := strings.Split(input, "\n\n")

	monkies := make([]monkey, len(instructions))

	for i, instruction := range instructions {
		monkies[i] = newMonkey(instruction)
	}

	prod := 1
	for _, mnky := range monkies {
		prod *= mnky.divisor
	}

	for round := 1; round <= 10000; round++ {
		for i, mnky := range monkies {
			for _, itm := range mnky.items {
				itm = mnky.operation(itm)
				itm = (itm % prod)
				n := mnky.toMonkey[mnky.isDivisible(itm)]
				monkies[n].items = append(monkies[n].items, itm)
				monkies[i].numInspected++
			}
			monkies[i].items = []int{}
		}
	}

	inspected := []int{}
	for _, mnky := range monkies {
		inspected = append(inspected, mnky.numInspected)
	}
	sort.Ints(inspected)

	return inspected[len(inspected)-1] * inspected[len(inspected)-2]
}
