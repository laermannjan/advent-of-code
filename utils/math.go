package utils

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(integers ...int) int {
	result := integers[0]
	for i := 1; i < len(integers); i++ {
		cur := integers[i]
		result = result * cur / GCD(result, cur)
	}

	return result
}
