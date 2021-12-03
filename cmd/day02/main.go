package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/xylo04/aoc2021/internal/pilot"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	cmds := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		cmds = append(cmds, scanner.Text())
	}

	x, y := pilot.Parse(cmds)
	fmt.Printf("Ending position was %d, %d; mult is %d", x, y, x*y)
}
