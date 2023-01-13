package day18

import (
	"aoc2022/utils"
	"fmt"
	"strings"
)

func FirstStar() {
	utils.Star(18, 1, "surface area of lava droplets", firstStar)
}

func SecondStar() {
	utils.Star(18, 2, "exterior surface area of lava droplets", secondStar)
}

func firstStar(content string) int {
	cubes := parseCubes(content)
	return cubes.countSurfaces()
}

func secondStar(content string) int {
	cubes := parseCubes(content)
	return cubes.countExteriorSurfaces()
}

type point struct {
	x, y, z int
}

type cubes map[point]bool

func parseCubes(content string) cubes {
	lines := strings.Split(content, "\n")
	cubes := make(cubes)
	for _, line := range lines {
		var p point
		fmt.Sscanf(line, "%d,%d,%d", &p.x, &p.y, &p.z)
		cubes[p] = true
	}
	return cubes
}

func (cs cubes) countSurfaces() int {
	sides := []point{{-1,0,0},{1,0,0},{0,-1,0},{0,1,0},{0,0,-1},{0,0,1}} 
	surfaces := 0
	for c := range cs {
		for _, s := range sides {
			if !cs[point{c.x+s.x, c.y+s.y, c.z+s.z}] {
				surfaces++
			}
		}
	}
	return surfaces
}

func (cs cubes) countExteriorSurfaces() int {
	sides := []point{{-1,0,0},{1,0,0},{0,-1,0},{0,1,0},{0,0,-1},{0,0,1}} 
	min, max := cs.findBoundaries()
	water := make(cubes)
	start := point{min.x-1,min.y-1,min.z-1}
	// find where the water flows, water surrounds the bounding box of cubes
	water[start] = true
	flow := []point{start}
	for len(flow) > 0 {
		l := flow[len(flow)-1]
		flow = flow[:len(flow)-1]
		for _, s := range sides {
			p := point{l.x + s.x, l.y + s.y, l.z + s.z}
			if !water[p] && !cs[p] &&
				p.x >= min.x-1 && p.x <= max.x+1 &&
				p.y >= min.y-1 && p.y <= max.y+1 &&
				p.z >= min.z-1 && p.z <= max.z+1 {
				water[p] = true
				flow = append(flow, p)
			}
		}
	}
	surfaces := 0
	for c := range cs {
		for _, s := range sides {
			// count surfaces facing water
			if water[point{c.x+s.x, c.y+s.y, c.z+s.z}] {
				surfaces++
			}
		}
	}
	return surfaces
}

func (cs cubes) findBoundaries() (min point, max point) {
	for c := range cs {
		min = c
		max = c
		break
	}
	for c := range cs {
		if min.x > c.x { min.x = c.x }
		if min.y > c.y { min.y = c.y }
		if min.z > c.z { min.z = c.z }
		if max.x < c.x { max.x = c.x }
		if max.y < c.y { max.y = c.y }
		if max.z < c.z { max.z = c.z }
	}
	return
}