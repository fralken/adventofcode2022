package day14

import (
	"aoc2022/utils"
	"strings"
)

func FirstStar() {
	utils.Star(14, 1, "units of sand", firstStar)
}

func SecondStar() {
	utils.Star(14, 2, "units of sand", secondStar)
}

func firstStar(content string) int {
	rocks := parseRocks((content))
	return rocks.drop1(coord{500, 0})
}

func secondStar(content string) int {
	rocks := parseRocks((content))
	return rocks.drop2(coord{500, 0})
}

type coord struct {
	x, y int
}

type rock [2]coord
type rocks []rock

func parseRocks(content string) (rocks rocks) {
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		coords := strings.Split(line, " -> ")
		for i := range coords[:len(coords)-1] {
			xy := strings.Split(coords[i], ",")
			end1 := coord {
				utils.StringToInt(xy[0]),
				utils.StringToInt(xy[1]),
			}
			xy = strings.Split(coords[i+1], ",")
			end2 := coord {
				utils.StringToInt(xy[0]),
				utils.StringToInt(xy[1]),
			}
			rocks = append(rocks, [2]coord{end1, end2})
		}
	}
	return
}

func (r rocks) collide(c coord) bool {
	for _, rock := range r {
		if rock.hit(c) {
			return true
		}
	}
	return false
}

func (r rock) hit(c coord) bool {
	if r[0].x == r[1].x {
		if c.x == r[0].x {
			if (c.y <= r[0].y && c.y >= r[1].y) ||
				(c.y >= r[0].y && c.y <= r[1].y) {
					return true
			}
		}
	}
	if r[0].y == r[1].y {
		if c.y == r[0].y {
			if (c.x <= r[0].x && c.x >= r[1].x) ||
				(c.x >= r[0].x && c.x <= r[1].x) {
					return true
			}
		}
	}
	return false
}

func (r rocks) findBottom() int {
	val := 0
	for _, rock := range r {
		if rock[0].y > val { val = rock[0].y }
		if rock[1].y > val { val = rock[1].y }
	}
	return val
}

func (r rocks) drop1(start coord) int {
	sand := make(map[coord]bool)
	bottom := r.findBottom()
	falling := []coord{start}
	full := false
	for !full {
		p := falling[len(falling)-1]
		next := coord{p.x, p.y+1}
		if sand[next] || r.collide(next) {
			next = coord{p.x-1, p.y+1}
			if sand[next] || r.collide(next) {
				next = coord{p.x+1, p.y+1}
				if sand[next] || r.collide(next) {
					sand[p] = true
					falling = falling[:len(falling)-1]
					continue
				}
			}
		}
		full = p.y > bottom
		falling = append(falling, next)
	}
	return len(sand)
}

func (r rocks) drop2(start coord) int {
	sand := make(map[coord]bool)
	bottom := r.findBottom() + 2
	falling := []coord{start}
	for len(falling) > 0 {
		p := falling[len(falling)-1]
		next := coord{p.x, p.y+1}
		if sand[next] || next.y == bottom || r.collide(next) {
			next = coord{p.x-1, p.y+1}
			if sand[next] || next.y == bottom || r.collide(next) {
				next = coord{p.x+1, p.y+1}
				if  sand[next] || next.y == bottom || r.collide(next) {
					sand[p] = true
					falling = falling[:len(falling)-1]
					continue
				}
			}
		}
		falling = append(falling, next)
	}
	return len(sand)
}