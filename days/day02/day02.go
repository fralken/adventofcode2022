package day02

import (
	"aoc2022/utils"
	"strings"
)

func FirstStar() {
	utils.Star(2, 1, "total score", firstStar)
}

func SecondStar() {
	utils.Star(2, 2, "total score", secondStar)
}

func firstStar(content string) (s int) {
	lines := strings.Split(content, "\n")
	for _, l := range lines {
		s += score1(l)
	}
	return
}

func secondStar(content string) (s int) {
	lines := strings.Split(content, "\n")
	for _, l := range lines {
		s += score2(l)
	}
	return
}

// A < B < C < A
// X=A Y=B Z=C
// A=1 B=2 C=3
// lose=0 draw=3 win=6
// we play X Y Z
func score1(play string) (s int) {
	switch play {
	case "A X":
		s = 1 + 3
	case "A Y":
		s = 2 + 6
	case "A Z":
		s = 3
	case "B X":
		s = 1
	case "B Y":
		s = 2 + 3
	case "B Z":
		s = 3 + 6
	case "C X":
		s = 1 + 6
	case "C Y":
		s = 2
	case "C Z":
		s = 3 + 3
	}
	return
}

// A < B < C < A
// X=lose Y=draw Z=win
// A=1 B=2 C=3
// lose=0 draw=3 win=6
// we play X Y Z
func score2(play string) (s int) {
	switch play {
	case "A X":
		s = 3
	case "A Y":
		s = 1 + 3
	case "A Z":
		s = 2 + 6
	case "B X":
		s = 1
	case "B Y":
		s = 2 + 3
	case "B Z":
		s = 3 + 6
	case "C X":
		s = 2
	case "C Y":
		s = 3 + 3
	case "C Z":
		s = 1 + 6
	}
	return
}