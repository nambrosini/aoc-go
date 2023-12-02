package day02

import (
	"fmt"
	"github.com/nambrosini/aoc-go/util"
	"log"
	"strconv"
	"strings"
)

func Run() {
	input := generate(util.ReadInput(2))
	fmt.Println("Day 02")
	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

func generate(input string) []Game {
	var games []Game
	s := strings.Split(input, "\n")
	for _, game := range s {
		var g Game
		hands := strings.Split(strings.Split(game, ": ")[1], "; ")
		for _, hand := range hands {
			var h Hand
			cubes := strings.Split(hand, ", ")
			for _, cube := range cubes {
				c := strings.Split(cube, " ")
				v, err := strconv.Atoi(c[0])
				if err != nil {
					log.Fatal("Couldn't convert number", err)
				}
				h = append(h, Cube{
					Color: Color(c[1]),
					count: v,
				})
			}
			g = append(g, h)
		}
		games = append(games, g)
	}
	return games
}

var limits = map[Color]int{
	Red:   12,
	Green: 13,
	Blue:  14,
}

func part1(games []Game) int {
	sum := 0
	for i, game := range games {
		possible := true
		for _, hand := range game {
			if !checkPossible(hand) {
				possible = false
				break
			}
		}
		if possible {
			sum += i + 1
		}
	}
	return sum
}

func part2(games []Game) int {
	sum := 0
	for _, game := range games {
		colors := make(map[Color]int)
		for _, hand := range game {
			for _, cube := range hand {
				if colors[cube.Color] < cube.count {
					colors[cube.Color] = cube.count
				}
			}
		}
		res := 1
		for _, val := range colors {
			res *= val
		}
		sum += res
	}
	return sum
}

func checkPossible(hand Hand) bool {
	for _, cube := range hand {
		if cube.count > limits[cube.Color] {
			return false
		}
	}

	return true
}

type Hand []Cube
type Game []Hand

type Cube struct {
	Color
	count int
}

type Color string

const (
	Red   Color = "red"
	Blue  Color = "blue"
	Green Color = "green"
)
