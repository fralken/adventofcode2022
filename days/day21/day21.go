package day21

import (
	"aoc2022/utils"
	"strconv"
	"strings"
)

func FirstStar() {
	utils.Star(21, 1, "number that root monkey yells", firstStar)
}

func SecondStar() {
	utils.Star(21, 2, "number that human yells", secondStar)
}

func firstStar(content string) int {
	monkeys := parseMonkeys(content)
	return monkeys.getNumber("root")
}

func secondStar(content string) int {
	monkeys := parseMonkeys(content)
	root := monkeys["root"]
	if monkeys.findHuman(*root.left) {
		num := monkeys.getNumber(*root.right)
		return monkeys.findNumber(*root.left, num)
	} else { // human is in right branch
		num := monkeys.getNumber(*root.left)
		return monkeys.findNumber(*root.right, num)
	}
}

type monkey struct {
	num *int
	left, right *string
	op *string
}

type monkeys map[string]monkey

func parseMonkeys(content string) monkeys {
	lines := strings.Split(content, "\n")
	monkeys := make(monkeys)
	for _, line := range lines {
		parts := strings.Split(line, ": ")
		name := parts[0]
		num, ok := strconv.Atoi(parts[1])
		if ok == nil {
			monkeys[name] = monkey{ &num, nil, nil, nil }
		} else {
			op := strings.Split(parts[1], " ")
			monkeys[name] = monkey{ nil, &op[0], &op[2], &op[1] }
		}
	}
	return monkeys
}

func (ms monkeys) getNumber(name string) int {
	m := ms[name]
	if m.num != nil {
		return *m.num
	} else {
		return op(*m.op, ms.getNumber(*m.left), ms.getNumber(*m.right))
	}
}

func (ms monkeys) findHuman(name string) bool {
	if name == "humn" {
		return true
	}
	m := ms[name]
	if m.num != nil {
		return false
	} else {
		return ms.findHuman(*m.left) || ms.findHuman(*m.right)
	}
}

func (ms monkeys) findNumber(name string, match int) int {
	if ms[name].num != nil {
		return match
	} else {
		m := ms[name]
		if ms.findHuman(*m.left) {
			num := ms.getNumber(*m.right)
			match = reverseLeftOp(*m.op, match, num)
			return ms.findNumber(*m.left, match)
		} else { // human is in right branch
			num := ms.getNumber(*m.left)
			match = reverseRightOp(*m.op, match, num)
			return ms.findNumber(*m.right, match)
		}
	}
}

func op(op string, a, b int) int {
	switch op {
	case "+": return a + b
	case "-": return a - b
	case "*": return a * b
	case "/": return a / b
	}
	return 0
}

// reverse operation when number to find is on the left
// if "a" is the number to match, "x" is the number to find
// and "b" is the other operand, we have
// x + b = a --> x = a - b
// x - b = a --> x = a + b
// x * b = a --> x = a / b
// x / b = a --> x = a * b
// the result is the new number to match
func reverseLeftOp(op string, a, b int) int {
	switch op {
	case "+": return a - b
	case "-": return a + b
	case "*": return a / b
	case "/": return a * b
	}
	return 0
}

// reverse operation when number to find is on the right
// if "a" is the number to match, "x" is the number to find
// and "b" is the other operand, we have
// b + x = a --> x = a - b
// b - x = a --> x = b - a
// b * x = a --> x = a / b
// b / x = a --> x = b / a
// the result is the new number to match
func reverseRightOp(op string, a, b int) int {
	switch op {
	case "+": return a - b
	case "-": return b - a
	case "*": return a / b
	case "/": return b / a
	}
	return 0
}

