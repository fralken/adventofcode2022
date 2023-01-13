package day23

import "testing"

func TestFirstStar(t *testing.T) {
	content :=
`....#..
..###.#
#...#.#
.#...##
#.###..
##.#.##
.#..#..`
	result := firstStar(content)
	want := 110
	if result != want {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, want)
	}
}

func TestSecondStar(t *testing.T) {
	content :=
`....#..
..###.#
#...#.#
.#...##
#.###..
##.#.##
.#..#..`
	result := secondStar(content)
	want := 20
	if result != want {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, want)
	}
}
