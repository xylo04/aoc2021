package diagnostic

import (
	"strconv"
	"strings"
)

type DiagReport struct {
	Gamma        uint
	Epsilon      uint
	PowerConsume uint
	O2Gen        uint
	CO2Scrub     uint
	LifeSupport  uint
}

func CalcReport(rep string) DiagReport {
	numbers := strings.Split(rep, "\n")
	if numbers[len(numbers)-1] == "" {
		numbers = numbers[:len(numbers)-1]
	}
	report := DiagReport{}
	report.computeGammaEpsilon(numbers)
	report.computeO2Gen(numbers)
	report.computeCO2Scrub(numbers)
	report.LifeSupport = report.O2Gen * report.CO2Scrub
	return report
}

func (report *DiagReport) computeGammaEpsilon(numbers []string) {
	width := len(numbers[0])
	gammaStr := strings.Repeat(" ", width)
	epsilonStr := strings.Repeat(" ", width)
	for bit := 0; bit < width; bit++ {
		mcv := report.mostCommonValue(numbers, bit)
		if mcv == '1' {
			gammaStr = gammaStr[:bit] + "1" + gammaStr[bit+1:]
			epsilonStr = epsilonStr[:bit] + "0" + epsilonStr[bit+1:]
		} else {
			gammaStr = gammaStr[:bit] + "0" + gammaStr[bit+1:]
			epsilonStr = epsilonStr[:bit] + "1" + epsilonStr[bit+1:]
		}
	}

	gamma, _ := strconv.ParseInt(gammaStr, 2, 32)
	epsilon, _ := strconv.ParseInt(epsilonStr, 2, 32)
	report.Gamma = uint(gamma)
	report.Epsilon = uint(epsilon)
	report.PowerConsume = report.Gamma * report.Epsilon
}

func (report *DiagReport) computeO2Gen(numbers []string) {
	var remaining = make([]string, len(numbers))
	copy(remaining, numbers)
	width := len(numbers[0])
	for i := 0; i < width; i++ {
		mcv := report.mostCommonValue(remaining, i)
		filtered := []string{}
		for _, n := range remaining {
			if rune(n[i]) == mcv {
				filtered = append(filtered, n)
			}
		}
		if len(filtered) == 1 {
			o2Gen, _ := strconv.ParseInt(filtered[0], 2, 32)
			report.O2Gen = uint(o2Gen)
		}
		remaining = filtered
	}
}

func (report *DiagReport) computeCO2Scrub(numbers []string) {
	var remaining = make([]string, len(numbers))
	copy(remaining, numbers)
	width := len(numbers[0])
	for i := 0; i < width; i++ {
		mcv := report.mostCommonValue(remaining, i)
		lcv := '0'
		if mcv == '0' {
			lcv = '1'
		}
		filtered := []string{}
		for _, n := range remaining {
			if rune(n[i]) == lcv {
				filtered = append(filtered, n)
			}
		}
		if len(filtered) == 1 {
			co2Scrub, _ := strconv.ParseInt(filtered[0], 2, 32)
			report.CO2Scrub = uint(co2Scrub)
		}
		remaining = filtered
	}
}

func (report *DiagReport) mostCommonValue(numbers []string, bit int) rune {
	on := 0
	for n := 0; n < len(numbers); n++ {
		if numbers[n][bit] == '1' {
			on++
		}
	}
	mcv := '1'
	if on*2 < len(numbers) {
		mcv = '0'
	}
	return mcv
}
