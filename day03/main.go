package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("inputs/day03.txt")
	if err != nil {
		fmt.Println("Error reading file")
	}
	text := strings.TrimSpace(string(content))

	fmt.Println("part 1:", part1(text))
}

// ans: 993
func part1(text string) int {
	var validTriangles [][]int
	for _, triangle := range parseInput(text) {
		if isValidTriangle(triangle) {
			validTriangles = append(validTriangles, triangle)
		}
	}
	return len(validTriangles)
}

func isValidTriangle(sides []int) bool {
  valid1 := sides[0] + sides[1] > sides[2]
  valid2 := sides[1] + sides[2] > sides[0]
  valid3 := sides[2] + sides[0] > sides[1]
  return valid1 && valid2 && valid3
}

func parseInput(text string) [][]int {
	triangles := make([][]int, 0)
	for lineIdx, line := range strings.Split(text, "\n") {
		triangles = append(triangles, make([]int, 0))
		for _, sideStr := range strings.Split(line, " ") {
			if sideStr != "" {
				side, _ := strconv.Atoi(sideStr)
				triangles[lineIdx] = append(triangles[lineIdx], side)
			}
		}
	}
	return triangles
}
