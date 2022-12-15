package main

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/justingodden/advent-of-code-2022/utils"
)

func main() {
	fmt.Println(part1())
	fmt.Println(part2())
}

func part1() int {
	input := utils.ReadFile("input.txt")
	lines := strings.Split(input, "\n\n")

	var sum int

	for i, s := range lines {
		s := strings.Split(s, "\n")
		var a, b any
		json.Unmarshal([]byte(s[0]), &a)
		json.Unmarshal([]byte(s[1]), &b)

		if cmp(a, b) <= 0 {
			sum += i + 1
		}
	}

	return sum
}

func part2() int {
	input := utils.ReadFile("input.txt")
	lines := strings.Split(input, "\n\n")

	packets := []any{}

	for _, s := range lines {
		s := strings.Split(s, "\n")
		var a, b any
		json.Unmarshal([]byte(s[0]), &a)
		json.Unmarshal([]byte(s[1]), &b)
		packets = append(packets, a, b)
	}

	packets = append(packets, []any{[]any{2.}}, []any{[]any{6.}})
	sort.Slice(packets, func(i, j int) bool { return cmp(packets[i], packets[j]) < 0 })

	prod := 1

	for i, p := range packets {
		if fmt.Sprint(p) == "[[2]]" || fmt.Sprint(p) == "[[6]]" {
			prod *= i + 1
		}
	}

	return prod
}

func cmp(a, b any) int {
	as, aok := a.([]any)
	bs, bok := b.([]any)

	switch {
	case !aok && !bok:
		return int(a.(float64) - b.(float64))
	case !aok:
		as = []any{a}
	case !bok:
		bs = []any{b}
	}

	for i := 0; i < len(as) && i < len(bs); i++ {
		if c := cmp(as[i], bs[i]); c != 0 {
			return c
		}
	}

	return len(as) - len(bs)
}
