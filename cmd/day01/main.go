package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/xylo04/aoc2021/internal/depth"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	depths := make([]uint, 0)
	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		dep, _ := strconv.ParseUint(scanner.Text(), 10, 32)
		depths = append(depths, uint(dep))
	}

	inc := depth.Increases(depths)
	fmt.Printf("There were %d increases in depth", inc)
}
