package utils

import (
	"fmt"
	"os"
	"strconv"
)

func ReadFile(filePath string) string {
	data, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func StringsToInts(strs []string) []int {
	ints := make([]int, len(strs))
	for i, str := range strs {
		ints[i], _ = strconv.Atoi(str)
	}
	return ints
}

func StringToInt(str string) int {
	i, _ := strconv.Atoi(str)
	return i
}

func Star[T int|string](day int, star int, msg string, run func(string) T) {
	content := ReadFile(fmt.Sprintf("./input/day%02d.txt", day))
	fmt.Printf("day %2d.%d - %s: %v\n", day, star, msg, run(content))
}

// euclidean modulus
func Mod(n, m int) int {
	if n < 0 {
		return (n % m) + m
	} else {
		return n % m
	}
}

func Abs(a int) int {
	if a < 0 { return -a } else { return a }
}

func Max(a, b int) int {
	if a > b { return a } else { return b }
}