package day04

import (
	"errors"
	"strings"
)

func part1(input string) (int, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	sum := 0
	for _, line := range lines {
		count, err := countWinningCards(line)
		if err != nil {
			return 0, err
		}
		if count > 0 {
			sum += 1 << (count - 1)
		}
	}

	return sum, nil
}

func part2(input string) (int, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	counts := make([]int, len(lines))
	for i := range counts {
		counts[i] = 1
	}
	for current, line := range lines {
		count, err := countWinningCards(line)
		if err != nil {
			return 0, err
		}
		for i := 1; i <= count; i++ {
			counts[current+i] += counts[current]
		}
	}

	sum := 0
	for _, c := range counts {
		sum += c
	}

	return sum, nil
}

func countWinningCards(line string) (int, error) {
	index := strings.Index(line, ": ")
	if index < 0 {
		return 0, errors.New("missing colon in " + line)
	}
	winningNumbersRaw, numbersYouHaveRaw, found := strings.Cut(line[index+2:], " | ")
	if !found {
		return 0, errors.New("missing | separator in " + line)
	}
	winningNumbers := map[string]bool{}
	for _, n := range strings.Fields(winningNumbersRaw) {
		winningNumbers[n] = true
	}
	count := 0
	for _, n := range strings.Fields(numbersYouHaveRaw) {
		if winningNumbers[n] {
			count++
		}
	}
	return count, nil
}
