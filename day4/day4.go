package main

import (
	"aoc23/utils"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

func main() {
	lines, err := utils.ReadInput("day4")
	if err != nil {
		panic(err)
	}
	fmt.Println("Part 1: ", part1(lines))
	fmt.Println("Part 2: ", part2(lines))
}

func part1(cards []string) (sum int) {
	for _, card := range cards {
		_, winning, numbers := getNumbers(card)
		cardMatch := 0.0
		for n := range numbers {
			if slices.Contains(winning, numbers[n]) {
				cardMatch += 1
			}
		}
		if cardMatch > 0 {
			sum += int(math.Pow(2, cardMatch-1))
		}
	}
	return
}

func part2(cards []string) (sum int) {
	cardMap := map[int][][][]int{}
	idSlice := []int{}
	for _, card := range cards {
		id, winning, numbers := getNumbers(card)
		cardMap[id] = append(cardMap[id], [][]int{winning, numbers})
		idSlice = append(idSlice, id)
	}
	for _, id := range idSlice {
		scratchcards := cardMap[id]
		for _, scratchcard := range scratchcards {
			sum += 1
			winning, numbers := scratchcard[0], scratchcard[1]
			cardMatch := 0.0
			for n := range numbers {
				if slices.Contains(winning, numbers[n]) {
					cardMatch += 1
				}
			}
			if cardMatch > 0 {
				for i := id + 1; i < id+1+int(cardMatch); i++ {
					if i > len(idSlice) {
						break
					}
					cardMap[i] = append(cardMap[i], cardMap[i][0])
				}
			}
		}
	}
	return
}

func getNumbers(card string) (id int, winning []int, numbers []int) {
	split := strings.Split(card, ":")
	idString, _ := strings.CutPrefix(split[0], "Card ")
	id, _ = strconv.Atoi(strings.Trim(idString, " "))
	numbersSplit := strings.Split(split[1], " | ")
	winningSlice := strings.Split(numbersSplit[0], " ")
	for _, w := range winningSlice {
		w = strings.Trim(w, " ")
		if w == "" {
			continue
		}
		n, _ := strconv.Atoi(w)
		winning = append(winning, n)
	}
	numbersSlice := strings.Split(numbersSplit[1], " ")
	for _, w := range numbersSlice {
		w = strings.Trim(w, " ")
		if w == "" {
			continue
		}
		n, _ := strconv.Atoi(w)
		numbers = append(numbers, n)
	}
	return
}
