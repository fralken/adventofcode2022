package day10

import (
	"aoc2022/utils"
	"fmt"
	"strings"
)

func FirstStar() {
	utils.Star(10, 1, "sum of six signal strengths", firstStar)
}

func SecondStar() {
	utils.Star(10, 2, "crt display letters", secondStar)
}

func firstStar(content string) int {
	lines := strings.Split(content, "\n")
	cycle := 0
	strength := 0
	check := 20
	x := 1
	for _, line := range lines {
		if line == "noop" {
			cycle++
			strength += checkCycle(cycle, x, &check)
		} else {
			cycle++
			strength += checkCycle(cycle, x, &check)
			cycle++
			strength += checkCycle(cycle, x, &check)
			x += utils.StringToInt(line[5:])
		}
	}
	return strength
}

func secondStar(content string) string {
	lines := strings.Split(content, "\n")
	cycle := 0
	x := 1
	var crt [240]byte
	for _, line := range lines {
		if line == "noop" {
			cycle++
			display(&crt, cycle, x)
		} else {
			cycle++
			display(&crt, cycle, x)
			cycle++
			display(&crt, cycle, x)
			x += utils.StringToInt(line[5:])
		}
	}
	image := "\n"
	for i := 0; i < 240; i += 40 {
		image = fmt.Sprintf("%s%s\n", image, string(crt[i:i+40]))
	}
	return image
}

func checkCycle(cycle int, x int, check *int) (strength int) {
	if cycle == *check && *check <= 220 {
		strength = cycle * x
		*check += 40
	}
	return
}

func display(crt *[240]byte, cycle, x int) {
	c := cycle - 1
	q := c % 40
	if q >= x - 1 && q <= x + 1 {
		(*crt)[c] = '#'
	} else {
		(*crt)[c] = ' '
	}
}