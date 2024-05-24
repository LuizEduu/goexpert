package matematica

// utilizando generics
func Sum[T int | float64](m map[string]T) T {
	var sum T

	for _, v := range m {
		sum += v
	}

	return sum
}
