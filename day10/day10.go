package main

import (
	"aoc23/utils"
	"fmt"
	"slices"
)

func main() {
	lines, err := utils.GetInput("day10")
	if err != nil {
		panic(err)
	}
	fmt.Println("part 1:", part1(lines))
	fmt.Println("part 2:", part2(lines))
}

func part1(lines []string) int {
	var startI, startJ int
	for i, line := range lines {
		for j, r := range line {
			if r == 'S' {
				startI, startJ = i, j
			}
		}
	}
	currents := [][2]int{
		{startI, startJ},
	}
	visited := map[[2]int]int{currents[0]: 0}
	steps := 0
	for len(currents) > 0 {
		newCurrents := make([][2]int, 0)
		for _, current := range currents {
			i, j := current[0], current[1]
			visited[current] = steps
			moves := allowedMoves(lines, i, j)
			for _, move := range moves {
				nextI, nextJ := i+move[0], j+move[1]
				if _, found := visited[[2]int{nextI, nextJ}]; !found {
					newCurrents = append(newCurrents, [2]int{nextI, nextJ})
					// visited[] = 1
				}
			}
		}
		currents = newCurrents
		steps++
	}
	max := -1
	for _, v := range visited {
		if v > max {
			max = v
		}
	}
	return max
}

func prettyPrint(lines []string, boundaries map[[2]int]int, inside map[[2]int]int, replace bool) {
	for i, line := range lines {
		for j, r := range line {
			if _, found := boundaries[[2]int{i, j}]; found {
				if replace {
					fmt.Print(string(printMap[r]))
				} else {
					fmt.Print(string(r))
				}
			} else if _, found := inside[[2]int{i, j}]; found {
				fmt.Print("I")
			} else {
				fmt.Print("O")
			}
		}
		fmt.Println()
	}
}

var printMap = map[rune]rune{
	'S': 'S',
	'.': ' ',
	'|': '┃',
	'-': '━',
	'J': '┛',
	'F': '┏',
	'7': '┓',
	'L': '┗',
}

var (
	Right = [2]int{0, 1}
	Left  = [2]int{0, -1}
	Up    = [2]int{-1, 0}
	Down  = [2]int{1, 0}
)

var moves = [][2]int{
	{0, -1}, // Left
	{0, 1},  // Right
	{-1, 0}, // Up
	{1, 0},  // Down
}

var allowedPipes = map[[2]int][]rune{
	Right: {'-', 'J', '7'},
	Left:  {'-', 'F', 'L'},
	Up:    {'|', 'F', '7'},
	Down:  {'|', 'J', 'L'},
}

func allowedMoves(lines []string, i, j int) (allowed [][2]int) {
	currentPosition := lines[i][j]
	for _, move := range moves {
		newI, newJ := i+move[0], j+move[1]
		if newI < 0 || newJ < 0 {
			continue
		}
		if newI >= len(lines) {
			continue
		}
		if newJ >= len(lines[0]) { // Every line should have the same length
			continue
		}
		newPosition := rune(lines[newI][newJ])
		switch currentPosition {
		case 'S': // Allowed: Right: - J 7 Left: - F L Up: | F 7 Down: | J L
			if slices.Contains(allowedPipes[move], newPosition) {
				allowed = append(allowed, move)
			}
		case '-':
			if (move == Right || move == Left) && slices.Contains(allowedPipes[move], newPosition) {
				allowed = append(allowed, move)
			}
		case '|':
			if (move == Up || move == Down) && slices.Contains(allowedPipes[move], newPosition) {
				allowed = append(allowed, move)
			}
		case 'J':
			if (move == Up || move == Left) && slices.Contains(allowedPipes[move], newPosition) {
				allowed = append(allowed, move)
			}
		case '7':
			if (move == Down || move == Left) && slices.Contains(allowedPipes[move], newPosition) {
				allowed = append(allowed, move)
			}
		case 'F':
			if (move == Down || move == Right) && slices.Contains(allowedPipes[move], newPosition) {
				allowed = append(allowed, move)
			}
		case 'L':
			if (move == Up || move == Right) && slices.Contains(allowedPipes[move], newPosition) {
				allowed = append(allowed, move)
			}
		}

	}
	return allowed
}

func part2(lines []string) int {
	var startI, startJ int
	for i, line := range lines {
		for j, r := range line {
			if r == 'S' {
				startI, startJ = i, j
			}
		}
	}
	currents := [][2]int{
		{startI, startJ},
	}
	visited := map[[2]int]int{currents[0]: 0}
	steps := 0
	for len(currents) > 0 {
		newCurrents := make([][2]int, 0)
		for _, current := range currents {
			i, j := current[0], current[1]
			visited[current] = steps
			moves := allowedMoves(lines, i, j)
			for _, move := range moves {
				nextI, nextJ := i+move[0], j+move[1]
				if _, found := visited[[2]int{nextI, nextJ}]; !found {
					newCurrents = append(newCurrents, [2]int{nextI, nextJ})
				}
			}
		}
		currents = newCurrents
		steps++
	}
	inside := map[[2]int]int{}
	// prettyPrint(lines, visited, inside, true)
	for i, line := range lines {
		in := false
		for j, rune := range line {
			if _, found := visited[[2]int{i, j}]; found {
				if rune == '|' || rune == 'L' || rune == 'J' {
                    in = !in
                }
			} else {
				if in {
					inside[[2]int{i, j}] = 1
				}
			}
		}
	}
	// prettyPrint(lines, visited, inside, true)
	return len(inside)
}
