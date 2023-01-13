package day23

import (
	"aoc2022/utils"
	"fmt"
	"reflect"
	"strings"
)

func FirstStar() {
	utils.Star(23, 1, "number of empty ground tiles", firstStar)
}

func SecondStar() {
	utils.Star(23, 2, "first round where no elf moves", secondStar)
}

func firstStar(content string) int {
	elves := parseElves(content)
	for i := 0; i < 10; i++ {
		elves = elves.move(i)
	}
	return elves.emptyGround()
}

func secondStar(content string) int {
	elves := parseElves(content)
	i := 0
	for {
		newElves := elves.move(i)
		i++
		if reflect.DeepEqual(elves, newElves) {
			break
		}
		elves = newElves
	}
	return i
}

type coord struct {
	r, c int
}
type elves map[coord]coord

func parseElves(content string) elves {
	lines := strings.Split(content, "\n")
	elves := make(elves)
	for i, line := range lines {
		for j, c := range line {
			if c == '#' {
				elves[coord{i, j}] = coord{i, j}
			}
		}
	}
	return elves
}

func nextMove(i, offset int) int {
	return (i + offset) % 4
}

func (es elves) move(offset int) elves {
	nextPos := make(map[coord]int)
	for k := range es {
		if es.needsToMove(k) {
			for i := 0; i < 4; i++ {
				canMove := false
				dir := nextMove(i, offset)
				switch dir {
				case 0: // North
					_, ok1 := es[coord{k.r - 1, k.c - 1}]
					_, ok2 := es[coord{k.r - 1, k.c}]
					_, ok3 := es[coord{k.r - 1, k.c + 1}]
					if !(ok1 || ok2 || ok3) {
						es[k] = coord{k.r - 1, k.c}
						nextPos[es[k]]++
						canMove = true
					}
				case 1: // South
					_, ok1 := es[coord{k.r + 1, k.c - 1}]
					_, ok2 := es[coord{k.r + 1, k.c}]
					_, ok3 := es[coord{k.r + 1, k.c + 1}]
					if !(ok1 || ok2 || ok3) {
						es[k] = coord{k.r + 1, k.c}
						nextPos[es[k]]++
						canMove = true
					}
				case 2: // West
					_, ok1 := es[coord{k.r - 1, k.c - 1}]
					_, ok2 := es[coord{k.r, k.c - 1}]
					_, ok3 := es[coord{k.r + 1, k.c - 1}]
					if !(ok1 || ok2 || ok3) {
						es[k] = coord{k.r, k.c - 1}
						nextPos[es[k]]++
						canMove = true
					}
				case 3: // East
					_, ok1 := es[coord{k.r - 1, k.c + 1}]
					_, ok2 := es[coord{k.r, k.c + 1}]
					_, ok3 := es[coord{k.r + 1, k.c + 1}]
					if !(ok1 || ok2 || ok3) {
						es[k] = coord{k.r, k.c + 1}
						nextPos[es[k]]++
						canMove = true
					}
				}
				if canMove {
					break
				}
			}
		}
	}
	newElves := make(elves)
	for k, v := range es {
		if w, ok := nextPos[v]; ok && w <= 1 {
			newElves[v] = v
		} else {
			newElves[k] = k
		}
	}
	return newElves
}

func (es elves) emptyGround() int {
	minr, minc, maxr, maxc := es.bounds()
	return (maxr - minr + 1) * (maxc - minc + 1) - len(es)
}

func (es elves) bounds() (minr, minc, maxr, maxc int) {
	for k := range es {
		minr = k.r
		maxr = k.r
		minc = k.c
		maxc = k.c
		break
	}
	for k := range es {
		if minr > k.r { minr = k.r }
		if minc > k.c { minc = k.c }
		if maxr < k.r { maxr = k.r }
		if maxc < k.c { maxc = k.c }
	}
	return
}

func (es elves) needsToMove(c coord) bool {
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			d := coord{c.r + i, c.c + j}
			if _, ok := es[d]; ok && c != d {
				return true
			}
		}
	}
	return false
}

func (es elves) Print() {
	minr, minc, maxr, maxc := es.bounds()
	h := maxr - minr + 1
	w := maxc - minc + 1
	grid := make([][]byte, h)
	for i := 0; i < h; i++ {
		grid[i] = make([]byte, w)
		for j := 0; j < w; j++ {
			if _, ok := es[coord{minr + i, minc + j}]; ok {
				grid[i][j] = '#'
			} else {
				grid[i][j] = '.'
			}
		}
	}
	fmt.Printf("%d,%d\n", minr, minc)
	for i := 0; i < h; i++ {
		fmt.Printf("%s\n", string(grid[i]))
	}
	fmt.Printf("%d,%d\n", maxr, maxc)
}