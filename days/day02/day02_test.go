package day02

import "testing"

func TestFirstStar(t *testing.T) {
	content := 
	`A Y
B X
C Z`
	result := firstStar(content)
	want := 15
	if result != want {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, want)
	}
}

func TestSecondStar(t *testing.T) {
	content := 
	`A Y
B X
C Z`
	result := secondStar(content)
	want := 12
	if result != want {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, want)
	}
}
