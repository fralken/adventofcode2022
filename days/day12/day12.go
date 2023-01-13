package day12

import (
	"aoc2022/utils"
	"container/heap"
	"strings"
)

func FirstStar() {
	utils.Star(12, 1, "shortest path from start to end", firstStar)
}

func SecondStar() {
	utils.Star(12, 2, "minimum shortest path", secondStar)
}

func firstStar(content string) int {
	grid, start, end := parseMap(content)
	return grid.findPath(start, end)
}

func secondStar(content string) int {
	grid, _, end := parseMap(content)
	starts := grid.findStarts()
	min := 0
	for _, start := range starts {
		steps := grid.findPath(start, end) 
		if steps > 0 && (min == 0 || min > steps) {
			min = steps
		}
	}
	return min
}

type coord [2]int

type grid []string

type path struct {
	pos coord
	length int
}

type paths []*path

func parseMap(content string) (grid grid, start coord, end coord) {
	grid = strings.Split(content, "\n")
	for r, row := range grid {
		for c := range row {
			if grid[r][c] == 'S' {
				start = coord{r, c}
				grid[r] = strings.Replace(grid[r], "S", "a", 1)
			} else if grid[r][c] == 'E' {
				end = coord{r, c}
				grid[r] = strings.Replace(grid[r], "E", "z", 1)
			}
		}
	}
	return
}

func (g grid) findPath(start coord, end coord) int {
	h := len(g) - 1
	w := len(g[0]) - 1
	p := paths{}
	heap.Init(&p)
	visited := make(map[coord]bool)
	visited[coord{0, 0}] = true
	neighbours := []coord{{-1, 0}, {0, -1}, {0, 1}, {1, 0}}
	i, j, c := start[0], start[1], 0
	for i != end[0] || j != end[1] {
		for _, n := range neighbours {
			y, x := i+n[0], j+n[1]
			if y >= 0 && y <= h && x >= 0 && x <= w {
				k := coord{y, x}
				if g[y][x] <= g[i][j] + 1 && !visited[k] {
					visited[k] = true
					heap.Push(&p, &path{ k, c + 1 })
				}
			}
		}
		if len(p) > 0 {
			path := heap.Pop(&p).(*path)
			i, j, c = path.pos[0], path.pos[1], path.length
		} else {
			c = -1
			break
		}
	}
	return c
}

func (g grid) findStarts() (starts []coord) {
	for r, row := range g {
		for c := range row {
			if g[r][c] == 'a' {
				starts = append(starts, coord{r,c})
			}
		}
	}
	return
}

func (p paths) Len() int { return len(p) }

func (p paths) Less(i, j int) bool {
	return p[i].length < p[j].length
}

func (p paths) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *paths) Push(x any) {
	item := x.(*path)
	*p = append(*p, item)
}

func (p *paths) Pop() any {
	old := *p
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	*p = old[:n-1]
	return item
}