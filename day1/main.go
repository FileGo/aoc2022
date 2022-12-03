package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	buf, err := os.ReadFile("./puzzle.dat")
	if err != nil {
		panic(err)
	}

	elfI := 0
	elves := make(map[int]int)

	for _, line := range strings.Split(string(buf), "\n") {
		if len(line) == 0 {
			elfI++
			continue
		}

		cal, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		elves[elfI] += cal
	}

	maxCal := 0

	for _, cal := range elves {
		if cal > maxCal {
			maxCal = cal
		}
	}

	fmt.Printf("Maximum calories carried by a single elf is: %d\n", maxCal)

	elvesSlice := make([]int, len(elves))
	for _, elf := range elves {
		elvesSlice = append(elvesSlice, elf)
	}

	sort.Slice(elvesSlice, func(i, j int) bool {
		return elvesSlice[i] > elvesSlice[j]
	})

	threeElves := 0
	for i := 0; i < 3; i++ {
		threeElves += elvesSlice[i]
	}

	fmt.Printf("Three top elves are carrying %d callories.\n", threeElves)

}
