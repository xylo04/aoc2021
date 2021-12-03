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

	rep := diagnostic.CalcReport(string(file))
	fmt.Printf("%+v", rep)
}
