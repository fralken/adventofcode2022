package day05

import (
	"aoc2022/utils"
	"fmt"
	"strings"
)

func FirstStar() {
	utils.Star(5, 1, "crates on top of each stack", firstStar)
}

func SecondStar() {
	utils.Star(5, 2, "crates on top of each stack", secondStar)
}

func firstStar(content string) string {
	stacks, moves := parseStacksMoves(content)
	stacks.move1(&moves)
	return stacks.top()
}

func secondStar(content string) string {
	stacks, moves := parseStacksMoves(content)
	stacks.move2(&moves)
	return stacks.top()
}

type move struct {
	q, f, t int
}

type stacks [][]byte
type moves []move

func parseStacksMoves(content string) (stacks, moves) {
	blocksPart := strings.Split(content, "\n\n")
	stacksPart := strings.Split(blocksPart[0], "\n")
	moves := strings.Split(blocksPart[1], "\n")
	stacksCount := len(strings.Split(stacksPart[len(stacksPart)-1], "   "))
	stackArray := make(stacks, stacksCount)
	for i := len(stacksPart)-2; i >= 0; i-- {
		l := stacksPart[i]
		for j := 0; j < stacksCount; j++ {
			if len(l) > 1 && l[1] != ' ' {
				stackArray[j] = append(stackArray[j], l[1])
			}
			if len(l) > 3 { l = l[4:] }
		}
	}
	movesArray := make([]move, len(moves))
	for l, line := range moves {
		var m move
		fmt.Sscanf(line, "move %d from %d to %d", &m.q, &m.f, &m.t)
		m.f -= 1
		m.t -= 1
		movesArray[l] = m
	}
	return stackArray, movesArray
}

func (s stacks) move1(moves *moves) {
	for _, m := range *moves {
		for i := 0; i < m.q; i++ {
			s[m.t] = append(s[m.t], s[m.f][len(s[m.f])-1])
			s[m.f] = s[m.f][:len(s[m.f])-1]
		}
	}
}

func (s stacks) move2(moves *moves) {
	for _, m := range *moves {
		s[m.t] = append(s[m.t], s[m.f][len(s[m.f])-m.q:]...)
		s[m.f] = s[m.f][:len(s[m.f])-m.q]
	}
}

func (ss stacks) top() string {
	var top []byte
	for _, s := range ss {
		top = append(top, s[len(s)-1])
	}
	return string(top)
}