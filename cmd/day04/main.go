package main

import (
	"fmt"
	"log"
	"os"

	"github.com/xylo04/aoc2021/internal/bingo"
)

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	game := bingo.ParseGame(string(file))
	winningBoard := game.Play()
	fmt.Printf("%+v", winningBoard)
}
