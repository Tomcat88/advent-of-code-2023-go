package main

import (
	"aoc23/utils"
	"fmt"
	"strconv"
	"strings"

)

func main() {
	lines, err := utils.ReadInput("day1")
	if err != nil {
		panic(err)
	}
	digits := map[string]int{
		"0": 0,
		"1": 1,
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
		"9": 9,
	}
	fmt.Println("part 1", process(lines, digits))

	digits = map[string]int{
		"0":     0,
		"1":     1,
		"2":     2,
		"3":     3,
		"4":     4,
		"5":     5,
		"6":     6,
		"7":     7,
		"8":     8,
		"9":     9,
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
	fmt.Println("part 2", process(lines, digits))
}

func process(lines []string, digits map[string]int) (sum int) {
	for _, line := range lines {
		found := make([]int, 0)
		for i := 0; i < len(line); i++ {
			for k, d := range digits {
				if strings.HasPrefix(line[i:], k) {
					found = append(found, d)
				}
			}
		}
		if len(found) == 0 {
			continue
		}
		n, _ := strconv.Atoi(fmt.Sprintf("%d%d", found[0], found[len(found)-1]))
		sum += n
	}
	return
}
