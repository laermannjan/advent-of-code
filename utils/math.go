package utils

import "cmp"

// greatest common divisor (GCD) via Euclidean algorithm
func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func GCD(integers ...int) int {
	_gcd := gcd(integers[0], integers[1])
	for i := 2; i < len(integers); i++ {
		_gcd = gcd(_gcd, integers[i])
	}
	return _gcd
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

type Comparator[T any] func(a, b T) bool

func ArgMinFunc[T any](s []T, cmp func(a, b T) int) int {
	if len(s) < 1 {
		panic("ArgMinFunc: empty list")
	}
	ii := 0
	for i, v := range s {
		if cmp(v, s[ii]) < 0 {
			ii = i
		}
	}
	return ii
}

func ArgMin[T cmp.Ordered](s []T) int {
	if len(s) < 1 {
		panic("ArgMin: empty list")
	}
	return ArgMinFunc(s, func(a, b T) int { return cmp.Compare(a, b) })
}

func ArgMaxFunc[T any](s []T, cmp func(a, b T) int) int {
	if len(s) < 1 {
		panic("ArgMaxFunc: empty list")
	}
	ii := 0
	for i, v := range s {
		if cmp(v, s[ii]) > 0 {
			ii = i
		}
	}
	return ii
}

func ArgMax[T cmp.Ordered](s []T) int {
	if len(s) < 1 {
		panic("ArgMax: empty list")
	}
	return ArgMaxFunc(s, func(a, b T) int { return cmp.Compare(a, b) })
}
