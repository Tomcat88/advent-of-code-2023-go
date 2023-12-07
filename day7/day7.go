package main

import (
	"aoc23/utils"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

const (
	FiveOfKind  = 6
	FourOfKind  = 5
	FullHouse   = 4
	ThreeOfKind = 3
	TwoPair     = 2
	OnePair     = 1
	HighCard    = 0
)

var powers = map[rune]int{
	'A': 14,
	'K': 13,
	'Q': 12,
	'J': 11,
	'T': 10,
}

func main() {
	lines, err := utils.ReadInput("day7")
	if err != nil {
		panic(err)
	}
	fmt.Println("Part 1:", part1(lines))
}

type Hand struct {
	hand     string
	t        int
	bid      int
	wildcard int
}

func part1(lines []string) (sum int) {
	var powers = map[rune]int{
		'A': 14,
		'K': 13,
		'Q': 12,
		'J': 11,
		'T': 10,
	}
	hands := make([]Hand, len(lines))
	for i, handBid := range lines {
		split := strings.Split(handBid, " ")
		hand, bid := split[0], split[1]
		t := getHandType(hand)
		bidInt, _ := strconv.Atoi(bid)
		hands[i] = Hand{hand, t, bidInt, -1}
		// fmt.Println(i, hand, bid, t)
	}
	slices.SortFunc(hands, func(a, b Hand) int {
		if a.t < b.t {
			return -1
		} else if a.t > b.t {
			return 1
		} else {
			for i, ra := range a.hand {
				rb := rune(b.hand[i])
				pa, founda := powers[ra]
				pb, foundb := powers[rb]
				if !founda {
					pa, _ = strconv.Atoi(string(ra))
				}
				if !foundb {
					pb, _ = strconv.Atoi(string(rb))
				}
				if pa < pb {
					return -1
				} else if pa > pb {
					return 1
				}
			}
		}
		return 0
	})
	for i, hand := range hands {
		win := (i + 1) * hand.bid
		fmt.Println(i, hand.hand, hand.t, hand.bid, win)
		sum += win
	}
	return
}

func part2(lines []string) (sum int) {
	var powers = map[rune]int{
		'A': 14,
		'K': 13,
		'Q': 12,
		'T': 11,
		'J': 2,
	}
	hands := make([]Hand, len(lines))
	for i, handBid := range lines {
		split := strings.Split(handBid, " ")
		hand, bid := split[0], split[1]
		t := getHandType(hand)
		bidInt, _ := strconv.Atoi(bid)
		hands[i] = Hand{hand, t, bidInt, -1}
		// fmt.Println(i, hand, bid, t)
	}
	slices.SortFunc(hands, func(a, b Hand) int {
		if a.t < b.t {
			return -1
		} else if a.t > b.t {
			return 1
		} else {
			for i, ra := range a.hand {
				rb := rune(b.hand[i])
				pa, founda := powers[ra]
				pb, foundb := powers[rb]
				if !founda {
					pa, _ = strconv.Atoi(string(ra))
				}
				if !foundb {
					pb, _ = strconv.Atoi(string(rb))
				}
				if pa < pb {
					return -1
				} else if pa > pb {
					return 1
				}
			}
		}
		return 0
	})
	for i, hand := range hands {
		win := (i + 1) * hand.bid
		fmt.Println(i, hand.hand, hand.t, hand.bid, win)
		sum += win
	}
	return
}

func occurrences(hand string) (m map[rune]int) {
	m = make(map[rune]int)
	for _, c := range hand {
		_, ok := m[c]
		if !ok {
			m[c] = 1
		} else {
			m[c] += 1
		}
	}
	return
}

func getHandType(hand string) int {
	occ := occurrences(hand)
	keys := Keys(occ)
	if len(occ) == 1 {
		return FiveOfKind
	}
	if len(occ) == 2 {
		if occ[keys[0]] == 4 || occ[keys[1]] == 4 {
			return FourOfKind
		}
		if occ[keys[0]] == 3 || occ[keys[1]] == 3 {
			return FullHouse
		}
	}
	if len(occ) == 3 {
		three := false
		twos := 0
		for _, k := range keys {
			o := occ[k]
			if o == 3 {
				three = true
			}
			if o == 2 {
				twos++
			}
		}
		if three {
			return ThreeOfKind
		} else {
			if twos == 2 {
				return TwoPair
			} else if twos == 1 {
				return OnePair
			}
		}
	}
	if len(occ) == 4 {
		for _, k := range keys {
			o := occ[k]
			if o == 2 {
				return OnePair
			}
		}
	}
	return HighCard
}

func Keys(m map[rune]int) (keys []rune) {
	for k := range m {
		keys = append(keys, k)
	}
	return
}
