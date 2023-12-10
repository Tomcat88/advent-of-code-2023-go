package main

import (
	"aoc23/utils"
	"fmt"
	"slices"
)

func main() {
	lines, err := utils.GetInput("day10")
	// lines, err := utils.GetInput("day10")
	if err != nil {
		panic(err)
	}
	fmt.Println("part 1:", part1(lines))
	// fmt.Println("part 2:", part2(lines))
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
	// startString := joinToString(currents[0], ",")
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

func inside(lines []string, boundaries map[[2]int]int, i, j int, boundaryCount int) bool {
	current := lines[i][j]
	_, isBoundary := boundaries[[2]int{i, j}]
	fmt.Println(current, i, j, isBoundary, boundaryCount)
	return false
}

func prettyPrint(lines []string, boundaries map[[2]int]int, outside map[[2]int]int, replace bool) {
	for i, line := range lines {
		for j, r := range line {
			if _, found := boundaries[[2]int{i, j}]; found {
				if replace {
					fmt.Print(string(printMap[r]))
				} else {
					fmt.Print(string(r))
				}
			} else if _, found := outside[[2]int{i, j}]; found {
				fmt.Print("O")
			} else {
				fmt.Print("X")
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

var squeezeAllowedPipes = map[[2]int][]rune{
	Right: []rune{'-', 'J', '7'},
	Left:  []rune{'-', 'F', 'L'},
	Up:    []rune{'|', 'F', '7'},
	Down:  []rune{'|', 'J', 'L'},
}

func allowedOutsiseMoves(lines []string, boundaries map[[2]int]int, i, j int) (allowed [][2]int) {
	// currentPosition := lines[i][j]
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
		if boundaries[[2]int{newI, newJ}] > 0 {
			continue
		}
		allowed = append(allowed, move)
	}
	return
}

var allowedPipes = map[[2]int][]rune{
	Right: []rune{'-', 'J', '7'},
	Left:  []rune{'-', 'F', 'L'},
	Up:    []rune{'|', 'F', '7'},
	Down:  []rune{'|', 'J', 'L'},
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
				// fmt.Println(string(currentPosition), string(newPosition))
				allowed = append(allowed, move)
			}
		case '-':
			if (move == Right || move == Left) && slices.Contains(allowedPipes[move], newPosition) {
				// fmt.Println(string(currentPosition), string(newPosition))
				allowed = append(allowed, move)
			}
		case '|':
			if (move == Up || move == Down) && slices.Contains(allowedPipes[move], newPosition) {
				// fmt.Println(string(currentPosition), string(newPosition))
				allowed = append(allowed, move)
			}
		case 'J':
			if (move == Up || move == Left) && slices.Contains(allowedPipes[move], newPosition) {
				// fmt.Println(string(currentPosition), string(newPosition))
				allowed = append(allowed, move)
			}
		case '7':
			if (move == Down || move == Left) && slices.Contains(allowedPipes[move], newPosition) {
				// fmt.Println(string(currentPosition), string(newPosition))
				allowed = append(allowed, move)
			}
		case 'F':
			if (move == Down || move == Right) && slices.Contains(allowedPipes[move], newPosition) {
				// fmt.Println(string(currentPosition), string(newPosition))
				allowed = append(allowed, move)
			}
		case 'L':
			if (move == Up || move == Right) && slices.Contains(allowedPipes[move], newPosition) {
				// fmt.Println(string(currentPosition), string(newPosition))
				allowed = append(allowed, move)
			}
		}

	}
	return allowed
}

/* func joinToString(ints [2]int, sep string) string {
	slice := make([]string, len(ints))
	for i, v := range ints {
		slice[i] = strconv.Itoa(v)
	}
	return strings.Join(slice, sep)
} */

func part2(lines []string) int {
	var startI, startJ int
	for i, line := range lines {
		for j, r := range line {
			if r == 'S' {
				fmt.Println("Starting from", i, j)
				startI, startJ = i, j
			}
		}
	}
	currents := [][2]int{
		{startI, startJ},
	}
	// startString := joinToString(currents[0], ",")
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
	fmt.Println("visited", visited)
	outsideVisited := map[[2]int]int{{0, 0}: 1}
	for i, line := range lines {
		for j, _ := range line {
			boundaryCount := 0
			if inside(lines, visited, i, j, boundaryCount) {
				outsideVisited[[2]int{i, j}] = 1
			}
		}
	}
	prettyPrint(lines, visited, outsideVisited, true)
	/* boundaries := utils.Keys(visited)
	 for len(currents) > 0 {
		for _, current := range currents {
			i, j := current[0], current[1]
			if slices.Contains(boundaries, [2]int{i, j}) {
				continue
			}

		}

	} */
	return 0
}
