package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("inputs/day04.txt")
	if err != nil {
		fmt.Println("Error reading file")
	}
	text := strings.TrimSpace(string(content))

	fmt.Println("part 1:", part1(text))
}

type Room struct {
	encryptedName string
	id            int
	checksum      string
}

func parseRoom(input string) Room {
	idx1 := strings.LastIndex(input, "-")
	idx2 := strings.Index(input, "[")

	encryptedName := input[:idx1]
	idStr := input[idx1+1 : idx2]
	id, _ := strconv.Atoi(idStr)
	checksum := input[idx2+1 : len(input)-1]

	return Room{encryptedName, id, checksum}
}

func (room Room) computeChecksum() string {
	// count chars into buckets
	charCounts := make(map[string]int)
	noDashName := strings.ReplaceAll(room.encryptedName, "-", "")
	for _, char := range strings.Split(noDashName, "") {
		charCounts[char] += 1
	}

	// bucket characters by their frequency
	charCountBuckets := make(map[int][]string)
	for char, count := range charCounts {
		charCountBuckets[count] = append(charCountBuckets[count], char)
	}

	// get list of frequencies available
	frequencies := make([]int, 0)
	for char := range charCountBuckets {
		frequencies = append(frequencies, char)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(frequencies)))

	// compute checksum using alphabetisation to break ties
	checksumLong := ""
	for _, freq := range frequencies {
		letters := charCountBuckets[freq]
		sort.Strings(letters)
		for _, letter := range letters {
			checksumLong += letter
		}
	}
	return checksumLong[:5]
}

func part1(text string) int {
	sumOfRoomIds := 0
	for _, line := range strings.Split(text, "\n") {
		room := parseRoom(line)
		if room.computeChecksum() == room.checksum {
			sumOfRoomIds += room.id
		}
	}
	return sumOfRoomIds
}
