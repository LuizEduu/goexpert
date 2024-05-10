package main

import "fmt"

func sum(a, b *int) int {
	*a = 24

	return *a + *b
}

func main() {
	num1 := 10
	num2 := 20

	fmt.Println(sum(&num1, &num2))

	fmt.Println(num1)
}

/* func main() {
	a := 15

	fmt.Printf("ref em memória de a %p e valor %v\n", &a, a)
	//fmt.Println(&a) & utilizado para buscar a referencia em memória da variável

	var pointer *int = &a // cria uma variável do tipo ponteiro de int e recebe a referencia em memória de a

	//fmt.Printf("ref em memória de pointer %p e valor %v\n", pointer, *pointer)

	*pointer = 20

	fmt.Printf("ref em memória de a %p e valor %v\n", &a, a)
} */
