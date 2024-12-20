package helpers

// handles only positove coordinates


func szudzikPair(x int, y int) int {
	if x >= y {
		return (x * x) + x + y
	}
	return (y * y) + x
}


//handles negative numbers and big sets too 
func szudzikPairSigned(x int, y int) int {
	c := szudzikPair(convert(x), convert(y))

	if ((x >= 0) && (y < 0)) || ((x < 0) && (y >= 0)) {
		return -c - 1
	}

	return c
}

func convert(a int) int {
	if a >= 0 {
		return 2 * a
	}
	return (2 * a) * -1
}
