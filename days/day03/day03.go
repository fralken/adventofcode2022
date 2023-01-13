package day03

import (
	"aoc2022/utils"
	"strings"
)

func FirstStar() {
	utils.Star(3, 1, "sum of priorities", firstStar)
}

func SecondStar() {
	utils.Star(3, 2, "sum of priorities", secondStar)
}

func firstStar(content string) int {
	ruckSacks := parseRucksaks(content)
	items := ruckSacks.findCommonItems(findCommonItem)
	return sumPriorities(items)
}

func secondStar(content string) int {
	ruckSacks := parseRucksaksGroups(content)
	items := ruckSacks.findCommonItems(findCommonItemInGroup)
	return sumPriorities(items)
}

type ruckSacks []string
type ruckSacksGroup []ruckSacks

func parseRucksaks(content string) (ruckSacks ruckSacksGroup) {
	lines := strings.Split(content,"\n")
	for i, line := range lines {
		ruckSacks = append(ruckSacks, []string{})
		split := len(line)/2
		ruckSacks[i] = append(ruckSacks[i], line[:split])
		ruckSacks[i] = append(ruckSacks[i], line[split:])
	}
	return
}

func parseRucksaksGroups(content string) (groups ruckSacksGroup) {
	lines := strings.Split(content,"\n")
	for i :=0; i < len(lines); i += 3 {
		ruckSacks := make([]string, 3)
		ruckSacks[0] = lines[i]
		ruckSacks[1] = lines[i+1]
		ruckSacks[2] = lines[i+2]
		groups = append(groups, ruckSacks)
	}
	return
}

func (rsg ruckSacksGroup) findCommonItems(find func(ruckSacks) byte) []byte {
	common := make([]byte, len(rsg))
	for i, ruckSack := range rsg {
		common[i] = find(ruckSack)
	}
	return common
}

func findCommonItem(sacks ruckSacks) byte {
	for _, f := range sacks[0] {
		for _, s := range sacks[1] {
			if f == s { 
				return byte(f) 
			}
		}
	}
	return 0
}

func findCommonItemInGroup(sacks ruckSacks) byte {
	for _, f := range sacks[0] {
		for _, s := range sacks[1] {
			for _, t := range sacks[2] {
				if f == s && s == t { 
					return byte(f) 
				}
			}
		}
	}
	return 0
}

func priority(item byte) int {
	switch {
		case 'a' <= item && item <= 'z': return int(item - 'a' + 1)
		case 'A' <= item && item <= 'Z': return int(item - 'A' + 27)
	}
	return 0
}

func sumPriorities(items []byte) (sum int) {
	for _, item := range items {
		sum += priority(item)
	}
	return sum
}