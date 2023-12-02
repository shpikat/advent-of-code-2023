package day01

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

func part1(input string) (int, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	sum := 0
	for _, line := range lines {
		first := strings.IndexFunc(line, unicode.IsDigit)
		last := strings.LastIndexFunc(line, unicode.IsDigit)
		sum += int(line[first]-'0')*10 + int(line[last]-'0')
	}

	return sum, nil
}

func part2(input string) (int, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	var digits = map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	var sb strings.Builder
	sb.WriteByte('(')
	for digit := range digits {
		sb.WriteString(digit)
		sb.WriteByte('|')
	}
	sb.WriteString("[0-9])")
	regex := sb.String()

	// add greedy search with beginning/end anchoring to reduce iterations and result in a single capturing group
	forth, err := regexp.Compile(regex + ".*$")
	if err != nil {
		return 0, err
	}
	back, err := regexp.Compile("^.*" + regex)
	if err != nil {
		return 0, err
	}

	sum := 0
	for _, line := range lines {
		found := forth.FindStringSubmatch(line)[1]
		first, ok := digits[found]
		if !ok {
			first = int(found[0] - '0')
		}

		fmt.Println(line)
		found = back.FindStringSubmatch(line)[1]

		last, ok := digits[found]
		if !ok {
			last = int(found[0] - '0')
		}

		sum += first*10 + last
	}
	return sum, nil
}
