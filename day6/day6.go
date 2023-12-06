package main

import (
	"aoc23/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	lines, err := utils.ReadInput("day6")
	if err != nil {
		panic(err)
	}
	fmt.Println("Part 1:", part1(lines))
	fmt.Println("Part 2:", part2(lines))
}

func part1(lines []string) (mul int) {
	mul = 1
	times := extractInts(lines[0], "Time:")
	distances := extractInts(lines[1], "Distance:")
	for i, t := range times {
		r := distances[i]
		min, max := getMinMax(t, r)
		mul *= max - min + 1
	}
	return
}

func part2(lines []string) int {
	t := 42899189
	r := 308117012911467
	min, max := getMinMax(t, r)
	return max - min + 1
}

func getMinMax(t int, r int) (int, int) {
	min := ((float64(t) - (math.Sqrt(math.Pow(float64(t), 2) - float64(4*r)))) / 2)
	max := ((float64(t) + (math.Sqrt(math.Pow(float64(t), 2) - float64(4*r)))) / 2)
	if min == math.Round(min) {
		min += 1
	} else {
		min = math.Ceil(min)
	}
	if max == math.Round(max) {
		max -= 1
	} else {
		max = math.Floor(max)
	}
	return int(min), int(max)
}

func extractInts(s string, prefix string) (ints []int) {
	withoutPrefix, _ := strings.CutPrefix(s, prefix)
	strSlice := strings.Split(withoutPrefix, " ")
	for _, t := range strSlice {
		t = strings.Trim(t, " ")
		if t == "" {
			continue
		}
		time, _ := strconv.Atoi(t)
		ints = append(ints, time)
	}
	return
}
