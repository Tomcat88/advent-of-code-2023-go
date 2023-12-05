package main

import (
	"aoc23/utils"
	"fmt"
	"strconv"
)

func main() {
	lines, err := utils.ReadInput("day3")
	if err != nil {
		panic(err)
	}
	fmt.Println("Part 1:", part1(lines))
	fmt.Println("Part 2:", part2(lines))
}

func part1(lines []string) (sum int) {
	for i, line := range lines {
		number := ""
		symbolNear := false
		for j, c := range line {
			_, NaN := strconv.Atoi(string(c))
			if NaN == nil {
				number += string(c)
				if len(isSymbolNear(lines, i, j, isSymbol)) != 0 {
					symbolNear = true
				}
			}
			if (NaN != nil || j == len(line)-1) && number != "" && symbolNear {
				n, err := strconv.Atoi(number)
				if err != nil {
					panic(err)
				}
				sum += n
				number = ""
				symbolNear = false
			} else if number != "" && NaN != nil {
				number = ""
				symbolNear = false
			}
		}
	}
	return
}

func part2(lines []string) (sum int) {
	// create a map gear coors -> numbers
	// when found two, mul together and sum
	near := map[string]int{}
	for i, line := range lines {
		number := ""
		gearNear := ""
		fmt.Println(line)
		for j, c := range line {
			_, NaN := strconv.Atoi(string(c))
			if NaN == nil {
				number += string(c)
				if coords:= isSymbolNear(lines, i, j, isGear); len(coords) > 0 {
					gearNear = fmt.Sprintf("%d,%d", coords[0][0], coords[0][1]) 
					// fmt.Println("found gear", coords[0][0],coords[0][1], number)
				}
			}
			if (NaN != nil || j == len(line)-1) && number != "" && gearNear == "" {
				number = ""
				gearNear = ""
			} else if number != "" && NaN != nil {
				currentNumber, err := strconv.Atoi(number)
				if err != nil {
					panic(err)
				}
				numberNear, found := near[gearNear]
				if !found {
					near[gearNear] = currentNumber
				} else {
					fmt.Println(currentNumber, numberNear)
					sum += numberNear * currentNumber
					near[gearNear] = currentNumber
				}
				number = ""
				gearNear = ""
			}
		}
		gearNear = ""
	}
	return
}

func isSymbolNear(lines []string, i, j int, isSymbolFn func(byte) bool) [][]int {
	coords := [][]int{
		{i + 1, j},
		{i - 1, j},
		{i, j + 1},
		{i, j - 1},
		{i + 1, j + 1},
		{i + 1, j - 1},
		{i - 1, j + 1},
		{i - 1, j - 1},
	}
	for _, coord := range coords {
		if coord[0] >= 0 && coord[0] < len(lines) && coord[1] >= 0 && coord[1] < len(lines[0]) {
			if isSymbolFn(lines[coord[0]][coord[1]]) {
				return [][]int{coord}
			}
		}
	}
	return make([][]int, 0)
}

func isGear(b byte) bool {
	return string(b) == "*"
}

func isSymbol(b byte) bool {
	_, nan := strconv.Atoi(string(b))
	return nan != nil && string(b) != "."
}
