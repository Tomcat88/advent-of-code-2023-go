package main

import (
	"aoc23/utils"
	"strconv"
	"strings"
)

var config = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func main() {
	lines, err := utils.ReadInput("day2")
	if err != nil {
		panic(err)
	}
	println("Part 1: ", part1(lines))
	println("Part 2: ", part2(lines))
}

func part2(lines []string) (sum int) {
	for _, line := range lines {
		split := strings.Split(line, ":")
		sum += getMinimumConfig(split[1]) 
	}
	return
}

func getMinimumConfig(line string) (mul int) {
	sets := strings.Split(line, ";")
	minimumConfig := make(map[string]int, 0)
	mul = 1
	for _, set := range sets {
		cubes := strings.Split(set, ",")
		for _, cube := range cubes {
			cube = strings.Trim(cube, " ")
			game := strings.Split(cube, " ")
			n, _ := strconv.Atoi(game[0])
			color := game[1]
			min, ok := minimumConfig[color]
			if !ok || n > min {
				minimumConfig[color] = n
			}
		}
	}
	for _, c := range minimumConfig {
        mul *= c
    }
	return 
}

func part1(lines []string) int {
	sum := 0
	for _, line := range lines {
		split := strings.Split(line, ":")
		idString, _ := strings.CutPrefix(split[0], "Game ")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}
		if isGamePossible(split[1], config) {
			sum += id
		}
	}
	return sum
}


func isGamePossible(line string, config map[string]int) bool {
	sets := strings.Split(line, ";")
	for _, set := range sets {
		cubes := strings.Split(set, ",")
		for _, cube := range cubes {
			cube = strings.Trim(cube, " ")
			game := strings.Split(cube, " ")
			n, _ := strconv.Atoi(game[0])
			color := game[1]
			if config[color] < n {
				return false
			}
		}
	}
	return true
}

