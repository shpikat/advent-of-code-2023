package day02

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

func part1(input string) (int, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	expected := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	regex, err := regexp.Compile("^Game (\\d+): (.+)$")
	if err != nil {
		return 0, err
	}
	sum := 0

games:
	for _, line := range lines {
		matches := regex.FindStringSubmatch(line)

		subsets := strings.Split(matches[2], "; ")
		for _, subset := range subsets {
			cubes := strings.Split(subset, ", ")
			for _, cube := range cubes {
				number, colour, found := strings.Cut(cube, " ")
				if !found {
					return 0, errors.New("unexpected format for a line: " + line)
				}
				max, ok := expected[colour]
				if ok {
					n, err := strconv.Atoi(number)
					if err != nil {
						return 0, err
					}

					if n > max {
						continue games
					}
				}
			}
		}

		id, err := strconv.Atoi(matches[1])
		if err != nil {
			return 0, err
		}

		sum += id
	}

	return sum, nil
}

func part2(input string) (int, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	sum := 0

	for _, line := range lines {
		_, game, found := strings.Cut(line, ": ")
		if !found {
			return 0, errors.New("unexpected format for a line: " + line)
		}

		counts := map[string]int{}

		subsets := strings.Split(game, "; ")
		for _, subset := range subsets {
			cubes := strings.Split(subset, ", ")
			for _, cube := range cubes {
				number, colour, found := strings.Cut(cube, " ")
				if !found {
					return 0, errors.New("unexpected format for a line: " + line)
				}
				n, err := strconv.Atoi(number)
				if err != nil {
					return 0, err
				}

				if n > counts[colour] {
					counts[colour] = n
				}
			}
		}

		sum += counts["red"] * counts["green"] * counts["blue"]
	}

	return sum, nil
}
