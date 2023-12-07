package day07

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/maccam912/advent-of-code-2023/util"
)

type Hand struct {
	cards []rune
	bid   int
}

func handClassA(hand Hand) int {
	classes := map[rune]int{}
	for _, card := range hand.cards {
		classes[card] += 1
	}
	// Five of a kind: check for any class having 5
	// Four of a kind: check for any class having 4
	// Full house: check for any class having 3 and any other class having 2
	// Three of a kind: check for any class having 3
	// Two pair: check for any class having 2 and any other class having 2
	// One pair: check for any class having 2
	// High card: check for any class having 1
	for _, count := range classes {
		if count == 5 {
			return 8 // Five of a kind
		}
		if count == 4 {
			return 7 // Four of a kind
		}
	}
	if len(classes) == 2 {
		// Will be true for full house or for 4 of a kind, but 4 of a kind checked already
		return 6 // Full house
	}
	for _, count := range classes {
		if count == 3 {
			return 5 // Three of a kind
		}
	}
	if len(classes) == 3 {
		// Will be true for two pair or for three of a kind, but three of a kind checked already
		return 4 // Two pair
	}
	for _, count := range classes {
		if count == 2 {
			return 3 // One pair
		}
	}
	return 2 // High card
}

func handClassB(hand Hand) int {
	classes := map[rune]int{}
	jCount := 0
	for _, card := range hand.cards {
		if card != 'J' {
			classes[card] += 1
		} else {
			jCount++
		}
	}

	// Five of a kind: check for any class having 5
	// Four of a kind: check for any class having 4
	// Full house: check for any class having 3 and any other class having 2
	// Three of a kind: check for any class having 3
	// Two pair: check for any class having 2 and any other class having 2
	// One pair: check for any class having 2
	// High card: check for any class having 1

	highestCount := 0
	secondHighestCount := 0
	for _, count := range classes {
		if count > highestCount {
			secondHighestCount = highestCount
			highestCount = count
		} else if count > secondHighestCount {
			secondHighestCount = count
		}
	}

	if highestCount == 5 || highestCount+jCount >= 5 {
		return 8 // Five of a kind
	} else if highestCount == 4 || highestCount+jCount >= 4 {
		return 7 // Four of a kind
	} else if highestCount == 3 && secondHighestCount == 2 ||
		highestCount+jCount >= 3 && secondHighestCount == 2 ||
		highestCount == 3 && secondHighestCount+jCount >= 2 ||
		highestCount+1 >= 3 && secondHighestCount+1 >= 2 && jCount >= 2 {
		return 6 // Full house
	} else if highestCount == 3 || highestCount+jCount >= 3 {
		return 5 // Three of a kind
	} else if highestCount == 2 && secondHighestCount == 2 ||
		highestCount+jCount >= 2 && secondHighestCount == 2 ||
		highestCount == 2 && secondHighestCount+jCount >= 2 {
		return 4 // Two pair
	} else if highestCount == 2 || highestCount+jCount >= 2 {
		return 3 // One pair
	} else {
		return 2 // High card
	}
}

func orderCardsA(runeA rune, runeB rune) int {
	order := "AKQJT98765432"
	for i := 0; i < len(order); i++ {
		if runeA == rune(order[i]) && runeB != rune(order[i]) {
			return 1
		}
		if runeB == rune(order[i]) && runeA != rune(order[i]) {
			return -1
		}
	}
	return 0
}

func orderCardsB(runeA rune, runeB rune) int {
	order := "AKQT98765432J"
	for i := 0; i < len(order); i++ {
		if runeA == rune(order[i]) && runeB != rune(order[i]) {
			return 1
		}
		if runeB == rune(order[i]) && runeA != rune(order[i]) {
			return -1
		}
	}
	return 0
}

func orderHandsA(handA Hand, handB Hand) int {
	aClass := handClassA(handA)
	bClass := handClassA(handB)
	if aClass > bClass {
		return 1
	}
	if aClass < bClass {
		return -1
	}
	// Same class, compare cards in order
	for i := 0; i < len(handA.cards); i++ {
		v := orderCardsA(handA.cards[i], handB.cards[i])
		if v != 0 {
			return v
		}
	}
	return 0
}

func orderHandsB(handA Hand, handB Hand) int {
	aClass := handClassB(handA)
	bClass := handClassB(handB)
	if aClass > bClass {
		return 1
	}
	if aClass < bClass {
		return -1
	}
	// Same class, compare cards in order
	for i := 0; i < len(handA.cards); i++ {
		v := orderCardsB(handA.cards[i], handB.cards[i])
		if v != 0 {
			return v
		}
	}
	return 0
}

func parseInput(path string) []Hand {
	lines, _ := util.ReadLines(path)
	hands := []Hand{}
	for _, line := range lines {
		hand := Hand{}
		fields := strings.Fields(line)
		hand.cards = []rune(fields[0])
		hand.bid, _ = strconv.Atoi(fields[1])
		hands = append(hands, hand)
	}
	return hands
}

func A(path string) int {
	hands := parseInput(path)
	totalWinnings := 0
	slices.SortFunc(hands, orderHandsA)
	for i, hand := range hands {
		totalWinnings += hand.bid * (i + 1)
	}
	return totalWinnings
}

func B(path string) int {
	hands := parseInput(path)
	totalWinnings := 0
	slices.SortFunc(hands, orderHandsB)
	for i, hand := range hands {
		totalWinnings += hand.bid * (i + 1)
	}
	return totalWinnings
}

// Run runs the day 1 challenge.
func Run() {
	partA := A("day07/input.txt")
	fmt.Printf("Part A: %v\n", partA)

	partB := B("day07/input.txt")
	fmt.Printf("Part B: %v\n", partB)
}
