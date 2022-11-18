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

func part2(text string) string {
	keypad := Keypad2{0, 2}
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
	return map[Keypad]string{
		{0, 0}: "1",
		{1, 0}: "2",
		{2, 0}: "3",
		{0, 1}: "4",
		{1, 1}: "5",
		{2, 1}: "6",
		{0, 2}: "7",
		{1, 2}: "8",
		{2, 2}: "9",
	}[k]
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

type Keypad2 struct {
	x int
	y int
}

var xyMaxMap = map[int]int{
	0: 2,
	1: 3,
	2: 4,
	3: 3,
	4: 2,
}

var xyMinMap = map[int]int{
	0: 2,
	1: 1,
	2: 0,
	3: 1,
	4: 2,
}

func (k *Keypad2) move(move string) {
	switch move {
	case "U":
		k.y = max(k.y-1, xyMinMap[k.x])
	case "D":
		k.y = min(k.y+1, xyMaxMap[k.x])
	case "L":
		k.x = max(k.x-1, xyMinMap[k.y])
	case "R":
		k.x = min(k.x+1, xyMaxMap[k.y])
	}
}

func (k Keypad2) toKey() string {
	return map[Keypad2]string{
		{2, 0}: "1",
		{1, 1}: "2",
		{2, 1}: "3",
		{3, 1}: "4",
		{0, 2}: "5",
		{1, 2}: "6",
		{2, 2}: "7",
		{3, 2}: "8",
		{4, 2}: "9",
		{1, 3}: "A",
		{2, 3}: "B",
		{3, 3}: "C",
		{2, 4}: "D",
	}[k]
}
