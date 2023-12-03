package day03

import "testing"

func TestDay(t *testing.T) {
	t.Run("test part 1", func(t *testing.T) {
		input := "467..114..\n...*......\n..35..633.\n......#...\n617*......\n.....+.58.\n..592.....\n......755.\n...$.*....\n.664.598.."
		got := part1(generate(input))
		want := 4361

		if got != want {
			t.Errorf("error, got %d want %d", got, want)
		}
	})
	t.Run("test part 2", func(t *testing.T) {
		input := "467..114..\n...*......\n..35..633.\n......#...\n617*......\n.....+.58.\n..592.....\n......755.\n...$.*....\n.664.598.."
		got := part2(generate(input))
		want := 467835

		if got != want {
			t.Errorf("error, got %d want %d", got, want)
		}
	})
}

func TestHelpers(t *testing.T) {
	t.Run("Check adjacent cells for symbols", func(t *testing.T) {
		input := generate("...\n.67\n.$.")

		got := hasAdjacentSymbol(1, 1, input)

		if !got {
			t.Errorf("error, got %t want %t", got, !got)
		}
	})
	t.Run("Get value for part", func(t *testing.T) {
		input := []rune{'.', '.', '6', '7'}

		got, err := getPartValue(2, input)
		if err != nil {
			t.Fatal("error", err)
		}
		want := 67

		if got != want {
			t.Errorf("error, got %d want %d", got, want)
		}
	})
	t.Run("Get gear ratio for part", func(t *testing.T) {
		input := generate("467..114..\n...*......\n..35..633.\n......#...\n617*......\n.....+.58.\n..592.....\n......755.\n...$.*....\n.664.598..")

		got := getGearRatio(1, 3, input)
		want := 16345

		if got != want {
			t.Errorf("error, got %d want %d", got, want)
		}
	})
}
