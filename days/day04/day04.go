package day04

import (
	"aoc2022/utils"
	"fmt"
	"strings"
)

func FirstStar() {
	utils.Star(4, 1, "pairs containing the other", firstStar)
}

func SecondStar() {
	utils.Star(4, 2, "pairs overlapping", secondStar)
}

func firstStar(content string) int {
	pairs := parsePairs(content)
	count := pairs.countContains()
	return count
}

func secondStar(content string) int {
	pairs := parsePairs(content)
	count := pairs.countOverlaps()
	return count
}

type ranges struct {
	a, b, c, d int
}

type pairs []ranges

func parsePairs(content string) pairs {
	lines := strings.Split(content, "\n")
	pairs := pairs{}
	for _, line := range lines {
		var p ranges
		fmt.Sscanf(line, "%d-%d,%d-%d", &p.a, &p.b, &p.c, &p.d)
		pairs = append(pairs, p)
	}
	return pairs
}

func (ps pairs) countContains() (count int) {
	for _, p := range ps {
		if p.a <= p.c && p.b >= p.d ||
			p.c <= p.a && p.d >= p.b {
			count++
		}
	}
	return
}

func (ps pairs) countOverlaps() (count int) {
	for _, p := range ps {
		if p.a <= p.c && p.c <= p.b ||
			p.c <= p.a && p.a <= p.d {
			count++
		}
	}
	return
}