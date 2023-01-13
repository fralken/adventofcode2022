package day09

import (
	"aoc2022/utils"
	"strings"
)

func FirstStar() {
	utils.Star(9, 1, "position visited by tail rope", firstStar)
}

func SecondStar() {
	utils.Star(9, 2, "position visited by tail rope", secondStar)
}

func firstStar(content string) int {
	return moveRope(content, 2)
}

func secondStar(content string) int {
	return moveRope(content, 10)
}

type pos struct {
	x, y int
}

func moveRope(content string, size int) int {
	lines := strings.Split(content, "\n")
	rope := make([]pos, size)
	tail := size - 1
	for i := 0; i < size; i++ {
		rope[i] = pos{ 0, 0 }
	}
	visited := make(map[pos]bool, 0)
	visited[rope[tail]] = true
	for _, line := range lines {
		c := utils.StringToInt(line[2:])
		for i := 0; i < c; i++ {
			switch line[0] {
			case 'R':
				rope[0].x++
			case 'L':
				rope[0].x--
			case 'U':
				rope[0].y--
			case 'D':
				rope[0].y++
			}
			for i := 0; i < tail; i++ {
				update(&rope[i], &rope[i+1])
			}
			visited[rope[tail]] = true
		}
	}
	return len(visited)
}

func update(head *pos, tail *pos) {
	dx := utils.Abs(head.x - tail.x)
	dy := utils.Abs(head.y - tail.y)
	if dx > 1 || dx == 1 && dy > 1 {
		if head.x > tail.x { tail.x++ } else { tail.x-- } 
	}
	if dy > 1 || dy == 1 && dx > 1 {
		if head.y > tail.y { tail.y++ } else { tail.y-- } 
	}
}
