package day22

import "testing"

func TestFirstStar(t *testing.T) {
	content :=
`        ...#
        .#..
        #...
        ....
...#.......#
........#...
..#....#....
..........#.
        ...#....
        .....#..
        .#......
        ......#.

10R5L5R10L4R5L5`
	result := firstStar(content)
	want := 6032
	if result != want {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, want)
	}
}

func TestSecondStar(t *testing.T) {
	content :=
`        ...#
        .#..
        #...
        ....
...#.......#
........#...
..#....#....
..........#.
        ...#....
        .....#..
        .#......
        ......#.

10R5L5R10L4R5L5`
	result := secondStar(content)
	want := 5031
	if result != want {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, want)
	}
}
