package integermath

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Relu(x int) int {
	if x < 0 {
		return 0
	}
	return x
}
