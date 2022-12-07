package main

import (
	"fmt"
	"os"
)

func do(input []byte, num int) int {
	for i := num; i < len(input); i++ {
		matchFound := true

		for j := i - 1; j >= i-num; j-- {
			if input[i] == input[j] {
				matchFound = false
				break
			}

			for k := j - 1; k >= i-num; k-- {
				if input[j] == input[k] {
					matchFound = false
					break
				}
			}

			if !matchFound {
				break
			}
		}

		if matchFound {
			return i + 1
		}
	}

	return -1
}

func main() {
	buf, err := os.ReadFile("./puzzle.dat")
	if err != nil {
		panic(err)
	}

	// Part1
	numChar := 4 - 1
	res := do(buf, numChar)
	fmt.Printf("Result (part 1): %d\n", res)

	// Part 2
	numChar = 14 - 1
	res = do(buf, numChar)
	fmt.Printf("Result (part 2): %d\n", res)

}
