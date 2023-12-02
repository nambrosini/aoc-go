package day01

import (
	"github.com/nambrosini/aoc-go/util"
	"log"
	"slices"
	"strconv"
	"strings"
	"unicode"
)

func Run() {
	input := util.ReadInput(1)
	log.Printf("Part 1: %d\n", part1(input))
	log.Printf("Part 2: %d\n", part2(input))
}

func part1(input string) int {
	lines := strings.Split(input, "\n")
	sum := 0
	for _, line := range lines {
		var nums []int
		for _, c := range line {
			if unicode.IsDigit(c) {
				v, err := strconv.Atoi(string(c))
				if err != nil {
					log.Fatal("Coulnd't conver num", err)
				}
				nums = append(nums, v)
			}
		}
		sum += nums[0]*10 + nums[len(nums)-1]
	}
	return sum
}

func part2(input string) int {
	lines := strings.Split(input, "\n")
	numbers := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	sum := 0
	for _, line := range lines {
		var nums []int
		for i, c := range line {
			if unicode.IsDigit(c) {
				nums = append(nums, toDigit(c))
				continue
			}
			for x := 3; x <= 5; x++ {
				if i+x > len(line) {
					break
				}
				sub := line[i : i+x]
				index := slices.Index(numbers, sub)
				if index != -1 {
					nums = append(nums, index+1)
					break
				}
			}
		}
		sum += nums[0]*10 + nums[len(nums)-1]
	}
	return sum
}

func toDigit(c rune) int {
	v, err := strconv.Atoi(string(c))
	if err != nil {
		log.Fatal("Couldn't convert num", err)
	}
	return v
}
