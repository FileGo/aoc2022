package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	rock     = 1
	paper    = 2
	scissors = 3

	resLoss = 0
	resDraw = 3
	resWin  = 6
)

func sanitize(input []byte) string {
	ret := string(input)
	ret = strings.ReplaceAll(ret, "X", "A")
	ret = strings.ReplaceAll(ret, "Y", "B")
	ret = strings.ReplaceAll(ret, "Z", "C")

	return ret
}

func getPoint(input byte) int {
	return int(input) - 64
}

func getRes(input int) int {
	return (input - 1) * 3
}

func main() {
	buf, err := os.ReadFile("./puzzle.dat")
	if err != nil {
		panic(err)
	}

	bufS := sanitize(buf)
	score := 0

	for _, round := range strings.Split(bufS, "\n") {
		var res int

		me := getPoint(round[2])
		oppo := getPoint(round[0])

		if me == oppo {
			// Draw
			res = resDraw
		} else {
			if me == paper {
				if oppo == rock {
					res = resWin
				} else if oppo == scissors {
					res = resLoss
				}
			} else if me == rock {
				if oppo == scissors {
					res = resWin
				} else if oppo == paper {
					res = resLoss
				}
			} else if me == scissors {
				if oppo == rock {
					res = resLoss
				} else if oppo == paper {
					res = resWin
				}
			}
		}

		score += res + me
		//fmt.Printf("%d %d: %d\n", oppo, me, res+me)
	}

	fmt.Printf("Total score (part 1): %d\n", score)

	// Part 2
	score = 0
	for _, round := range strings.Split(bufS, "\n") {
		var me int
		oppo := getPoint(round[0])
		res := getRes(getPoint(round[2]))

		switch res {
		case resDraw:
			me = oppo
			break
		case resWin:
			switch oppo {
			case paper:
				me = scissors
				break
			case scissors:
				me = rock
				break
			case rock:
				me = paper
			}
		case resLoss:
			switch oppo {
			case paper:
				me = rock
				break
			case scissors:
				me = paper
				break
			case rock:
				me = scissors
			}
		}

		score += res + me
		//fmt.Printf("%d %d: %d\n", oppo, me, res+me)
	}

	fmt.Printf("Total score (part 2): %d\n", score)

}
