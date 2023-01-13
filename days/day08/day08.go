package day08

import (
	"aoc2022/utils"
	"strings"
)

func FirstStar() {
	utils.Star(8, 1, "trees visible from outside", firstStar)
}

func SecondStar() {
	utils.Star(8, 2, "highest scenic score", secondStar)
}

func firstStar(content string) int {
	grid := makeGrid(content)
	visible := 2 * (len(grid[0]) + len(grid) - 2)
	for r := 1; r < len(grid) - 1; r++ {
		for c := 1; c < len(grid[0]) - 1; c++ {
			visible += grid.isVisible(r, c)
		}
	}
	return visible
}

func secondStar(content string) int {
	grid := makeGrid(content)
	max := 0
	for r := 1; r < len(grid) - 1; r++ {
		for c := 1; c < len(grid[0]) - 1; c++ {
			s := grid.score(r, c)
			if s > max { max = s }
		}
	}
	return max
}

type grid []string

func makeGrid(content string) grid {
	return strings.Split(content, "\n")
}

func (g grid) isVisible(r, c int) int {
	v := g[r][c]
	count := 0
	for i := 0; i < c; i++ {
		if g[r][i] >= v { 
			count++
			break
		}
	}
	for i := c + 1; i < len(g[r]); i++ {
		if g[r][i] >= v { 
			count++
			break
		}
	}
	for i := 0; i < r; i++ {
		if g[i][c] >= v { 
			count++
			break
		}
	}
	for i := r + 1; i < len(g); i++ {
		if g[i][c] >= v { 
			count++
			break
		}
	}
	if count == 4 { return 0 } else { return 1 }
}

func (g grid) score(r, c int) int {
	v := g[r][c]
	left := 0
	for i := c - 1 ; i >= 0; i-- {
		left++
		if g[r][i] >= v {	break }
	}
	right := 0
	for i := c + 1; i < len(g[r]); i++ {
		right++
		if g[r][i] >= v {	break }
	}
	top := 0
	for i := r - 1; i >= 0; i-- {
		top++
		if g[i][c] >= v {	break }
	}
	bottom := 0
	for i := r + 1; i < len(g); i++ {
		bottom++
		if g[i][c] >= v {	break }
	}
	return left * right * top * bottom
}