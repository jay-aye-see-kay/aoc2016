package main

import (
	"io/ioutil"
	"strings"
	"testing"
)

func TestParseRoom(t *testing.T) {
	sample1 := "aaaaa-bbb-z-y-x-123[abxyz]"
	actual1 := parseRoom(sample1)
	expected1 := Room{"aaaaa-bbb-z-y-x", 123, "abxyz"}

	if actual1 != expected1 {
		t.Errorf("Expected %+v but received %+v", expected1, actual1)
	}

	sample2 := "not-a-real-room-404[oarel]"
	actual2 := parseRoom(sample2)
	expected2 := Room{"not-a-real-room", 404, "oarel"}

	if actual2 != expected2 {
		t.Errorf("Expected %+v but received %+v", expected2, actual2)
	}
}

func TestSample1(t *testing.T) {
	sample := `aaaaa-bbb-z-y-x-123[abxyz]
a-b-c-d-e-f-g-h-987[abcde]
not-a-real-room-404[oarel]
totally-real-room-200[decoy]`

	actual := part1(sample)
	expected := 1514
	if actual != expected {
		t.Errorf("Expected %+v but received %+v", expected, actual)
	}
}

func TestDay1(t *testing.T) {
	content, _ := ioutil.ReadFile("../inputs/day04.txt")
	text := strings.TrimSpace(string(content))
	actual := part1(text)
	expected := 245102
	if actual != expected {
		t.Errorf("Expected %+v but received %+v", expected, actual)
	}
}

func TestShiftCipher(t *testing.T) {
	actual := shiftCipher("qzmt-zixmtkozy-ivhz", 343)
	expected := "very encrypted name"
	if actual != expected {
		t.Errorf("Expected \"%+v\" but received \"%+v\"", expected, actual)
	}
}

func TestDay2(t *testing.T) {
	content, _ := ioutil.ReadFile("../inputs/day04.txt")
	text := strings.TrimSpace(string(content))
	actual := part2(text)
	expected := 324
	if actual != expected {
		t.Errorf("Expected %+v but received %+v", expected, actual)
	}
}
