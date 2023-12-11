package main

import (
	"aoc23/utils"
	"fmt"
	"math"
	"slices"
)

type Pair [2]int

const EXPANSION_FACTOR = 1000000

func main() {
	lines, err := utils.GetInput("day11")
	if err != nil {
		panic(err)
	}
	fmt.Println("part 1:", part1(lines))
	fmt.Println("part 2:", part2(lines))
}

func expandGalaxy(galaxy []string) (expanded []string, expansionI []int, expansionJ []int) {
	for i := 0; i < len(galaxy); i++ {
		expandRow := true
		row := ""
		for j := 0; j < len(galaxy[0]); j++ {
			expandRow = expandRow && galaxy[i][j] == '.'
			expandColumn := true
			for k := 0; k < len(galaxy); k++ {
				expandColumn = expandColumn && galaxy[k][j] == '.'
			}
			if expandColumn {
				row = row + "."
				if !slices.Contains(expansionJ, j) {
					expansionJ = append(expansionJ, j)
				}
			}
			row += string(galaxy[i][j])
		}
		if expandRow {
			expanded = append(expanded, row)
			expansionI = append(expansionI, i)
		}
		expanded = append(expanded, row)
	}
	return
}

func part1(lines []string) (sum int) {
	expanded, _, _ := expandGalaxy(lines)
	galaxies := make([]Pair, 0)
	for i, line := range expanded {
		for j, r := range line {
			if r == '#' {
				galaxies = append(galaxies, Pair{i, j})
			}
		}
	}
	cmbs := combinations(galaxies)
	for _, c := range cmbs {
		d := manhattanDistance(c[0], c[1])
		sum += d
	}
	return
}

func part2(lines []string) (sum int) {
	_, expansionI, expansionJ := expandGalaxy(lines)
	galaxies := make([]Pair, 0)
	for i, line := range lines {
		for j, r := range line {
			if r == '#' {
				galaxies = append(galaxies, Pair{i, j})
			}
		}
	}
	cmbs := combinations(galaxies)
	for _, c := range cmbs {
		expanded1 := expandPair(c[0], expansionI, expansionJ)
		expanded2 := expandPair(c[1], expansionI, expansionJ)
		d := manhattanDistance(expanded1, expanded2)
		sum += d
	}
	return
}

func expandPair(pair Pair, expansionI, expansionJ []int) Pair {
	iFactor := 0
	for _, i := range expansionI {
		if i < pair[0] {
			iFactor++
		}
	}
	jFactor := 0
	for _, j := range expansionJ {
		if j < pair[1] {
			jFactor++
		}
	}
	return Pair{pair[0] + ((EXPANSION_FACTOR - 1) * iFactor), pair[1] + ((EXPANSION_FACTOR - 1) * jFactor)}
}

func combinations(pairs []Pair) (combinations [][2]Pair) {
	for i := 0; i < len(pairs); i++ {
		for j := i + 1; j < len(pairs); j++ {
			combinations = append(combinations, [2]Pair{pairs[i], pairs[j]})
		}
	}
	return
}

func manhattanDistance(galaxy1, galaxy2 Pair) int {
	x1, y1 := galaxy1[0], galaxy1[1]
	x2, y2 := galaxy2[0], galaxy2[1]

	return int(math.Abs(float64(x1-x2)) + math.Abs(float64(y1-y2)))
}
