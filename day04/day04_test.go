package day04

import (
	"testing"

	"github.com/shpikat/advent-of-code-2023/internal"
)

const (
	sample1 = `
Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11
`

	part1Sample = 13
	part1Answer = 24542

	part2Sample = 30
	part2Answer = 8736438
)

func TestPart1(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		want  int
	}{
		{"sample 1", sample1, part1Sample},
		{"puzzle input", internal.ReadInput(t, "./testdata/input.txt"), part1Answer},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := part1(tc.input)
			if err != nil {
				t.Errorf("Error: %v", err)
			}
			if got != tc.want {
				t.Errorf("Got: %v, want: %v", got, tc.want)
			}
		})
	}

}

func TestPart2(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		want  int
	}{
		{"sample 2", sample1, part2Sample},
		{"puzzle input", internal.ReadInput(t, "./testdata/input.txt"), part2Answer},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := part2(tc.input)
			if err != nil {
				t.Errorf("Error: %v", err)
			}
			if got != tc.want {
				t.Errorf("Got: %v, want: %v", got, tc.want)
			}
		})
	}
}

func Benchmark(b *testing.B) {
	input := internal.ReadInput(b, "./testdata/input.txt")
	parts := []struct {
		name   string
		fn     func(input string) (int, error)
		answer int
	}{
		{"part1", part1, part1Answer},
		{"part2", part2, part2Answer},
	}

	for _, part := range parts {
		b.Run(part.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				got, err := part.fn(input)
				if err != nil {
					b.Errorf("Error: %v", err)
				}
				if got != part.answer {
					b.Errorf("Got: %v, want: %v", got, part.answer)
				}
			}
		})
	}
}
