package matematica

// utilizando generics
func Sum[T int | float64](m map[string]T) T { // primeira letra maiuscula para deixar publico, se for minuscula fica privado
	var sum T

	for _, v := range m {
		sum += v
	}

	return sum
}
