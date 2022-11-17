package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

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

func main() {
	content, err := ioutil.ReadFile("inputs/day01.txt")
	if err != nil {
		fmt.Println("Error reading file")
	}

	pos := Position{0, 0, 0}

	text := strings.TrimSpace(string(content))
	for _, move := range strings.Split(text, ",") {
		chars := strings.Split(strings.TrimSpace(move), "")
		direction := chars[0]
		magnitude, err := strconv.Atoi(strings.Join(chars[1:], ""))
		if err != nil {
			fmt.Println("Error parsing string to int")
		}

		pos.rotate(direction)
		pos.move(magnitude)
	}

	result := math.Abs(float64(pos.x)) + math.Abs(float64(pos.y))
	fmt.Println("result:", result)
}
