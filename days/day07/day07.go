package day07

import (
	"aoc2022/utils"
	"strings"
)

func FirstStar() {
	utils.Star(7, 1, "sum of dir sizes below 100000", firstStar)
}

func SecondStar() {
	utils.Star(7, 2, "size of dir to be deleted", secondStar)
}

func firstStar(content string) int {
	root := parseCommands(content)
	return root.sumDirsLessThan(100000)
}

func secondStar(content string) int {
	root := parseCommands(content)
	return root.freeSpace(70000000 - root.size, 30000000, root.size)
}

type item struct {
	name string
	size int
	items []item
	parent *item
}

func parseCommands(content string) *item {
	root := &item{ "/", 0, []item{}, nil }
	current := root
	lines := strings.Split(content, "\n")
	for l := 0; l < len(lines); l++ {
		if lines[l][0] == '$' { // command
			if lines[l][2:4] == "cd" {
				switch lines[l][5:] {
				case "/":
					current = root
				case "..":
					if current.parent != nil {
						current = current.parent
					}
				default:
					name := lines[l][5:]
					it := current.findItem(name)
					if it == nil {
						current = current.addDir(name)
					} else {
						current = it
					}
				}
			}
		} else { // output
			if lines[l][0:3] == "dir" {
				name := lines[l][4:]
				it := current.findItem(name)
				if it == nil {
					current.addDir(name)
				}
			} else { // file
				parts := strings.Split(lines[l], " ")
				size := utils.StringToInt(parts[0])
				name := parts[1]
				it := current.findItem(name)
				if it == nil {
					current.addFile(name, size)
				}
			}
		}
	}
	return root
}

func (tree *item) findItem(name string) *item {
	for i := range tree.items {
		if tree.items[i].name == name {
			return &tree.items[i]
		}
	}
	return nil
}

func (tree *item) addDir(name string) *item {
	newItem := &item{ name, 0, []item{}, tree}
	tree.items = append(tree.items, *newItem)
	return newItem
}

func (tree *item) addFile(name string, size int) *item {
	newItem := &item{ name, size, nil, tree}
	tree.items = append(tree.items, *newItem)
	for t := tree; t != nil; t = t.parent {
		t.size += size
	}
	return newItem
}

func (root *item) sumDirsLessThan(size int) (sum int) {
	if root.size < size {
		sum += root.size
	}
	for _, it := range root.items {
		if it.items != nil {
			sum += it.sumDirsLessThan(size)
		}
	}
	return
}

func (root *item) freeSpace(free int, required int, del int) int {
	for _, it := range root.items {
		if it.items != nil {
			if free + it.size >= required && it.size < del {
				del = it.size
			}
			del = it.freeSpace(free, required, del)
		}
	}
	return del
}