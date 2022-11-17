package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("inputs/day01.txt")
	if err != nil {
		fmt.Println("Error reading file")
	}
	text := strings.TrimSpace(string(content))

	fmt.Println("part 1:", part1(text)) // 246
	fmt.Println("part 2:", part2(text)) // 124
}

// how far away is the final location
func part1(input string) int {
	pos := Position{0, 0, 0}

	for _, move := range strings.Split(input, ",") {
		direction, magnitude := parseMoveStr(move)
		pos.rotate(direction)
		pos.move(magnitude)
	}

	return intAbs(pos.x) + intAbs(pos.y)
}

// how far away is the first location we visit twice
func part2(input string) int {
	pos := Position{0, 0, 0}
	vistied := History{}

	for _, move := range strings.Split(input, ",") {
		direction, magnitude := parseMoveStr(move)
		pos.rotate(direction)

		for step := 0; step < magnitude; step++ {
			if magnitude > 0 {
				pos.move(1)
			} else {
				pos.move(-1)
			}
			isReturnVisit := vistied.visit(pos)
			if isReturnVisit {
				return intAbs(pos.x) + intAbs(pos.y)
			}
		}
	}

	return intAbs(pos.x) + intAbs(pos.y)
}

// std library only contains this function for floats
func intAbs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// parse a movement string, like R7 into it's direction and magnitude
func parseMoveStr(move string) (string, int) {
	chars := strings.Split(strings.TrimSpace(move), "")
	direction := chars[0]
	magnitude, err := strconv.Atoi(strings.Join(chars[1:], ""))
	if err != nil {
		fmt.Println("Error parsing string to int")
	}
	return direction, magnitude
}

type Position struct {
	x         int
	y         int
	direction int // 0 = north, 1 = east, 2 = south, 3 = west
}

func (p *Position) rotate(change string) {
	// rotate direction
	switch change {
	case "L":
		p.direction -= 1
	case "R":
		p.direction += 1
	}

	// handle wrap around cases
	if p.direction == -1 {
		p.direction = 3
	} else if p.direction == 4 {
		p.direction = 0
	}
}

func (p *Position) move(magnitude int) {
	switch p.direction {
	case 0: // north
		p.x += magnitude
	case 1: // east
		p.y += magnitude
	case 2: // south
		p.x -= magnitude
	case 3: // west
		p.y -= magnitude
	}
}

type Location struct {
	x int
	y int
}

type History map[Location]int

// Keep track of locations visit, return true if repeat visit to location
func (history History) visit(pos Position) bool {
	location := Location{pos.x, pos.y}
	history[location] += 1
	if history[location] > 1 {
		return true
	} else {
		return false
	}
}
