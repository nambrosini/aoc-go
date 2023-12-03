package day03

import (
	"fmt"
	"github.com/nambrosini/aoc-go/util"
	"log"
	"slices"
	"strconv"
	"strings"
	"unicode"
)

func Run() {
	input := generate(util.ReadInput(3))
	fmt.Println("Day 03")
	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

func generate(input string) [][]rune {
	schema := [][]rune{}
	for _, l := range strings.Split(input, "\n") {
		row := []rune{}
		for _, c := range l {
			row = append(row, c)
		}
		schema = append(schema, row)
	}
	return schema
}

func part1(input [][]rune) int {
	lastNum := 0
	sum := 0
	for i, row := range input {
		for j, e := range row {
			if unicode.IsDigit(e) && hasAdjacentSymbol(i, j, input) {
				v, err := getPartValue(j, row)
				if err != nil {
					log.Fatal("Couldn't get value part", err)
				}
				if v == lastNum {
					continue
				}
				sum += v
				lastNum = v
			} else {
				lastNum = 0
			}
		}
	}
	return sum
}

func part2(input [][]rune) int {
	res := 0
	for i, row := range input {
		for j, e := range row {
			if e == '*' {
				res += getGearRatio(i, j, input)
			}
		}
	}
	return res
}

func getGearRatio(x, y int, schema [][]rune) int {
	var gears []int
	for i := -1; i <= 1; i++ {
		if x+i < 0 || x+i >= len(schema) {
			continue
		}
		for j := -1; j <= 1; j++ {
			if y+j < 0 && y+j >= len(schema[0]) {
				continue
			}
			el := schema[x+i][y+j]
			if unicode.IsDigit(el) {
				v, err := getPartValue(y+j, schema[x+i])
				if err != nil {
					log.Fatal("Error", err)
				}
				if !slices.Contains(gears, v) {
					gears = append(gears, v)
					if len(gears) == 2 {
						return gears[0] * gears[1]
					}
				}
			}
		}
	}
	return 0
}

func hasAdjacentSymbol(x, y int, schema [][]rune) bool {
	for i := -1; i <= 1; i++ {
		if x+i < 0 || x+i >= len(schema) {
			continue
		}
		for j := -1; j <= 1; j++ {
			if y+j < 0 || y+j >= len(schema[0]) {
				continue
			}
			el := schema[x+i][y+j]
			if el != '.' && !unicode.IsDigit(el) {
				return true
			}
		}
	}
	return false
}

func getPartValue(x int, row []rune) (int, error) {
	start, end := x, x
	startFound, endFound := false, false

	for i := x - 1; !startFound; i-- {
		if i >= 0 && unicode.IsDigit(row[i]) {
			start = i
		} else {
			startFound = true
		}
	}

	for i := x + 1; !endFound; i++ {
		if i < len(row) && unicode.IsDigit(row[i]) {
			end = i
		} else {
			endFound = true
		}
	}

	return strconv.Atoi(string(row[start : end+1]))
}
