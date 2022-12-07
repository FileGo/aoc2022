package main

import "os"

func main() {
	buf, err := os.ReadFile("example.dat")
	if err != nil {
		panic(err)
	}

	bufS := string(buf)
}
