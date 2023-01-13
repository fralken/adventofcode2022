package day16

import (
	"aoc2022/utils"
	"fmt"
	"strings"
)

func FirstStar() {
	utils.Star(16, 1, "maximum pressure", firstStar)
}

func SecondStar() {
	utils.Star(16, 2, "maximum pressure", secondStar)
}

func firstStar(content string) int {
	valves := parseValves(content)
	graph := valves.shortestPaths()
	bitmask := valves.makeBitmask()
	return valves.goTo1(0, graph, "AA", 0, 0, 30, bitmask)
}

func secondStar(content string) int {
	valves := parseValves(content)
	graph := valves.shortestPaths()
	bitmask := valves.makeBitmask()
	// to reduce the number of possible paths we discard flows that are lower than minFlow,
	// that is half of the flow obtained in the available time,
	// the sum of two flows we compute later must be greater than minFlow
	minFlow := valves.goTo1(0, graph, "AA", 0, 0, 26, bitmask) / 2
	flows := valves.goTo2(0, graph, "AA", 0, 0, 26, minFlow, bitmask)
	max := 0
	for i, fi := range flows[:len(flows)-1] {
		for _, fj := range flows[i+1:] {
			// if paths do not overlap we can sum the flows
			if fi.valves & fj.valves == 0 {
				flow := fi.flow + fj.flow
				if flow > max { max = flow }
			}

		}
	}
	return max
}

type valve struct {
	rate int
	links string
}
type graph map[string]valve
type paths map[string]int
type graphPaths map[string]paths
type route struct {
	flow int
	valves uint32
}
type bitmask map[string]uint32

func parseValves(content string) graph {
	lines := strings.Split(content, "\n")
	valves := make(graph)
	for _, line := range lines {
		parts := strings.Split(line, "; ")
		if strings.HasPrefix(parts[1], "tunnels lead to valves ") {
			parts[1] = parts[1][len("tunnels lead to valves "):]
		} else {
			parts[1] = parts[1][len("tunnel leads to valve "):]
		}
		var name string
		var rate int
		fmt.Sscanf(parts[0], "Valve %s has flow rate=%d;", &name, &rate)
		valves[name] = valve{ rate,	parts[1] }
	}
	return valves
}

func (valves graph) shortestPaths() graphPaths {
	graph := make(graphPaths)
	max := len(valves)
	for v := range valves {
		graph[v] = make(paths)
		for w := range valves {
			if v == w {
				graph[v][w] = 0
			} else if strings.Contains(valves[v].links, w) {
				graph[v][w] = 1
			} else {
				graph[v][w] = max
			}
		}
	}
	for i := range valves {
		for j := range valves {
			for k := range valves {
				if graph[j][k] > graph[j][i] + graph [i][k] {
					graph[j][k] = graph[j][i] + graph [i][k]
				}
			}
		}
	}
	return graph
}

func (valves graph) goTo1(opened uint32, paths graphPaths, current string, flow int, time int, end int, bitmask bitmask) int {
	max := flow
	for k := range valves {
		if opened & bitmask[k] > 0 || k == current || valves[k].rate == 0 {
			continue
		}
		elapsed := paths[current][k] + 1
		if time + elapsed > end {
			continue
		}
		nextFlow := flow + (end - time - elapsed) * valves[k].rate
		nextTime := time + elapsed
		if next := valves.goTo1(opened | bitmask[k], paths, k, nextFlow, nextTime, end, bitmask); next > max {
			max = next
		}
	}
	return max
}

func (valves graph) goTo2(opened uint32, paths graphPaths, current string, flow int, time int, end int,
	minFlow int, bitmask bitmask) []route {
	routes := []route{}
	if flow > minFlow {
		routes = append(routes, route{flow, opened})
	}
	for k := range valves {
		if (bitmask[k] & opened) > 0 || k == current || valves[k].rate == 0 {
			continue
		}
		elapsed := paths[current][k] + 1
		if time + elapsed > end {
			continue
		}
		nextFlow := flow + (end - time - elapsed) * valves[k].rate
		nextTime := time + elapsed
		next := valves.goTo2(opened | bitmask[k], paths, k, nextFlow, nextTime, end, minFlow, bitmask)
		routes = append(routes, next...)
	}
	return routes
}

// we use a bitmask to store valves already opened
// we consider only valves that have positive rate
// that luckily are less than 32, without this
// optimization the computation takes too long
func (valves graph) makeBitmask() bitmask {
	bitmask := make(bitmask)
	b := 0
	for k, v := range valves {
		if v.rate > 0 {
			bitmask[k] = 1 << b
			b++
		}
	}
	return bitmask
}