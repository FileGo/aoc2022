package main

import (
	"fmt"
	"os"
	"strings"
)

func getPriority(input byte) int {
	ascii := int(input)

	if ascii >= 97 && ascii <= 122 {
		return ascii - 96
	} else if ascii >= 65 && ascii <= 90 {
		return ascii - 64 + 26
	}

	return -1
}

func contains(needle byte, haystack []byte) bool {
	for _, item := range haystack {
		if needle == item {
			return true
		}
	}

	return false
}

func main() {
	buf, err := os.ReadFile("./puzzle.dat")
	if err != nil {
		panic(err)
	}

	total := 0
	rucksacks := strings.Split(string(buf), "\n")

	for _, rucksack := range rucksacks {
		first := rucksack[:len(rucksack)/2]
		second := rucksack[len(rucksack)/2:]

		sum := 0
		var commons []byte

		for i := 0; i < len(first); i++ {
			for j := 0; j < len(second); j++ {
				if first[i] == second[j] {
					if !contains(first[i], commons) {
						sum += getPriority(first[i])
						commons = append(commons, first[i])
					}
				}
			}
		}

		total += sum
	}

	fmt.Printf("Sum of priorities (part 1): %d\n", total)

	// Part 2
	nGroups := len(rucksacks) / 3
	total = 0

	for g := 0; g < nGroups; g++ {
		sum := 0
		var badges []byte
		r := g * 3
		for i := 0; i < len(rucksacks[r]); i++ {
			elem := rucksacks[r][i]
			if contains(elem, []byte(rucksacks[r+1])) && contains(elem, []byte(rucksacks[r+2])) {
				if !contains(elem, badges) {
					sum += getPriority(elem)
					badges = append(badges, elem)
				}
			}
		}

		total += sum
	}

	fmt.Printf("Sum of priorities (part 2): %d\n", total)
}
