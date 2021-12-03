package main

import (
	"fmt"
	"log"
	"os"

	"github.com/xylo04/aoc2021/internal/diagnostic"
)

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	gamma, epsilon := diagnostic.Report(string(file))
	fmt.Printf("Gamma %d, epsilon %d; power consumption %d", gamma, epsilon, gamma*epsilon)
}
