package day04

import (
	"fmt"
	"github.com/nambrosini/aoc-go/util"
	"log"
	"math"
	"strconv"
	"strings"
)

const DAY = 4

func Run() {
	input := generate(util.ReadInput(DAY))
	fmt.Printf("Day 0%d\n", DAY)
	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

func generate(input string) []int {
	lines := strings.Split(input, "\n")
	var cards []int
	for _, line := range lines {
		split := strings.Replace(strings.Split(line, ": ")[1], "| ", "", 1)
		numbers := strings.Split(split, " ")
		var res []int
		for _, n := range numbers {
			if n != "" {
				v, err := strconv.Atoi(n)
				if err != nil {
					log.Fatal("Error", err)
				}
				res = append(res, v)
			}
		}
		cards = append(cards, len(removeDuplicates(res)))
	}

	return cards
}

func removeDuplicates(input []int) []int {
	encountered := map[int]bool{}
	var result []int

	for _, value := range input {
		if encountered[value] == false {
			encountered[value] = true
			result = append(result, value)
		}
	}

	return result
}

func getMax(cards []int) int {
	maxLen := 0
	for _, card := range cards {
		if card > maxLen {
			maxLen = card
		}
	}
	return maxLen
}

func part1(cards []int) int {
	maxLen := getMax(cards)
	sum := 0
	for _, card := range cards {
		if card != maxLen {
			sum += int(math.Pow(2, float64(maxLen)-float64(card)-1))
		}
	}
	return sum
}

func part2(cards []int) int {
	maxLen := getMax(cards)
	res := map[int]int{}
	sum := 0
	for i, card := range cards {
		res[i] += 1
		sum += res[i]
		for j := 1; j <= maxLen-card; j++ {
			res[i+j] += res[i]
		}
	}
	return sum
}
