package main

import (
	"aoc23/utils"
	"fmt"
)

func main() {
	lines, err := utils.GetInput("day16")
	if err != nil {
		panic(err)
	}
	fmt.Println("part 1:", part1(lines, Current{Point{0, 0}, Point{0, 1}}))
	fmt.Println("part 2:", part2(lines))
}

type Point [2]int

/* const (
	Down  = [2]int{1, 0}
	Up    = [2]int{-1, 0}
	Right = [2]int{0, 1}
	Left  = [2]int{0, -1}
) */

type Current struct {
	pos Point
	dir Point
}

func part1(grid []string, start Current) int {
	currents := make([]Current, 0)
	visited := make(map[[2]Point]int)
	mirrors := make(map[[2]Point]bool)
	currents = append(currents, start)
	for len(currents) > 0 {
		newCurrents := []Current{}
		for i := 0; i < len(currents); i++ {
			current := currents[i]
			switch grid[current.pos[0]][current.pos[1]] {
			case '.':
				key := [2]Point{current.pos, current.dir}
				visited[key] = visited[key] + 1
				nextPos := Point{current.pos[0] + current.dir[0], current.pos[1] + current.dir[1]}
				if _, alreadyVisited := visited[[2]Point{nextPos, current.dir}]; !alreadyVisited && isValidPos(len(grid), len(grid[0]), nextPos) {
					newCurrents = append(newCurrents, Current{nextPos, current.dir})
				}
			case '\\':
				mirrors[[2]Point{current.pos, current.dir}] = true
				newDir := Point{current.dir[1], current.dir[0]} // invert
				nextPos := Point{current.pos[0] + newDir[0], current.pos[1] + newDir[1]}
				if _, alreadyVisited := visited[[2]Point{nextPos, newDir}]; !alreadyVisited && isValidPos(len(grid), len(grid[0]), nextPos) {
					if _, mirrorVisited := mirrors[[2]Point{nextPos, newDir}]; !mirrorVisited {
						newCurrents = append(newCurrents, Current{nextPos, newDir})
					}
				}
			case '/':
				mirrors[[2]Point{current.pos, current.dir}] = true
				newDir := Point{-1 * current.dir[1], -1 * current.dir[0]} // invert
				nextPos := Point{current.pos[0] + newDir[0], current.pos[1] + newDir[1]}
				if _, alreadyVisited := visited[[2]Point{nextPos, newDir}]; !alreadyVisited && isValidPos(len(grid), len(grid[0]), nextPos) {
					if _, mirrorVisited := mirrors[[2]Point{nextPos, newDir}]; !mirrorVisited {
						newCurrents = append(newCurrents, Current{nextPos, newDir})
					}
				}
			case '|':
				mirrors[[2]Point{current.pos, current.dir}] = true
				upDir := Point{-1, 0}
				downDir := Point{1, 0}
				nextPos := Point{current.pos[0] + upDir[0], current.pos[1] + upDir[1]}
				if _, alreadyVisited := visited[[2]Point{nextPos, upDir}]; !alreadyVisited && isValidPos(len(grid), len(grid[0]), nextPos) {
					if _, mirrorVisited := mirrors[[2]Point{nextPos, upDir}]; !mirrorVisited {
						newCurrents = append(newCurrents, Current{nextPos, upDir})
					}
				}
				nextPos = Point{current.pos[0] + downDir[0], current.pos[1] + downDir[1]}
				if _, alreadyVisited := visited[[2]Point{nextPos, downDir}]; !alreadyVisited && isValidPos(len(grid), len(grid[0]), nextPos) {
					if _, mirrorVisited := mirrors[[2]Point{nextPos, downDir}]; !mirrorVisited {
						newCurrents = append(newCurrents, Current{nextPos, downDir})
					}
				}
			case '-':
				mirrors[[2]Point{current.pos, current.dir}] = true
				leftDir := Point{0, -1}
				rightDir := Point{0, 1}
				nextPos := Point{current.pos[0] + leftDir[0], current.pos[1] + leftDir[1]}
				if _, alreadyVisited := visited[[2]Point{nextPos, leftDir}]; !alreadyVisited && isValidPos(len(grid), len(grid[0]), nextPos) {
					if _, mirrorVisited := mirrors[[2]Point{nextPos, leftDir}]; !mirrorVisited {
						newCurrents = append(newCurrents, Current{nextPos, leftDir})
					}
				}
				nextPos = Point{current.pos[0] + rightDir[0], current.pos[1] + rightDir[1]}
				if _, alreadyVisited := visited[[2]Point{nextPos, rightDir}]; !alreadyVisited && isValidPos(len(grid), len(grid[0]), nextPos) {
					if _, mirrorVisited := mirrors[[2]Point{nextPos, rightDir}]; !mirrorVisited {
						newCurrents = append(newCurrents, Current{nextPos, rightDir})
					}
				}
			}
		}
		currents = newCurrents
		// fmt.Println(currents)
	}
	energizedMap := make(map[Point]int)
	for k, _ := range visited {
		energizedMap[k[0]] = energizedMap[k[0]] + 1
	}

	/* for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if e, found := energizedMap[Point{i, j}]; found {
				if e > 1 {
					fmt.Printf("%d", e)
				} else {
					fmt.Print("#")
				}
			} else {
				fmt.Print(string(grid[i][j]))
			}
		}
		fmt.Println()
	} */
	uniqueMirrors := make(map[Point]bool)
	for k, _ := range mirrors {
		uniqueMirrors[k[0]] = true
	}
	return len(uniqueMirrors) + len(energizedMap)
}

func isValidPos(maxI, maxJ int, pos Point) bool {
	return pos[0] >= 0 && pos[0] < maxI && pos[1] >= 0 && pos[1] < maxJ
}

func part2(grid []string) (max int) {
	fmt.Println("Top")
	for i := 0; i < len(grid[0]); i++ {
		value := part1(grid, Current{Point{0, i}, Point{1, 0}})
		if value > max {
			max = value
		}
	}
	fmt.Println("Bottom")
	for i := 0; i < len(grid[0]); i++ {
		value := part1(grid, Current{Point{len(grid) - 1, i}, Point{-1, 0}})
		if value > max {
			max = value
		}
	}
	fmt.Println("Left")
	for i := 0; i < len(grid); i++ {
		value := part1(grid, Current{Point{i, 0}, Point{0, 1}})
		if value > max {
			max = value
		}
	}
	fmt.Println("Right")
	for i := 0; i < len(grid); i++ {
		value := part1(grid, Current{Point{len(grid[0]) - 1, 0}, Point{0, -1}})
		if value > max {
			max = value
		}
	}
	return
}
