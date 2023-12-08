package main

import (
	"aoc23/utils"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

const (
	FiveOfKind  = 7
	FourOfKind  = 6
	FullHouse   = 5
	ThreeOfKind = 4
	TwoPair     = 3
	OnePair     = 2
	HighCard    = 1
)

func main() {
	lines, err := utils.ReadInput("day7")
	if err != nil {
		panic(err)
	}
	fmt.Println("Part 1:", part1(lines))
	fmt.Println("Part 2:", part2(lines))
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
		t := getHandType(hand, false)
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
		t := getHandType(hand, true)
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
					pa++
				}
				if !foundb {
					pb, _ = strconv.Atoi(string(rb))
					pb++
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
		if strings.Contains(hand.hand, "J") && hand.t == OnePair {
			fmt.Println(i, hand.hand, hand.t, hand.bid, win)
		}
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

func getHandType(hand string, joker bool) int {
	occ := occurrences(hand)
	keys := Keys(occ)
	jokerOcc := 0
	if joker {
		jokerOcc = occ['J']
	}
	if len(occ) == 1 {
		return FiveOfKind
	}
	if len(occ) == 2 {
		if occ[keys[0]] == 4 || occ[keys[1]] == 4 {
			if jokerOcc > 0 {
				return FiveOfKind
			}
			return FourOfKind
		}
		if occ[keys[0]] == 3 || occ[keys[1]] == 3 {
			switch {
			case jokerOcc == 1:
				return FourOfKind
			case jokerOcc == 2:
				return FiveOfKind
			case jokerOcc == 3:
				return FiveOfKind
			}
			return FullHouse
		}
	}
	if len(occ) == 3 {
		three := false
		twos := make([]rune, 0)
		for _, k := range keys {
			o := occ[k]
			if o == 3 {
				three = true
			}
			if o == 2 {
				twos = append(twos, k)
			}
		}
		if three {
			switch jokerOcc {
			case 1:
				return FourOfKind
			case 2:
				return FiveOfKind
			case 3:
				return FourOfKind
			}
			return ThreeOfKind
		} else {
			if len(twos) == 2 {
				switch {
				case jokerOcc == 2:
					return FourOfKind
				case jokerOcc == 1:
					return FullHouse
				}
				return TwoPair
			} else if len(twos) == 1 {
				switch {
				case jokerOcc == 1:
					return ThreeOfKind
				case jokerOcc == 2 && !slices.Contains(twos, 'J'):
					return FourOfKind
				case jokerOcc == 3:
					return FiveOfKind
				}
				return OnePair
			}
		}
	}
	if len(occ) == 4 {
		for _, k := range keys {
			o := occ[k]
			if o == 2 {
				if k != 'J' {
					switch jokerOcc {
					case 1:
						return ThreeOfKind
					case 2:
						return FourOfKind
					case 3:
						return FiveOfKind
					}
				} else {
					return ThreeOfKind
				}
				return OnePair
			}
		}
	}
	if jokerOcc == 1 {
        return OnePair
    }
	return HighCard
}

func Keys(m map[rune]int) (keys []rune) {
	for k := range m {
		keys = append(keys, k)
	}
	return
}
