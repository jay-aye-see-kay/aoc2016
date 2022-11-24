package main

import (
	"io/ioutil"
	"strings"
	"testing"
)

func TestSample1(t *testing.T) {
	sample := `aaaaa-bbb-z-y-x-123[abxyz]
a-b-c-d-e-f-g-h-987[abcde]
not-a-real-room-404[oarel]
totally-real-room-200[decoy]`

	text := strings.TrimSpace(sample)
	res := part1(text)
	if res == 1514 {
		t.Logf("Success !")
	} else {
		t.Error("Expected 1514")
	}
}

func TestDay1(t *testing.T) {
	content, _ := ioutil.ReadFile("../inputs/day04.txt")
	text := strings.TrimSpace(string(content))
	actual := part1(text)
	expected := 245102
	if actual != expected {
		t.Errorf("Expected %+v but recienved %+v", expected, actual)
	}
}
