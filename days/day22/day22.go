package day22

import (
	"aoc2022/utils"
	"math"
	"strings"
)

func FirstStar() {
	utils.Star(22, 1, "final password (map)", firstStar)
}

func SecondStar() {
	utils.Star(22, 2, "final password (cube)", secondStar)
}

func firstStar(content string) int {
	grid, instr := parseGridAndDirections(content)
	return navigate(grid, instr)
}

func secondStar(content string) int {
	cube, instr := parseCubeAndDirections(content)
	return navigate(cube, instr)
}

type world interface {
	start() coord
	move(coord, coord) (coord, coord)
}
type grid []string
type coord struct {
	r, c int
}
type face struct {
	neighbours [4]*face
	pos coord
}
type cube struct {
	grid []string
	faces [6]*face
	side int
}

func parseGridAndDirections(content string) (grid, string) {
	parts := strings.Split(content, "\n\n")
	grid := strings.Split(parts[0], "\n")
	return grid, parts[1]
}

func parseCubeAndDirections(content string) (cube, string) {
	parts := strings.Split(content, "\n\n")
	cube := makeCube(parts[0])
	return cube, parts[1]
}

func navigate(world world, instr string) int {
	dir := coord{0, 1}
	pos := world.start()
	turn := false
	for i := 0; i < len(instr); {
		if turn {
			dir, i = parseTurn(instr, i, dir)
			turn = false
		} else {
			var steps int
			steps, i = parseMove(instr, i)
			for steps > 0 {
				pos, dir = world.move(pos, dir)
				steps--
			}
			turn = true
		}
	}
	return password(pos, dir)
}

func makeCube(content string) cube {
	count := 0
	for _, c := range content {
		if c == '.' || c == '#' { count++ }
	}
	side := int(math.Sqrt(float64(count / 6)))
	grid := strings.Split(content, "\n")

	// create faces
	var faces [6]*face
	f := 0
	for r := 0; r < len(grid); r += side {
		for c := 0; c < len(grid[r]); c += side {
			if grid[r][c] != ' ' {
				faces[f] = &face { [4]*face{}, coord{ r, c } }
				f++
			}
		}
	}

	// find direct neighbours (5 edges)
	for i, face := range faces {
		for j, other := range faces {
			if i != j {
				pos := face.pos
				if other.pos == (coord{ pos.r - side, pos.c }) { // north
					face.neighbours[0] = other
				} else if other.pos == (coord{ pos.r, pos.c + side }) { // east
					face.neighbours[1] = other
				} else if other.pos == (coord{ pos.r + side, pos.c }) { // south
					face.neighbours[2] = other
				} else if other.pos == (coord{ pos.r, pos.c - side }) { // west
					face.neighbours[3] = other
				}
			}
		}
	}

	// find folded neighbours (7 edges) moving clockwise
	cw := func (n int) int { return utils.Mod(n + 1, 4) }
	missing := 7
	for missing > 0 {
		for _, face := range faces {
			for n := 0; n < 4; n++ {
				if face.neighbours[n] == nil {
					fcw := face.neighbours[cw(n)]
					if fcw != nil {
						o := fcw.orientationOf(face)
						fcwn := fcw.neighbours[cw(o)]
						if fcwn != nil {
							face.neighbours[n] = fcwn
							o := fcwn.orientationOf(fcw)
							fcwn.neighbours[cw(o)] = face
							missing--
							break
						}
					}
				}
			}
		} 
	}

	return cube{ grid, faces, side }
}

func turn(dir byte, c coord) coord {
	dirs := map[byte]map[coord]coord{
		'R': {
				{-1, 0}: { 0, 1},
				{ 0,-1}: {-1, 0},
				{ 1, 0}: { 0,-1},
				{ 0, 1}: { 1, 0},
			},
		'L': {
				{-1, 0}: { 0,-1},
				{ 0,-1}: { 1, 0},
				{ 1, 0}: { 0, 1},
				{ 0, 1}: {-1, 0},
			},
	}
	return dirs[dir][c]
}

func dirValue(c coord) int {
	val := map[coord]int{
		{ 0, 1}: 0,
		{ 1, 0}: 1,
		{ 0,-1}: 2,
		{-1, 0}: 3,
	}
	return val[c]
}

func (g grid) start() coord {
	c := 0
	for c < len(g[0]) {
		if g[0][c] == '.' {
			break
		}
		c++
	}
	return coord{0,c}
}

