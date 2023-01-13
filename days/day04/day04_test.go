package day04

import "testing"

func TestFirstStar(t *testing.T) {
	content :=
	`2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`
	result := firstStar(content)
	want := 2
	if result != want {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, want)
	}
}

func TestSecondStar(t *testing.T) {
	content := 
	`2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`
	result := secondStar(content)
	want := 4
	if result != want {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, want)
	}
}
