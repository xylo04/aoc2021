package diagnostic

import (
	"strconv"
	"strings"
)

func Report(rep string) (uint, uint) {
	numbers := strings.Split(rep, "\n")
	if numbers[len(numbers)-1] == "" {
		numbers = numbers[:len(numbers)-1]
	}
	width := len(numbers[0])
	gammaStr := strings.Repeat(" ", width)
	epsilonStr := strings.Repeat(" ", width)
	for bit := 0; bit < width; bit++ {
		on := 0
		for n := 0; n < len(numbers); n++ {
			if numbers[n][bit] == '1' {
				on++
			}
		}
		if on*2 > len(numbers) {
			gammaStr = gammaStr[:bit] + "1" + gammaStr[bit+1:]
			epsilonStr = epsilonStr[:bit] + "0" + epsilonStr[bit+1:]
		} else {
			gammaStr = gammaStr[:bit] + "0" + gammaStr[bit+1:]
			epsilonStr = epsilonStr[:bit] + "1" + epsilonStr[bit+1:]
		}
	}

	gamma, _ := strconv.ParseInt(gammaStr, 2, 32)
	epsilon, _ := strconv.ParseInt(epsilonStr, 2, 32)
	return uint(gamma), uint(epsilon)
}
