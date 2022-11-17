package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("inputs/day02.txt")
	if err != nil {
		fmt.Println("Error reading file")
	}
	text := strings.TrimSpace(string(content))

	fmt.Println("part 1:", part1(text))
}

func part1(text string) string {
	keypad := Keypad{1, 1}
	out := ""

	for _, line := range strings.Split(text, "\n") {
		for _, char := range strings.Split(line, "") {
			keypad.move(char)
		}
		out += keypad.toKey()
	}

	return out
}

type Keypad struct {
	x int
	y int
}

func (k *Keypad) move(move string) {
	switch move {
	case "U":
		k.y = max(k.y-1, 0)
	case "D":
		k.y = min(k.y+1, 2)
	case "L":
		k.x = max(k.x-1, 0)
	case "R":
		k.x = min(k.x+1, 2)
	}
}

func (k Keypad) toKey() string {
	if k.x == 0 && k.y == 0 {
		return "1"
	} else if k.x == 1 && k.y == 0 {
		return "2"
	} else if k.x == 2 && k.y == 0 {
		return "3"
	} else if k.x == 0 && k.y == 1 {
		return "4"
	} else if k.x == 1 && k.y == 1 {
		return "5"
	} else if k.x == 2 && k.y == 1 {
		return "6"
	} else if k.x == 0 && k.y == 2 {
		return "7"
	} else if k.x == 1 && k.y == 2 {
		return "8"
	} else if k.x == 2 && k.y == 2 {
		return "9"
	} else {
		panic("unknown key pos")
	}
}

func min(x int, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}
