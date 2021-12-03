package pilot

import (
	"strconv"
	"strings"
)

func Parse(cmds []string) (uint, uint) {
	// horizontal position
	var x uint
	// depth
	var y uint
	for _, c := range cmds {
		dir := c[:strings.Index(c, " ")]
		magStr := c[strings.Index(c, " ")+1:]
		mag, _ := strconv.ParseUint(magStr, 10, 32)

		switch dir {
		case "forward":
			x += uint(mag)
		case "up":
			// up means depth is decreasing
			y -= uint(mag)
		case "down":
			// down means depth is increasing
			y += uint(mag)
		}
	}
	return x, y
}
