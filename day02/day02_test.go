package main

import (
	"fmt"
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
		fmt.Printf("DEBUGPRINT[5]: day02_test.go:25: res=%+v\n", res)
		t.Error("Expected 1985")
	}
}
