package main

import (
	"aoc23/utils"
	"fmt"
)

func main() {
	lines, err := utils.GetInput("day14")
	if err != nil {
		panic(err)
	}
	fmt.Println("part 1", part1(lines))
	fmt.Println("part 2", part2(lines))
}

func part2(lines []string) (sum int) {
	mins := make([]int, len(lines[0]))
	tilted := make([][]rune, len(lines))
	for i := 0; i < len(lines); i++ {
		tilted[i] = make([]rune, len(lines[i]))
	}
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			r := lines[i][j]
			switch r {
			case 'O':
				if i > mins[j] {
					tilted[mins[j]][j] = 'O'
					tilted[i][j] = '.'
					mins[j] = mins[j] + 1
				} else {
					tilted[i][j] = 'O'
					mins[j] = i + 1
				}
			case '#':
				mins[j] = i + 1
				tilted[i][j] = '#'
			default:
				tilted[i][j] = rune(r)
			}
		}
	}
	for i := 0; i < len(lines); i++ {
		fmt.Println(i, string(tilted[i]))
		for j := 0; j < len(lines[0]); j++ {
			if tilted[i][j] == 'O' {
                sum += len(lines) - i
            }
		}
	}
	return
}

func part1(lines []string) (sum int) {
	return
}
