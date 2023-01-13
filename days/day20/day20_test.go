package day20

import "testing"

func TestFirstStar(t *testing.T) {
	content :=
`1
2
-3
3
-2
0
4`
	result := firstStar(content)
	want := 3
	if result != want {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, want)
	}
}

func TestSecondStar(t *testing.T) {
	content :=
`1
2
-3
3
-2
0
4`
	result := secondStar(content)
	want := 1623178306
	if result != want {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, want)
	}
}
