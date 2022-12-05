package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/justingodden/advent-of-code-2022/utils"
)

func main() {
	part1()
	part2()
}

func part1() {
	input := utils.ReadFile("input.txt")
	stacksInput := strings.Split(input, "\n\n")[0]
	commandsInput := strings.Split(input, "\n\n")[1]

	stackMap := map[int][]rune{}
	stacksInputLines := strings.Split(stacksInput, "\n")
	stacksInputLines = stacksInputLines[:len(stacksInputLines)-1]

	for i := len(stacksInputLines) - 1; i >= 0; i-- {
		pos := 1
		for j, ch := range stacksInputLines[i] {
			if (j-1)%4 == 0 {
				if unicode.IsLetter(ch) {
					stackMap[pos] = append(stackMap[pos], ch)
				}
				pos++
			}
		}
	}

	commandsInputLines := strings.Split(commandsInput, "\n")

	for _, cmd := range commandsInputLines {
		cmdArr := strings.Split(cmd, " ")
		n, _ := strconv.Atoi(cmdArr[1])
		from, _ := strconv.Atoi(cmdArr[3])
		to, _ := strconv.Atoi(cmdArr[5])

		for i := 0; i < n; i++ {
			l := len(stackMap[from])
			val := stackMap[from][l-1]
			stackMap[from] = stackMap[from][:l-1]
			stackMap[to] = append(stackMap[to], val)
		}
	}

	for i := 1; i <= 9; i++ {
		l := len(stackMap[i])
		fmt.Printf("%c", stackMap[i][l-1])
	}
	fmt.Println()
}

func part2() {
	input := utils.ReadFile("input.txt")
	stacksInput := strings.Split(input, "\n\n")[0]
	commandsInput := strings.Split(input, "\n\n")[1]

	stackMap := map[int][]rune{}
	stacksInputLines := strings.Split(stacksInput, "\n")
	stacksInputLines = stacksInputLines[:len(stacksInputLines)-1]

	for i := len(stacksInputLines) - 1; i >= 0; i-- {
		pos := 1
		for j, ch := range stacksInputLines[i] {
			if (j-1)%4 == 0 {
				if unicode.IsLetter(ch) {
					stackMap[pos] = append(stackMap[pos], ch)
				}
				pos++
			}
		}
	}

	commandsInputLines := strings.Split(commandsInput, "\n")

	for _, cmd := range commandsInputLines {
		cmdArr := strings.Split(cmd, " ")
		n, _ := strconv.Atoi(cmdArr[1])
		from, _ := strconv.Atoi(cmdArr[3])
		to, _ := strconv.Atoi(cmdArr[5])

		l := len(stackMap[from])

		crates := stackMap[from][l-n:]
		stackMap[from] = stackMap[from][:l-n]
		stackMap[to] = append(stackMap[to], crates...)
	}

	for i := 1; i <= 9; i++ {
		l := len(stackMap[i])
		fmt.Printf("%c", stackMap[i][l-1])
	}
	fmt.Println()
}
