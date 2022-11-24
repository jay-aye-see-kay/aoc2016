package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
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

func part1(text string) int {
	sumOfRoomIds := 0

	for _, line := range strings.Split(text, "\n") {
		// splitLine on `[` and remove trailing `]` -> have the checksum
		splitLine := strings.Split(line, "[")
		expectedChecksum := strings.ReplaceAll(splitLine[1], "]", "")

		re1 := regexp.MustCompile(`\d+`)
		roomId, _ := strconv.Atoi(re1.FindString(line))

		// remove non-alpha chars
		re2 := regexp.MustCompile("[^a-z]")
		ename := re2.ReplaceAllString(splitLine[0], "")

		// count chars into buckets
		charCounts := make(map[string]int)
		for _, char := range strings.Split(ename, "") {
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

		actualChecksum := checksumLong[:5]

		if actualChecksum == expectedChecksum {
			sumOfRoomIds += roomId
		}
	}

	return sumOfRoomIds
}
