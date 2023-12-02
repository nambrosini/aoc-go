package day01

import "testing"

func TestDay(t *testing.T) {
	t.Run("test part 1", func(t *testing.T) {
		input := "1abc2\npqr3stu8vwx\na1b2c3d4e5f\ntreb7uchet"
		got := part1(input)
		want := 142

		if got != want {
			t.Errorf("error, got %d want %d", got, want)
		}
	})
	t.Run("test part 2", func(t *testing.T) {
		input := "two1nine\neightwothree\nabcone2threexyz\nxtwone3four\n4nineeightseven2\nzoneight234\n7pqrstsixteen"
		got := part2(input)
		want := 281

		if got != want {
			t.Errorf("error, got %d want %d", got, want)
		}
	})
}
