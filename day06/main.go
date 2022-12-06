package main

import (
	"fmt"

	"github.com/justingodden/advent-of-code-2022/utils"
)

func main() {
	fmt.Println(part1())
	fmt.Println(part2())
}

func part1() int {
	input := utils.ReadFile("input.txt")
	for i := 3; i < len(input); i++ {
		chars := map[byte]bool{
			input[i-3]: true,
			input[i-2]: true,
			input[i-1]: true,
			input[i]:   true,
		}
		if len(chars) == 4 {
			return i + 1
		}
	}
	return 0
}

func part2() int {
	input := utils.ReadFile("input.txt")
	for i := 13; i < len(input); i++ {
		chars := map[byte]bool{
			input[i-13]: true,
			input[i-12]: true,
			input[i-11]: true,
			input[i-10]: true,
			input[i-9]:  true,
			input[i-8]:  true,
			input[i-7]:  true,
			input[i-6]:  true,
			input[i-5]:  true,
			input[i-4]:  true,
			input[i-3]:  true,
			input[i-2]:  true,
			input[i-1]:  true,
			input[i]:    true,
		}
		if len(chars) == 14 {
			return i + 1
		}
	}
	return 0
}
