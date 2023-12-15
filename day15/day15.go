package main

import (
	"aoc23/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	content, err := utils.GetInputAsString("day15")
	if err != nil {
		panic(err)
	}
	content, _ = strings.CutSuffix(content, "\n")
	fmt.Println("part 1", part1(content))
	fmt.Println("part 2", part2(content))
}

func part1(input string) (sum int) {
	sequences := strings.Split(input, ",")
	for _, s := range sequences {
		sum += hash(s)
	}
	return
}

func hash(s string) (hash int) {
	for _, r := range s {
		hash += int(r)
		hash *= 17
		hash = hash % 256
	}
	return
}

type Box struct {
	label  string
	length int
}

func part2(input string) (sum int) {
	m := make(map[int][]Box)
	sequences := strings.Split(input, ",")
	for _, s := range sequences {
		isRemove := strings.HasSuffix(s, "-")
		if isRemove {
			label, _ := strings.CutSuffix(s, "-")
			boxIdx := hash(label)
			boxes, anyBoxes := m[boxIdx]
			if anyBoxes {
				newBoxes := make([]Box, 0)
				for _, b := range boxes {
					if b.label != label {
						newBoxes = append(newBoxes, b)
					}
				}
				m[boxIdx] = newBoxes
			}
		} else {
			split := strings.Split(s, "=")
			label := split[0]
			boxIdx := hash(label)
			boxes, _ := m[boxIdx]
			length, _ := strconv.Atoi(split[1])
			newBoxes := make([]Box, 0)
			found := false
			for _, b := range boxes {
				if b.label == label {
					found = true
					newBoxes = append(newBoxes, Box{label, length})
				} else {
					newBoxes = append(newBoxes, b)
				}
			}
			if !found {
				newBoxes = append(newBoxes, Box{label, length})
			}
			m[boxIdx] = newBoxes
		}
	}

	boxIndexes := utils.Keys(m)
	for _, boxIdx := range boxIndexes {
		boxes := m[boxIdx]
		for i, box := range boxes {
			sum += box.length * (1 + i) * (1 + boxIdx)
		}
	}
	return
}
