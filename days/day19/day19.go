package day19

import (
	"aoc2022/utils"
	"fmt"
	"strings"
)

func FirstStar() {
	utils.Star(19, 1, "quality level", firstStar)
}

func SecondStar() {
	utils.Star(19, 2, "product of open geodes", secondStar)
}

func firstStar(content string) int {
	blueprints := parseBluePrints(content)
	sum := 0
	s := state{
		oreRobots: 1,
		clayRobots: 0,
		obsidianRobots: 0,
		geodeRobots: 0,
		ores: 0,
		clays: 0,
		obsidians: 0,
		geodes: 0,
		time: 24,
	}
	for _, b := range blueprints {
		geodes := b.openGeodes(s, false, false, false, 0)
		sum += geodes * b.id
	}
	return sum
}

func secondStar(content string) int {
	blueprints := parseBluePrints(content)
	if len(blueprints) > 3 { blueprints = blueprints[:3] }
	product := 1
	s := state{
		oreRobots: 1,
		clayRobots: 0,
		obsidianRobots: 0,
		geodeRobots: 0,
		ores: 0,
		clays: 0,
		obsidians: 0,
		geodes: 0,
		time: 32,
	}
	for _, b := range blueprints {
		product *= b.openGeodes(s, false, false, false, 0)
	}
	return product
}

type blueprint struct {
	id int
	oreOre int
	clayOre int
	obsidianOre int
	obsidianClay int
	geodeOre int
	geodeObsidian int
	maxOre int
}

type state struct {
	oreRobots int
	clayRobots int
	obsidianRobots int
	geodeRobots int
	ores int
	clays int
	obsidians int
	geodes int
	time int
}

func parseBluePrints(content string) []blueprint {
	lines := strings.Split(content, "\n")
	b := make([]blueprint, len(lines))
	for i, line := range lines {
		fmt.Sscanf(line, "Blueprint %d: Each ore robot costs %d ore. Each clay robot costs %d ore. Each obsidian robot costs %d ore and %d clay. Each geode robot costs %d ore and %d obsidian.",
			&b[i].id,
			&b[i].oreOre,
			&b[i].clayOre,
			&b[i].obsidianOre,
			&b[i].obsidianClay,
			&b[i].geodeOre,
			&b[i].geodeObsidian,
		)
		b[i].maxOre = utils.Max(b[i].oreOre, utils.Max(b[i].clayOre, utils.Max(b[i].obsidianOre, b[i].geodeOre)))
	}
	return b
}

func (b *blueprint) openGeodes(s state, oreSkipped, claySkipped, obsidianSkipped bool, best int) int {
	// prune this branch execution if the optimistic prediction is lower than the current best count
	if s.isPredictionWorse(best) { return best }

	// Robots can be built if these conditions are verified:
	// - there is enough material to build a robot
	// - the current number of robots is lower then the max amount
	//   of material needed to create other robots. This because it is
	//   not necessary to create more material per minute than needed to
	//   create a robot in each minute
	// If we could build a robot in the previous minute but didn't,
	// do not build it now but wait for a different robot to be built first
	canBuildOreRobot := !oreSkipped &&
		s.ores >= b.oreOre &&
		s.oreRobots < b.maxOre
	canBuildClayRobot := !claySkipped &&
		s.ores >= b.clayOre &&
		s.clayRobots < b.obsidianClay
	canBuildObsidianRobot := !obsidianSkipped &&
		s.ores >= b.obsidianOre &&
		s.clays >= b.obsidianClay &&
		s.obsidianRobots < b.geodeObsidian
	canBuildGeodeRobot := s.ores >= b.geodeOre &&
		s.obsidians >= b.geodeObsidian

	s.ores += s.oreRobots
	s.clays += s.clayRobots
	s.obsidians += s.obsidianRobots
	s.geodes += s.geodeRobots
	s.time -= 1

	if s.time == 0 { return s.geodes }

	// give priority to geode robots, this is the best choice
	// In this case there is no need to test building other robots
	// so we return the result after this execution
	if canBuildGeodeRobot {
		s.geodeRobots += 1
		s.ores -= b.geodeOre
		s.obsidians -= b.geodeObsidian
		return b.openGeodes(s, false, false, false, best)
	}
	// for the other robots we compare which is the best case and
	// compare with the current best case
	if canBuildObsidianRobot {
		s := s // make a copy of the state
		s.obsidianRobots += 1
		s.ores -= b.obsidianOre
		s.clays -= b.obsidianClay
		best = utils.Max(best, b.openGeodes(s, false, false, false, best))
	}
	if canBuildClayRobot {
		s := s // make a copy of the state
		s.clayRobots += 1
		s.ores -= b.clayOre
		best = utils.Max(best, b.openGeodes(s, false, false, false, best))
	}
	if canBuildOreRobot {
		s := s // make a copy of the state
		s.oreRobots += 1
		s.ores -= b.oreOre
		best = utils.Max(best, b.openGeodes(s, false, false, false, best))
	}
	// try not building robots now but wait for a different robot to be built first
	// this case is meaningful when at least one robot could not be built in this minute
	if !canBuildOreRobot || !canBuildClayRobot || !canBuildObsidianRobot {
		best = utils.Max(best, b.openGeodes(s, canBuildOreRobot, canBuildClayRobot, canBuildObsidianRobot, best))
	}

	return best
}

// This is an optimistic prediction that is compared to the current best geodes count.
// If the prediction is lower, prune this branch execution.
// The prediction considers an optimistic case where from now on a new geode robot is
// created, providing a new geode per minute per robot. This in practice is the sum
// of first integers up to (time - 1), that is equal to (time - 1) * time / 2
// The optimistic count is:
// - current number of opened geodes
// - count of future geodes opened by existing robots every minute
// - count of geodes opened by robots created every minute from now on
func (s state) isPredictionWorse(best int) bool {
	return s.geodes + s.geodeRobots * s.time + (s.time - 1) * s.time / 2 < best
}
