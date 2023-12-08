package main

import (
	"aoc23/utils"
	"fmt"
	"strings"
)

func main() {
	lines, err := utils.ReadInput("day8")
	if err != nil {
		panic(err)
	}

	fmt.Println("Part 1:", part1(lines))
	fmt.Println("Part 2:", part2(lines))
}

func part1(lines []string) (steps int) {
	instructions := lines[0]
	m := make(map[string][]string)
	replacer := strings.NewReplacer("=", "", " ", "", ",", "", "(", "", ")", "")
	for i := 2; i < len(lines); i++ {
		entry := replacer.Replace(lines[i])
		m[entry[0:3]] = []string{entry[3:6], entry[6:9]}
	}
	var current = "AAA"
	for current != "ZZZ" {
		for _, i := range instructions {
			switch i {
			case 'L':
				current = m[current][0]
			case 'R':
				current = m[current][1]
			}
			steps++
		}
	}

	return
}

func part2(lines []string) int {
	instructions := lines[0]
	m := make(map[string][]string)
	replacer := strings.NewReplacer("=", "", " ", "", ",", "", "(", "", ")", "")
	currents := make([]string, 0)
	for i := 2; i < len(lines); i++ {
		entry := replacer.Replace(lines[i])
		m[entry[0:3]] = []string{entry[3:6], entry[6:9]}
		if entry[2] == 'A' {
			currents = append(currents, entry[0:3])
		}
	}
	steps := make([]int, 0)
	for _, current := range currents {
		currentSteps := 0
		for current[2] != 'Z' {
			for _, i := range instructions {
				switch i {
				case 'L':
					current = m[current][0]
				case 'R':
					current = m[current][1]
				}
				currentSteps++
			}
		}
		steps = append(steps, currentSteps)
	}
	return utils.LcmSlice(steps)
}

