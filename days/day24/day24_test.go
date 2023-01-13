package day24

import "testing"

func TestFirstStar(t *testing.T) {
	content :=
`#.######
#>>.<^<#
#.<..<<#
#>v.><>#
#<^v^^>#
######.#`
	result := firstStar(content)
	want := 18
	if result != want {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, want)
	}
}

func TestSecondStar(t *testing.T) {
	content :=
`#.######
#>>.<^<#
#.<..<<#
#>v.><>#
#<^v^^>#
######.#`
	result := secondStar(content)
	want := 54
	if result != want {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, want)
	}
}
