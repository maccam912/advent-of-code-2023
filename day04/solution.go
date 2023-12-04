package day04

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/maccam912/advent-of-code-2023/util"
)

type Card struct {
	ID             int
	WinningNumbers []int
	YourNumbers    []int
	Copies         int
}

func parseCard(line string) Card {
	gameNum, numbers, _ := strings.Cut(line, ":")
	_, idStr, _ := strings.Cut(strings.TrimSpace(gameNum), " ")
	id, _ := strconv.Atoi(idStr)
	winningNumbersStr, yourNumbersStr, _ := strings.Cut(strings.TrimSpace(numbers), "|")
	winningNumbers := []int{}
	for _, num := range strings.Fields(strings.TrimSpace(winningNumbersStr)) {
		parsedNum, _ := strconv.Atoi(num)
		winningNumbers = append(winningNumbers, parsedNum)
	}
	yourNumbers := []int{}
	for _, num := range strings.Fields(strings.TrimSpace(yourNumbersStr)) {
		parsedNum, _ := strconv.Atoi(num)
		yourNumbers = append(yourNumbers, parsedNum)
	}

	return Card{id, winningNumbers, yourNumbers, 1}
}

func parseCards(path string) []Card {
	lines, _ := util.ReadLines(path)
	cards := []Card{}
	for _, line := range lines {
		cards = append(cards, parseCard(line))
	}
	return cards
}

func (card Card) CountWinningNumbers() int {
	winningSet := map[int]bool{}
	for _, num := range card.WinningNumbers {
		winningSet[num] = true
	}
	count := 0
	for _, num := range card.YourNumbers {
		if winningSet[num] {
			count++
		}
	}
	return count
}

func A(path string) int {
	cards := parseCards(path)
	score := 0
	for _, card := range cards {
		count := card.CountWinningNumbers()
		if count > 0 {
			cardScore := math.Pow(2, float64(count-1))
			score += int(cardScore)
		}
	}
	return score
}

func B(path string) int {
	cards := parseCards(path)
	for i, card := range cards {
		count := card.CountWinningNumbers()
		for j := 0; j < count; j++ {
			idx := j + i + 1
			cards[idx].Copies += card.Copies
		}
	}
	// return sum of copies
	totalCopies := 0
	for _, card := range cards {
		totalCopies += card.Copies
	}
	return totalCopies
}

// Run runs the day 4 challenge.
func Run() {
	partA := A("day04/input.txt")
	fmt.Printf("Part A: %v\n", partA)

	partB := B("day04/input.txt")
	fmt.Printf("Part B: %v\n", partB)
}
