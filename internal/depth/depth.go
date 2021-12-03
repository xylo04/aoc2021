package depth

func Increases(d []uint) uint {
	if len(d) < 2 {
		return 0
	}
	var c = uint(0)
	for i := 1; i < len(d); i++ {
		if d[i-1] < d[i] {
			c++
		}
	}
	return c
}
