package main

import (
	"aoc23/utils"
	"fmt"
)

func main() {
	lines, err := utils.ReadInput("day9")
	if err != nil {
		panic(err)
	}
	f, b := getForecasts(lines)
	fmt.Println("part 1", f)
	fmt.Println("part 2", b)
}

func getForecasts(lines []string) (forward, backward int) {
	for _, line := range lines {
		values, err := utils.AtoiSplit(line, " ")
		if err != nil {
			panic(err)
		}
		differences := make([][]int, 1)
		differences[0] = values
		stop := false
		for j := 0; !stop; j++ {
			diffs := differences[j]
			if j == len(differences)-1 && !stop {
				differences = append(differences, make([]int, len(diffs)-1))
			}
			all0 := true
			for k := 0; k < len(diffs)-1; k++ {
				diff := diffs[k+1] - diffs[k]
				differences[j+1][k] = diff
				all0 = all0 && diff == 0
			}
			if all0 {
				stop = true
			}
		}
		for j := len(differences) - 2; j >= 0; j-- {
			diff := differences[j]
			prev := differences[j+1]
			diff = append(diff, diff[len(diff)-1]+prev[len(prev)-1]) // part 1
			diff = append([]int{diff[0] - prev[0]}, diff...)         // part 2
			differences[j] = diff
		}
		forward += differences[0][len(differences[0])-1] // part 1
		backward += differences[0][0]                    // part 2
	}
	return
}