func (g grid) move(pos, dir coord) (coord, coord) {
	r := pos.r
	c := pos.c
	if dir.r == 0 {
		c = utils.Mod(c + dir.c, len(g[r]))
		for g[r][c] == ' ' {
			c = utils.Mod(c + dir.c, len(g[r]))
		}
	} else if dir.c == 0 {
		r = utils.Mod(r + dir.r, len(g))
		for c >= len(g[r]) || g[r][c] == ' ' { 
			r = utils.Mod(r + dir.r, len(g))
		}
	}
	if g[r][c] == '#' {
		return pos, dir
	} else {
		return coord{r, c}, dir
	}
}

func (c cube) start() coord {
	return coord{ c.faces[0].pos.r, c.faces[0].pos.c }
}

func (cb cube) move(pos, dir coord) (newPos, newDir coord) {
	face, r, c := cb.findFace(pos)
	newFace := face
	newDir = dir
	edge := cb.side - 1
	if dir.r == 0 { // move horizontally
		c += dir.c
		if c < 0 {
			newFace = face.neighbours[3] // west
			o := newFace.orientationOf(face)
			if o == 0 { // west to north
				newDir = coord{ 1, 0 }
				c = r
				r = 0
			} else if o == 1 { // west to east
				c = edge
			} else if o == 2 { // west to south
				newDir = coord{ -1, 0 }
				c = edge - r
				r = edge
			} else if o == 3 { // west to west
				newDir = coord{ 0, 1 }
				c = 0
				r = edge - r
			}
		} else if c > edge {
			newFace = face.neighbours[1] // east
			o := newFace.orientationOf(face)
			if o == 0 { // east to north
				newDir = coord{ 1, 0 }
				c = edge - r
				r = 0
			} else if o == 1 { // east to east
				newDir = coord{ 0, -1 }
				c = edge
				r = edge - r
			} else if o == 2 { // east to south
				newDir = coord{ -1, 0 }
				c = r
				r = edge
			} else if o == 3 { // east to west
				c = 0
			}
		}
	} else if dir.c == 0 { // move vertically
		r += dir.r
		if r < 0 {
			newFace = face.neighbours[0] // north
			o := newFace.orientationOf(face)
			if o == 0 { // north to north
				newDir = coord{ 1, 0 }
				r = 0
				c = edge - c
			} else if o == 1 { // north to east
				newDir = coord{ 0, -1 }
				r = edge - c
				c = edge
			} else if o == 2 { // north to south
				r = edge
			} else if o == 3 { // north to west
				newDir = coord{ 0, 1 }
				r = c
				c = 0
			}
		} else if r > edge {
			newFace = face.neighbours[2] // south
			o := newFace.orientationOf(face)
			if o == 0 { // south to north
				r = 0
			} else if o == 1 { // south to east
				newDir = coord{ 0, -1 }
				r = c
				c = edge
			} else if o == 2 { // south to south
				newDir = coord { -1, 0 }
				r = edge
				c = edge - c
			} else if o == 3 { // south to west
				newDir = coord{ 0, 1 }
				r = edge - c
				c = 0
			}
		}
	}
	newPos = coord{ newFace.pos.r + r, newFace.pos.c + c }
	if cb.grid[newPos.r][newPos.c] == '#' {
		newPos = pos
		newDir = dir
	}
	return
}

func (cb cube) findFace(pos coord) (*face, int, int) {
	r := pos.r % cb.side
	c := pos.c % cb.side
	p := coord{ pos.r - r, pos.c - c }
	for _, f := range cb.faces {
		if f.pos == p {
			return f, r, c
		}
	}
	return nil, -1, -1
}

func (f face) orientationOf(n *face) int {
	for i := range f.neighbours {
		if f.neighbours[i] == n {
			return i
		}
	}
	return -1
}

func parseMove(instr string, id int) (int, int) {
	newId := id
	for newId < len(instr) && instr[newId] >= '0' && instr[newId] <= '9' {
		newId++
	}
	return utils.StringToInt(instr[id:newId]), newId
}

func parseTurn(instr string, id int, dir coord) (coord, int) {
	return turn(instr[id], dir), id + 1
}

func password(pos, dir coord) int {
	return 1000 * (pos.r + 1) + 4 * (pos.c + 1) + dirValue(dir)
}
