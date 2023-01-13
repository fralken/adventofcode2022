package day01

import (
	"aoc2022/utils"
	"strings"
)

func FirstStar() {
	utils.Star(1, 1, "max total calories", firstStar)
}

func SecondStar() {
	utils.Star(1, 2, "sum top 3 calories", secondStar)
}

func firstStar(content string) int {
	elves := extractElves(content)
	return elves.max()
}

func secondStar(content string) int {
	elves := extractElves(content)
	return elves.sumTop(3)
}

type elves []int

func extractElves(content string) elves {
	elvesList := strings.Split(content, "\n\n")
	elves := make(elves, len(elvesList))
	for e, elf := range elvesList {
		calories := strings.Split(elf, "\n")
		sum := 0
		for _, c := range calories {
			sum += utils.StringToInt(c)
		}
		elves[e] = sum
	}
	return elves
}

func (es elves) max() int {
	max := 0
	for _, e := range es {
		if e > max { max = e }
	}
	return max
}

func (es elves) sumTop(n int) int {
	top := make([]int, 0)
	for _, e := range es {
		if len(top) < n {
			top = append(top, e)
		} else {
			for i, t := range top {
				if e > t {
					top[i] = e
					break
				}
			}
		}
	}
	sum := 0
	for _, t := range top {
		sum += t
	}
	return sum
}
