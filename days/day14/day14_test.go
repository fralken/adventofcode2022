package day14

import "testing"

func TestFirstStar(t *testing.T) {
	content :=
`498,4 -> 498,6 -> 496,6
503,4 -> 502,4 -> 502,9 -> 494,9`
	result := firstStar(content)
	want := 24
	if result != want {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, want)
	}
}

func TestSecondStar(t *testing.T) {
	content :=
`498,4 -> 498,6 -> 496,6
503,4 -> 502,4 -> 502,9 -> 494,9`
	result := secondStar(content)
	want := 93
	if result != want {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, want)
	}
}
