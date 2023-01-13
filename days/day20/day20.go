package day20

import (
	"aoc2022/utils"
	"strings"
)

func FirstStar() {
	utils.Star(20, 1, "sum of the three numbers", firstStar)
}

func SecondStar() {
	utils.Star(20, 2, "sum of the three decrypted numbers", secondStar)
}

func firstStar(content string) int {
	nums := parseNumbers(content)
	mixed := nums.mix(1)
	p := mixed.indexOf(0)
	return mixed.sumOfCoords(p)
}

func secondStar(content string) int {
	nums := parseNumbers(content)
	for i := range nums {
		nums[i] *= 811589153
	}
	mixed := nums.mix(10)
	p := mixed.indexOf(0)
	return mixed.sumOfCoords(p)
}

type numbers []int

func parseNumbers(content string) numbers {
	lines := strings.Split(content, "\n")
	return utils.StringsToInts(lines)
}

func (nums numbers) mix(times int) numbers {
	len := len(nums)
	indices := make(map[int]int, len)
	// nums can have repetitions so I keep a map
	// of (original indices -> new indices) to reference them
	for i := range nums { indices[i] = i }
	for t := 0; t < times; t++ {
		for i, n := range nums {
			if n != 0 {
				pos := indices[i]
				// I remove the element from the list
				// so the new position must be modulo (len - 1)
				// because now the list is one element shorter
				newPos := utils.Mod(pos + n, len - 1)
				indices[i] = newPos
				inc := 1 // if newPos < pos move other elements forward
				if newPos > pos { inc = -1 } // move other elements backward
				for k, v := range indices {
					if k != i && (newPos <= v && v < pos || pos < v && v <= newPos) {
						indices[k] += inc
					}
				}
			}
		}
	}
	mixed := make([]int, len)
	for i, n := range nums {
		mixed[indices[i]] = n
	}
	return mixed
}

func (nums numbers) indexOf(m int) (p int) {
	for i, n := range nums {
		if n == m {
			p = i
			break
		}
	}
	return
}

func (nums numbers) sumOfCoords(p int) (sum int) {
	for q := 1000; q <= 3000; q += 1000 {
		sum += nums[utils.Mod(p + q, len(nums))]	
	}
	return
}
