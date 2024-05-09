package main

import "fmt"

func main() {
	total := sum(1, 2, 3, 5, 6, 7, 4)

	fmt.Println(total)
}

func sum(numbers ...int) int {
	total := 0
	for _, value := range numbers {
		total += value
	}

	return total
}
