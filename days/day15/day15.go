package day15

import (
	"aoc2022/utils"
	"fmt"
	"strings"
)

func FirstStar() {
	utils.Star(15, 1, "positions that do not contain a beacon", firstStar)
}

func SecondStar() {
	utils.Star(15, 2, "tuning frequency", secondStar)
}

func firstStar(content string) int {
	sensors, beacons := parseSensors(content)
	return sensors.noBeacons(beacons, 2000000)
}

func secondStar(content string) int {
	sensors, _ := parseSensors(content)
	return sensors.findBeacon(4000000)
}

type coord struct {
	x, y int
}

type sensors map[coord]int
type beacons map[coord]bool

func parseSensors(content string) (sensors, beacons) {
	lines := strings.Split(content, "\n")
	sensors := make(map[coord]int)
	beacons := make(map[coord]bool)
	for _, line := range lines {
		var sx, sy, bx, by int
		fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sx, &sy, &bx, &by)
		m := manhattan(sx, sy, bx, by)
		sensors[coord{sx, sy}] = m
		beacons[coord{bx, by}] = true
	}
	return sensors, beacons
}

func (s sensors) noBeacons(b beacons, y int) int {
	no := make(map[int]bool)
	for k, v := range s {
		for x := k.x - v; x <= k.x + v; x++ {
			if manhattan(k.x, k.y, x, y) <= v && !b[coord{x, y}] {
				no[x] = true
			}
		}
	}
	return len(no)
}

func (s sensors) findBeacon(max int) int {
	for k, v := range s {
		for dx := v + 1; dx > 0; dx-- {
			dy := v + 1 - dx;
			arr := []coord{{dx,dy}, {-dx,dy}, {-dx,-dy}, {dx,-dy}}
			for _, a := range arr {
				x := k.x + a.x
				y := k.y + a.y
				if x >= 0 && x <= max && y >= 0 && y <= max &&
					!s.inside(x, y) {
					return x * 4000000 + y
				}
			}
		}
	}
	return 0
}

func (s sensors) inside(x, y int) bool {
	for k, v := range s {
		if manhattan(k.x, k.y, x, y) <= v {
			return true
		}
	}
	return false
}

func manhattan(sx, sy, bx, by int) int {
	return utils.Abs(bx - sx) + utils.Abs(by - sy)
}
