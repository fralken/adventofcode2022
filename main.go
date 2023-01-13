package main

import (
	"aoc2022/days/day01"
	"aoc2022/days/day02"
	"aoc2022/days/day03"
	"aoc2022/days/day04"
	"aoc2022/days/day05"
	"aoc2022/days/day06"
	"aoc2022/days/day07"
	"aoc2022/days/day08"
	"aoc2022/days/day09"
	"aoc2022/days/day10"
	"aoc2022/days/day11"
	"aoc2022/days/day12"
	"aoc2022/days/day13"
	"aoc2022/days/day14"
	"aoc2022/days/day15"
	"aoc2022/days/day16"
	"aoc2022/days/day17"
	"aoc2022/days/day18"
	"aoc2022/days/day19"
	"aoc2022/days/day20"
	"aoc2022/days/day21"
	"aoc2022/days/day22"
	"aoc2022/days/day23"
	"aoc2022/days/day24"
	"aoc2022/days/day25"
	"fmt"
	"os"
	"strconv"
)

func main() {
	days := [][]func(){
		{day01.FirstStar, day01.SecondStar},
		{day02.FirstStar, day02.SecondStar},
		{day03.FirstStar, day03.SecondStar},
		{day04.FirstStar, day04.SecondStar},
		{day05.FirstStar, day05.SecondStar},
		{day06.FirstStar, day06.SecondStar},
		{day07.FirstStar, day07.SecondStar},
		{day08.FirstStar, day08.SecondStar},
		{day09.FirstStar, day09.SecondStar},
		{day10.FirstStar, day10.SecondStar},
		{day11.FirstStar, day11.SecondStar},
		{day12.FirstStar, day12.SecondStar},
		{day13.FirstStar, day13.SecondStar},
		{day14.FirstStar, day14.SecondStar},
		{day15.FirstStar, day15.SecondStar},
		{day16.FirstStar, day16.SecondStar},
		{day17.FirstStar, day17.SecondStar},
		{day18.FirstStar, day18.SecondStar},
		{day19.FirstStar, day19.SecondStar},
		{day20.FirstStar, day20.SecondStar},
		{day21.FirstStar, day21.SecondStar},
		{day22.FirstStar, day22.SecondStar},
		{day23.FirstStar, day23.SecondStar},
		{day24.FirstStar, day24.SecondStar},
		{day25.FirstStar, day25.SecondStar},
	}

	if len(os.Args) > 1 {
		day, err := strconv.Atoi(os.Args[1])
		if err != nil || day > len(days) {
			fmt.Printf("you must enter a number between 1 and %d\n", len(days))
			os.Exit(1)
		}
		day--
		if len(os.Args) > 2 {
			star, err := strconv.Atoi(os.Args[2])
			if err != nil || star > len(days[day]) {
				fmt.Printf("you must enter a number between 1 and %d\n", len(days[day]))
				os.Exit(1)
			}
			star--
			days[day][star]()
		} else {
			for _, star := range days[day] {
				star()
			}
		}
	} else {
		for _, day := range days {
			for _, star := range day {
				star()
			}
		}
	}
}
