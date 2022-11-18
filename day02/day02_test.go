package main

import (
	"io/ioutil"
	"strings"
	"testing"
)

func TestSample1(t *testing.T) {
	sample := "ULL\nRRDDD\nLURDL\nUUUUD"
	text := strings.TrimSpace(sample)
	res := part1(text)
	if res == "1985" {
		t.Logf("Success !")
	} else {
		t.Error("Expected 1985")
	}
}

func TestDay1(t *testing.T) {
	content, _ := ioutil.ReadFile("../inputs/day02.txt")
	text := strings.TrimSpace(string(content))
	res := part1(text)
	if res == "45973" {
		t.Logf("Success !")
	} else {
		t.Error("Expected 45973")
	}
}

func TestSample2(t *testing.T) {
	sample := "ULL\nRRDDD\nLURDL\nUUUUD"
	text := strings.TrimSpace(sample)
	res := part2(text)
	if res == "5DB3" {
		t.Logf("Success !")
	} else {
		t.Error("Expected 5DB3")
	}
}

func TestDay2(t *testing.T) {
	content, _ := ioutil.ReadFile("../inputs/day02.txt")
	text := strings.TrimSpace(string(content))
	res := part2(text)
	if res == "27CA4" {
		t.Logf("Success !")
	} else {
		t.Error("Expected 27CA4")
	}
}
