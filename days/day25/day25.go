package day25

import (
	"aoc2022/utils"
	"fmt"
	"strings"
)

func FirstStar() {
	utils.Star(25, 1, "sum of SNAFU numbers", firstStar)
}

func SecondStar() {
	utils.Star(25, 2, "THE END ...", secondStar)
}

func firstStar(content string) string {
	lines := strings.Split(content, "\n")
	sum := 0
	for _, l := range lines {
		sum += snafuToDecimal(l)
	}
	return decimalToSnafu(sum)
}

func secondStar(content string) string {
	return "Thanks for watching"
}

func snafuToDecimal(s string) int {
	val := 0
	for mul, i := 1, len(s)-1; i >= 0; mul, i = mul * 5, i - 1 {
		if s[i] == '=' {
			val += -2 * mul
		} else if s[i] == '-' {
			val += -1 * mul
		} else {
			val += utils.StringToInt(string(s[i])) * mul
		}
	}
	return val
}

func decimalToSnafu(n int) string {
	val := ""
	for n > 0 {
		r := n % 5
		n = n / 5
		if r == 4 {
			n++
			val = "-" + val
		} else if r == 3 {
			n++
			val = "=" + val
		} else {
			val = fmt.Sprint(r) + val
		}
	}
	return val
}