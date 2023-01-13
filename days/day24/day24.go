package day24

import (
	"aoc2022/utils"
	"fmt"
	"strings"
)

func FirstStar() {
	utils.Star(24, 1, "minutes required to reach the goal", firstStar)
}

func SecondStar() {
	utils.Star(24, 2, "minutes required to reach the goal 3 times", secondStar)
}

func firstStar(content string) int {
	b, h, w := parseValley(content)
	start := coord{-1, 0}
	end := coord{h, w-1}
	return move(start, end, b, h, w, 1)
}

func secondStar(content string) int {
	b, h, w := parseValley(content)
	start := coord{-1, 0}
	end := coord{h, w-1}
	return move(start, end, b, h, w, 3)
}

type coord struct {
	r, c int
}
type blizzards map[rune]valley
type valley map[coord]bool
type paths map[int]valley

func parseValley(content string) (b blizzards, h int, w int) {
	lines := strings.Split(content, "\n")
	b = make(blizzards)
	h = len(lines) - 2
	w = len(lines[0]) - 2
	for r, line := range lines[1:len(lines)-1] {
		for c, x := range line[1:len(line)-1] {
			if x != '.' {
				if _, ok := b[x]; !ok {
					b[x] = make(map[coord]bool)
				}
				b[x][coord{r,c}] = true
			}
		}
	}
	return
}

func (b *blizzards) update(h int, w int) blizzards {
	newBl := make(blizzards)
	for k, v := range *b {
		newPos := make(map[coord]bool)
		for p := range v {
			if k == '^' {
				newPos[coord{utils.Mod(p.r - 1, h), p.c}] = true
			} else if k == '>' {
				newPos[coord{p.r, utils.Mod(p.c + 1, w)}] = true
			} else if k == 'v' {
				newPos[coord{utils.Mod(p.r + 1, h), p.c}] = true
			} else if k == '<' {
				newPos[coord{p.r, utils.Mod(p.c - 1, w)}] = true
			}
		}
		newBl[k] = newPos
	}
	return newBl
}

func (b blizzards) fillValley() valley {
	valley := make(valley)
	for _, v := range b {
		valley.merge(&v)
	}
	return valley
}

func (to *valley) merge(from *valley) {
	for k, v := range *from {
		(*to)[k] = v
	} 
}

func (ps paths) getShortest() (coord, int) {
	min := 0
	for k := range ps {
		if min == 0 || min > k {
			min = k
		}
	}
	var c coord
	for c = range ps[min] { break }
	delete(ps[min], c)
	if len(ps[min]) == 0 {
		delete(ps, min)
	}
	return c, min
}

func (ps paths) addPos(step int, pos coord) {
	if _, ok := ps[step]; !ok {
		ps[step] = make(valley)
	}
	ps[step][pos] = true
}

func move(start, end coord, b blizzards, h, w, times int) int {
	paths := make(paths)
	count := 0
	totalCount := 0
	for times > 0 {
		paths.addPos(count, start)
		blizzards := []blizzards{b}
		valleys := []valley{b.fillValley()}
		for len(paths) > 0 {
			pos, st := paths.getShortest()
			if count == 0 || count > st {
				st++
				if st == len(blizzards) {
					blizzards = append(blizzards, blizzards[st-1].update(h, w))
					valleys = append(valleys, blizzards[st].fillValley())
				}
				dir := []coord{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
				for _, d := range dir {
					np := coord{pos.r + d.r, pos.c + d.c}
					if np == end && (count == 0 || count > st) {
						count = st
					}
					if _, ok := valleys[st][np]; !ok && np.r >= 0 && np.r < h && np.c >= 0 && np.c < w {
						paths.addPos(st, np)
					}
				}
				if _, ok := valleys[st][pos]; !ok {
					paths.addPos(st, pos)
				}
			}
		}
		times--
		start, end = end, start
		b = blizzards[count]
		totalCount += count
		count = 0
	}
	return totalCount
}

func (v valley) Print(pos coord, h, w, count int) {
	grid := make([][]byte, h)
	for i := 0; i < h; i++ {
		grid[i] = make([]byte, w)
		for j := 0; j < w; j++ {
			if _, ok := v[coord{i, j}]; ok {
				grid[i][j] = 'X'
			} else {
				grid[i][j] = '.'
			}
		}
	}
	fmt.Printf("%d   %d,%d\n", count, pos.r, pos.c)
	for i := 0; i < h; i++ {
		fmt.Printf("%s\n", string(grid[i]))
	}
}
