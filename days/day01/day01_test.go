package day01

import "testing"

func TestFirstStar(t *testing.T) {
	content := `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`
	result := firstStar(content)
	want := 24000
	if result != want {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, want)
	}
}

func TestSecondStar(t *testing.T) {
	content := `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`
	result := secondStar(content)
	want := 45000
	if result != want {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, want)
	}
}
