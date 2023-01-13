package day08

import "testing"

func TestFirstStar(t *testing.T) {
	content :=
`30373
25512
65332
33549
35390`
	result := firstStar(content)
	want := 21
	if result != want {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, want)
	}
}

func TestSecondStar(t *testing.T) {
	content :=
`30373
25512
65332
33549
35390`
	result := secondStar(content)
	want := 8
	if result != want {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, want)
	}
}
