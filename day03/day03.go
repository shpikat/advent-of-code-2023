package day03

import (
	"strings"
)

const EMPTY = '.'

func part1(input string) (int, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	sum := 0
	for i := range lines {
		line := lines[i]
		var upper, lower string
		if i > 0 {
			upper = lines[i-1]
		} else {
			upper = createEmptyLine(len(line))
		}
		if i < len(lines)-1 {
			lower = lines[i+1]
		} else {
			lower = createEmptyLine(len(line))
		}

		acc := 0
		isAdjacent := false
		for j := range line {
			ch := line[j]

			hasFound := isSymbol(ch) || isSymbol(upper[j]) || isSymbol(lower[j])

			if isDigit(ch) {
				acc = acc*10 + int(ch-'0')
				isAdjacent = isAdjacent || hasFound
			} else {
				if acc != 0 {
					if isAdjacent || hasFound {
						sum += acc
					}
					acc = 0
				}
				isAdjacent = hasFound
			}
		}
		if acc != 0 && isAdjacent {
			sum += acc
		}
	}

	return sum, nil
}

func part2(input string) (int, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	type Coordinate struct {
		x, y int
	}
	const GEAR = '*'
	NIL := Coordinate{-1, -1}

	gears := map[Coordinate][]int{}
	for i := range lines {
		line := lines[i]
		var upper, lower string
		if i > 0 {
			upper = lines[i-1]
		} else {
			upper = createEmptyLine(len(line))
		}
		if i < len(lines)-1 {
			lower = lines[i+1]
		} else {
			lower = createEmptyLine(len(line))
		}

		acc := 0
		// Assuming only one gear can be adjacent to a number
		var gear Coordinate
		for j := range line {
			ch := line[j]

			var currentGear Coordinate
			if ch == GEAR {
				currentGear = Coordinate{j, i}
			} else if upper[j] == GEAR {
				currentGear = Coordinate{j, i - 1}
			} else if lower[j] == GEAR {
				currentGear = Coordinate{j, i + 1}
			} else {
				currentGear = NIL
			}

			if currentGear != NIL {
				gear = currentGear
			}

			if isDigit(ch) {
				acc = acc*10 + int(ch-'0')
			} else {
				if acc != 0 {
					if gear != NIL {
						gears[gear] = append(gears[gear], acc)
					}
					acc = 0
				}
				gear = currentGear
			}
		}
		if acc != 0 && gear != NIL {
			gears[gear] = append(gears[gear], acc)
		}
	}

	sum := 0
	for _, numbers := range gears {
		if len(numbers) == 2 {
			sum += numbers[0] * numbers[1]
		}
	}

	return sum, nil
}

func createEmptyLine(size int) string {
	return strings.Repeat(string(EMPTY), size)
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func isSymbol(ch byte) bool {
	return ch != EMPTY && !isDigit(ch)
}
