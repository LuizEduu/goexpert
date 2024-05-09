package main

func main() {
	result, result2 := sum(5, 5)

	println(result, result2)
}

func sum(a int, b int) (int, bool) {
	return a + b, true
}
