package main

import (
	"io/ioutil"
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	content, _ := ioutil.ReadFile("../inputs/day01.txt")
	text := strings.TrimSpace(string(content))
	res := part1(text)
	if res == 246 {
		t.Logf("Success !")
	} else {
		t.Error("Expected 246")
	}
}

func TestPart2(t *testing.T) {
	content, _ := ioutil.ReadFile("../inputs/day01.txt")
	text := strings.TrimSpace(string(content))
	res := part2(text)
	if res == 124 {
		t.Logf("Success !")
	} else {
		t.Error("Expected 124")
	}
}
