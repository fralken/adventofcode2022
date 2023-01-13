package day09

import "testing"

func TestFirstStar(t *testing.T) {
	content :=
`R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`
	result := firstStar(content)
	want := 13
	if result != want {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, want)
	}
}

func TestSecondStar1(t *testing.T) {
	content :=
`R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`
	result := secondStar(content)
	want := 1
	if result != want {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, want)
	}
}

func TestSecondStar2(t *testing.T) {
	content :=
`R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20`
	result := secondStar(content)
	want := 36
	if result != want {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, want)
	}
}
