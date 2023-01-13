package day25

import "testing"

func TestFirstStar1(t *testing.T) {
	content := []string{
		"1=-0-2",
		"12111",
		"2=0=",
		"21",
		"2=01",
		"111",
		"20012",
		"112",
		"1=-1=",
		"1-12",
		"12",
		"1=",
		"122",
	}
	want := []int{ 1747, 906, 198, 11, 201, 31, 1257, 32, 353, 107, 7, 3, 37 }
	for i := range content {
		result := snafuToDecimal(content[i])
		if result != want[i] {
			t.Errorf("Result was incorrect, got: %d, want: %d.", result, want)
		}
	}
}

func TestFirstStar2(t *testing.T) {
	content := []int{ 1747, 906, 198, 11, 201, 31, 1257, 32, 353, 107, 7, 3, 37 }
	want := []string{
		"1=-0-2",
		"12111",
		"2=0=",
		"21",
		"2=01",
		"111",
		"20012",
		"112",
		"1=-1=",
		"1-12",
		"12",
		"1=",
		"122",
	}
	for i := range content {
		result := decimalToSnafu(content[i])
		if result != want[i] {
			t.Errorf("Result was incorrect, got: %s, want: %s.", result, want)
		}
	}
}

func TestFirstStar3(t *testing.T) {
	content :=
`1=-0-2
12111
2=0=
21
2=01
111
20012
112
1=-1=
1-12
12
1=
122`
	result := firstStar(content)
	want := "2=-1=0"
	if result != want {
		t.Errorf("Result was incorrect, got: %s, want: %s.", result, want)
	}
}