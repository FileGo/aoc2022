package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	sectionsMax = 99
)

func main() {
	buf, err := os.ReadFile("./example.dat")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(buf), "\n")
	count := 0

	for _, line := range lines {
		var start1, start2, end1, end2 int

		_, err := fmt.Sscanf(line, "%d-%d,%d-%d", &start1, &end1, &start2, &end2)
		if err != nil {
			panic(err)
		}

		if (start2 >= start1 && end2 <= end1) || (start1 >= start2 && end1 <= end2) {
			count++
			//fmt.Printf("Fully contained schedule:\n%s\n", line)
		}
	}

	fmt.Printf("Number of fully contained schedules (part 1): %d\n", count)

	// Part 2
	count = 0
	for _, line := range lines {
		var start1, start2, end1, end2 int

		_, err := fmt.Sscanf(line, "%d-%d,%d-%d", &start1, &end1, &start2, &end2)
		if err != nil {
			panic(err)
		}

		overlap := false

		if start1 == start2 || end1 == end2 {
			overlap = true
		} else if start1 < start2 {
			diffStart := start2 - start1
			if end2 >= end1+diffStart {
				overlap = true
			}
		} else if start1 > start2 { // for readability
			diffStart := start1 - start2
			if end1 >= end2+diffStart {
				overlap = true
			}
		}

		fmt.Printf("%s: %t\n", line, overlap)

		if overlap {
			count++
		}
	}

	fmt.Printf("Number of fully contained schedules (part 2): %d\n", count)
}
