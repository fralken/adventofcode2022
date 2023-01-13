package day05

import "testing"

func TestFirstStar(t *testing.T) {
	content :=
`    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`
	result := firstStar(content)
	want := "CMZ"
	if result != want {
		t.Errorf("Result was incorrect, got: %s, want: %s.", result, want)
	}
}

func TestSecondStar(t *testing.T) {
	content :=
`    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`
	result := secondStar(content)
	want := "MCD"
	if result != want {
		t.Errorf("Result was incorrect, got: %s, want: %s.", result, want)
	}
}
