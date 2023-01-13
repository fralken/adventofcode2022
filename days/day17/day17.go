package day17

import (
	"aoc2022/utils"
)

func FirstStar() {
	utils.Star(17, 1, "height of tower of 2022 rocks", firstStar)
}

func SecondStar() {
	utils.Star(17, 2, "height of tower of 1000000000000 rocks", secondStar)
}

func firstStar(content string) int {
	return execute(content, 2022)
}

func secondStar(content string) int {
	return execute(content, 1000000000000)
}

type coord struct {
	x, y int
}

type rock struct {
	pos coord
	points []coord
	w int
	h int
}

type state struct {
	top, remainingRocks int
}

type key struct {
	rockId int
	jetId int
	caveDepth [7]int
}

type cave []uint8

func execute(jets string, maxRocks int) int {
	findCycles := true
	top := 0

	offset := coord{2, 4}
	down := coord{0, -1}

	// the cave is represented as an array of bitmasks 7 bit wide
	// where 1's represents rocks or floor,
	// it is initialized with the floor (all 1's)
	cave := cave{(1 << 7) - 1}

	// this keeps track of the state:
	// we check a repetition of rock's index, jet's index and
	// depth map of the cave, that is the distance from top
	// of the tower of tallest rocks at each horizontal position.
	// The state is the current top position and the number
	// of remaining rocks to fall.
	// When we find a repetition of rock's id, jet's id and depth map
	// we can compute the tower's height per cycle and
	// the number of falling rocks per cycle.
	cache := make(map[key]state)

	rocks := generateRocks()
	rockId := 0
	rockCount := len(rocks)

	jetId := 0
	jetsCount := len(jets)

	totalCyclesHeight := 0
	remainingRocks := maxRocks
	for remainingRocks > 0 {
		rock := rocks[rockId]
		rockId = (rockId + 1) % rockCount

		rock.pos.x = offset.x
		rock.pos.y = offset.y + top

		for {
			jet := jetDirection(jets[jetId])
			jetId = (jetId + 1) % jetsCount

			// move horizontally
			if cave.rockCanMove(&rock, &jet) {
				rock.pos.x += jet.x
			} 

			// move vertically
			if cave.rockCanMove(&rock, &down) {
				rock.pos.y += down.y
			} else {
				cave.update(&rock)
				rockHeight := rock.pos.y + rock.h - 1
				top = utils.Max(rockHeight, top)
				remainingRocks--
				if findCycles {
					depthMap := cave.computeDepthMap(top)
					k := key{rockId, jetId, depthMap}
					s, ok := cache[k]
					if ok {
						cycleHeight := top - s.top
						cycleRocks := s.remainingRocks - remainingRocks
						remainingCycles := remainingRocks / cycleRocks
						totalCyclesHeight = cycleHeight * remainingCycles
						remainingRocks = remainingRocks % cycleRocks
						findCycles = false
					} else {
						cache[k] = state{top, remainingRocks}
					}
				}
				break
			}
		}
	}
	return totalCyclesHeight + top
}

// generates a list of shapes with their width and height
// their position will be set when moving
func generateRocks() []rock {
	return []rock{
		{points: []coord{{0, 0}, {1, 0}, {2, 0}, {3, 0}}, h: 1, w: 4},
		{points: []coord{{1, 0}, {0, 1}, {1, 1}, {2, 1}, {1, 2}}, h: 3, w: 3},
		{points: []coord{{0, 0}, {1, 0}, {2, 0}, {2, 1}, {2, 2}}, h: 3, w: 3},
		{points: []coord{{0, 0}, {0, 1}, {0, 2}, {0, 3}}, h: 4, w: 1},
		{points: []coord{{0, 0}, {1, 0}, {0, 1}, {1, 1}}, h: 2, w: 2},
	}
}

// check if a rock collides against walls or floor of the cave
func (c *cave) rockCanMove(r *rock, jet *coord) bool {
	px := r.pos.x + jet.x
	py := r.pos.y + jet.y
	for _, i := range r.points {
		x := px + i.x
		if x < 0 || x > 6 {
			return false
		}
		y := py + i.y
		if len(*c) > y && (*c)[y] & (1 << x) != 0 {
			return false
		}
	}
	return true	
}

// update the tower of rocks in the cave
// resize the array representing the cave if more space is needed
func (c *cave) update(r *rock) {
	for _, i := range r.points {
		x := r.pos.x + i.x
		y := r.pos.y + i.y
		if len(*c) <= y {
			*c = append(*c, make([]uint8, y - len(*c) + 1)...)
		}
		(*c)[y] |= 1 << x
	}
}

// compute the distance from current top of the tower and
// the tallest rocks at each horizontal position
func (c *cave) computeDepthMap(top int) [7]int {
	res := [7]int{}
	for i := 0; i < 7; i++ {
		m := uint8(1 << i)
		for j := top; j >= 0; j-- {
			if (*c)[j] & m != 0 {
				res[i] = top - j
				break
			}
		}
	}
	return res
}

func jetDirection(jet byte) coord {
	if jet == '>' {
		return coord{1, 0}
	} else { // jet == '<'
		return coord{-1, 0}
	}
}
