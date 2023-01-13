package day06

import "testing"

func TestFirstStar(t *testing.T) {
	content := []string{
		"mjqjpqmgbljsphdztnvjfqwrcgsmlb",
		"bvwbjplbgvbhsrlpgdmjqwftvncz",
		"nppdvjthqldpwncqszvftbrmjlhg",
		"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
		"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
	}
	want := []int{ 7, 5, 6, 10, 11 }
	for i := range content {
		result := firstStar(content[i])
		if result != want[i] {
			t.Errorf("Result was incorrect, got: %d, want: %d.", result, want)
		}
	}
}

func TestSecondStar(t *testing.T) {
	content := []string{
		"mjqjpqmgbljsphdztnvjfqwrcgsmlb",
		"bvwbjplbgvbhsrlpgdmjqwftvncz",
		"nppdvjthqldpwncqszvftbrmjlhg",
		"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
		"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
	}
	want := []int{ 19, 23, 23, 29, 26}
	for i := range content {
		result := secondStar(content[i])
		if result != want[i] {
			t.Errorf("Result was incorrect, got: %d, want: %d.", result, want)
		}
	}
}
