package day17

import "testing"

func TestFirstStar(t *testing.T) {
	content := ">>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>"
	result := firstStar(content)
	want := 3068
	if result != want {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, want)
	}
}

func TestSecondStar(t *testing.T) {
	content := ">>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>"
	result := secondStar(content)
	want := 1514285714288
	if result != want {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, want)
	}
}
