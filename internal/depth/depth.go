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

const windowSize = 3

func IncreaseWindowSum(d []uint) uint {
	if len(d) < windowSize+1 {
		return 0
	}
	var c = uint(0)
	lastSum := sum(d[0:windowSize])
	for i := 1; i <= len(d)-windowSize; i++ {
		thisSum := sum(d[i : i+windowSize])
		if lastSum < thisSum {
			c++
		}
		lastSum = thisSum
	}
	return c
}

func sum(uints []uint) uint {
	acc := uint(0)
	for _, v := range uints {
		acc += v
	}
	return acc
}
