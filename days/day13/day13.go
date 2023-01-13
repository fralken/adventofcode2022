package day13

import (
	"aoc2022/utils"
	"fmt"
	"strings"
	"sort"
)

func FirstStar() {
	utils.Star(13, 1, "sum of indices of ordered pairs", firstStar)
}

func SecondStar() {
	utils.Star(13, 2, "decoder key of distressed signal", secondStar)
}

func firstStar(content string) int {
	pairs := parseSignals(content)
	sum := 0
	for i, pair := range pairs {
		if compare(pair[0], pair[1]) <= 0 {
			sum += (i + 1)
		}
	}
	return sum
}

func secondStar(content string) int {
	content = fmt.Sprintf("%s\n\n[[2]]\n[[6]]", content)
	pairs := parseSignals(content)
	list := []any{}
	for _, pair := range pairs {
		list = append(list, pair...)
	}
	sort.Slice(list, func(i, j int) bool {
		return compare(list[i], list[j]) < 0
	})
	val := 1
	for i, l := range list {
		s := fmt.Sprint(l)
		if s == "[[2]]" || s == "[[6]]" {
			val *= i + 1
		}
	}
	return val
}

func parseSignals(content string) (pairs [][]any) {
	parts := strings.Split(content, "\n\n")
	for _, part := range parts {
		lines := strings.Split(part, "\n")
		pairs = append(pairs, []any{ parseSignal(lines[0]), parseSignal(lines[1]) })
	}
	return
}

func parseSignal(content string) any {
	root := []any{}
	navigate(&root, content, 0)
	return root[0]
}

func navigate(n *([]any), content string, i int) int {
	for i < len(content) {
		if content[i] == '[' {
			next := []any{}
			i = navigate(&next, content, i+1)
			*n = append(*n, next)
		} else if content[i] >= '0' && content[i] <= '9' {
			j := i+1
			for content[j] >= '0' && content[j] <= '9' { j++ }
			*n = append(*n, any(utils.StringToInt(content[i:j])))
			i = j
		} else if content[i] == ',' {
			i++
		} else if content[i] == ']' {
			i++
			break 
		}
	}
	return i
}

func compare(left any, right any) int {
	lInt, lIsInt := left.(int)
	rInt, rIsInt := right.(int)
	if lIsInt && rIsInt {
		if lInt < rInt { return -1 }
		if lInt > rInt { return 1 }
	} else if lIsInt {
		return compare([]any{lInt}, right)
	} else if rIsInt {
		return compare(left, []any{rInt})
	} else {
		lList, _ := left.([]any)
		rList, _ := right.([]any)
		size := utils.Max(len(lList), len(rList))
		for i := 0; i < size; i++ {
			if i >= len(lList) { return -1 }
			if i >= len(rList) { return 1 }
			if ord := compare(lList[i], rList[i]); ord != 0 {
				return ord
			}
		}
	}
	return 0
}
