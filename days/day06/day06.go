package day06

import (
	"aoc2022/utils"
)

func FirstStar() {
	utils.Star(6, 1, "chars before first detected packet marker", firstStar)
}

func SecondStar() {
	utils.Star(6, 2, "chars before first detected start of msg", secondStar)
}

func firstStar(content string) int {
	return findMarker(content, 4)
}

func secondStar(content string) int {
	return findMarker(content, 14)
}

func findMarker(content string, length int) int {
	for i := 0; i < len(content) - length; i++ {
		found := true
		for j := 0; j < length-1 && found; j++ {
			for k := j+1; k < length && found; k++ {
				if content[i+j] == content[i+k] {
					found = false 
				}
			}
		}
		if found { return i + length }
	}
	return 0
}