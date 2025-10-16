package xmath

// Sign returns -1 for negative numbers, 1 for positive numbers, and 0 for zero.
// Example: direction := Sign(velocity.X)
func Sign[T ~float64 | ~int](x T) T {
	var zero T
	if x == zero {
		return zero
	}
	if x < zero {
		return -1
	}
	return 1
}

func Clamp[T ~float64 | ~int](value, lower, upper T) T {
	if value < lower {
		return lower
	}
	if value > upper {
		return upper
	}
	return value
}

func Abs[T ~float64 | ~int](v T) T {
	if v < 0 {
		return -v
	}
	return v
}
