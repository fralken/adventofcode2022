package day12

import "testing"

func TestFirstStar(t *testing.T) {
	content :=
`Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi`
	result := firstStar(content)
	want := 31
	if result != want {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, want)
	}
}

func TestSecondStar(t *testing.T) {
	content :=
`Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi`
	result := secondStar(content)
	want := 29
	if result != want {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, want)
	}
}
