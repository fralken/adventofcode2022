package day11

import (
	"aoc2022/utils"
	"strings"
)

func FirstStar() {
	utils.Star(11, 1, "level of monkey business after 20 rounds", firstStar)
}

func SecondStar() {
	utils.Star(11, 2, "level of monkey business after 10000 rounds", secondStar)
}

func firstStar(content string) int {
	monkeys := parseMonkeys(content)
	return play(monkeys, 20, func(a int) int { return a / 3 })
}

func secondStar(content string) int {
	monkeys := parseMonkeys(content)
	product := monkeys[0].test
	for _, m := range monkeys[1:] {
		product *= m.test
	}
	return play(monkeys, 10000, func(a int) int { return a % product })
}

type monkey struct {
	items []int
	operation func(int) int
	test int
	throwTrue int
	throwFalse int
	count int
}

func parseMonkeys(content string) []monkey {
	parts := strings.Split(content, "\n\n")
	monkeys := make([]monkey, len(parts))
	for i, part := range parts {
		lines := strings.Split(part, "\n")
		items := parseItems(lines[1])
		operation := parseOperation(lines[2])
		test := parseTest(lines[3])
		throwTrue := parseIfTrue(lines[4])
		throwFalse := parseIfFalse(lines[5])
		monkeys[i] = monkey{
			items,
			operation,
			test,
			throwTrue,
			throwFalse,
			0,
		}
	}
	return monkeys
}

func parseItems(line string) []int {
	return utils.StringsToInts(strings.Split(line[len("  Starting items: "):], ", "))
}

func parseOperation(line string) func(int) int {
	var op func(int, int) int
	operation := line[len("  Operation: new = old ")]
	if operation == '+' {
		op = func(a, b int) int { return a + b }
	} else { // operation == '*'
		op = func(a, b int) int { return a * b }
	}
	operand := line[len("  Operation: new = old ")+2:]
	if operand == "old" {
		return func(a int) int { return op(a, a) }
	} else {
		value := utils.StringToInt(operand)
		return func(a int) int { return op(a, value) }
	}
}

func parseTest(line string) int {
	return utils.StringToInt(line[len("  Test: divisible by "):])
}

func parseIfTrue(line string) int {
	return utils.StringToInt(line[len("    If true: throw to monkey "):])
}

func parseIfFalse(line string) int {
	return utils.StringToInt(line[len("    If false: throw to monkey "):])
}

func play(monkeys []monkey, rounds int, f func(int) int) int {
	for round := 0; round < rounds; round++ {
		for i := 0; i < len(monkeys); i++ {
			for len(monkeys[i].items) > 0 {
				monkeys[i].count++
				worry := f(monkeys[i].operation(monkeys[i].items[0]))
				var next int
				if worry % monkeys[i].test == 0 {
					next = monkeys[i].throwTrue
				} else {
					next = monkeys[i].throwFalse
				}
				monkeys[next].items = append(monkeys[next].items, worry)
				monkeys[i].items = monkeys[i].items[1:]
			}
		}
	}
	max1 := 0
	max2 := 0
	for i := 0; i < len(monkeys); i++ {
		if max1 < monkeys[i].count {
			max2 = max1
			max1 = monkeys[i].count
		} else if max2 < monkeys[i].count {
			max2 = monkeys[i].count
		}
	}
	return max1 * max2
}