package main

import "fmt"

func main() {

	value := func() int {
		return 1 + 2
	}() //função anonima, auto executável

	fmt.Println(value)

}
