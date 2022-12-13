package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/justingodden/advent-of-code-2022/utils"
)

func main() {
	fmt.Println(part1())
	part2()
}

func part1() int {
	input := utils.ReadFile("input.txt")
	cmds := strings.Split(input, "\n")
	X, cycle, idx, total := 1, 1, 0, 0

	for idx < len(cmds) {
		cmd := strings.Split(cmds[idx], " ")
		switch cmd[0] {
		case "noop":
			total += checkCycle(cycle, X)
			cycle++
			idx++

		default:
			total += checkCycle(cycle, X)
			n, _ := strconv.Atoi(cmd[1])
			cycle++
			total += checkCycle(cycle, X)
			cycle++
			X += n
			idx++
		}
	}

	return total
}

func part2() {
	input := utils.ReadFile("input.txt")
	cmds := strings.Split(input, "\n")
	X, cycle, idx := 1, 0, 0

	for idx < len(cmds) {
		cmd := strings.Split(cmds[idx], " ")
		switch cmd[0] {
		case "noop":
			if (cycle%40) == X || (cycle%40) == X+1 || (cycle%40) == X-1 {
				fmt.Print("🔴")
			} else {
				fmt.Print("⚫")
			}
			cycle++
			idx++

			if (cycle % 40) == 0 {
				fmt.Println()
			}

		default:
			if (cycle%40) == X || (cycle%40) == X+1 || (cycle%40) == X-1 {
				fmt.Print("🔴")
			} else {
				fmt.Print("⚫")
			}
			n, _ := strconv.Atoi(cmd[1])
			cycle++
			if (cycle % 40) == 0 {
				fmt.Println()
			}

			if (cycle%40) == X || (cycle%40) == X+1 || (cycle%40) == X-1 {
				fmt.Print("🔴")
			} else {
				fmt.Print("⚫")
			}
			cycle++
			X += n
			idx++
			if (cycle % 40) == 0 {
				fmt.Println()
			}
		}
	}
}

func checkCycle(cycle int, X int) int {
	switch cycle {
	case 20:
		return 20 * X
	case 60:
		return 60 * X
	case 100:
		return 100 * X
	case 140:
		return 140 * X
	case 180:
		return 180 * X
	case 220:
		return 220 * X
	default:
		return 0
	}
}
