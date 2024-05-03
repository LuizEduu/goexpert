package main

import "fmt"

/*
	Slices é uma estrutura de dados semelhantes ao array
	A diferença é que não é necessário definir um tamanho fixo em sua inicialização
	Slices por baixo dos panos utiliza array
	Eles são divididos em 3 partes
	Um ponteiro que "aponta" para o array que vai ser utilizado
	Um tamanho para saber até onde ir
	Uma capacidade para saber o quanto ele consegue armazenar de dados
*/

func main() {
	s := []int{20, 10, 45, 4, 5, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2}

	//fmt.Printf("Len=%d cap=%d %vn", len(s), cap(s), s)

	//fmt.Printf("Len=%d cap=%d %v", len(s[:0]), cap(s[:0]), s[:0]) // do inicio do meu slice para a direita, mostre 0 items
	//fmt.Printf("Len=%d cap=%d %v", len(s[:4]), cap(s[:4]), s[:4]) // do inicio do meu slice para a direita, mostre 4 items

	fmt.Printf("Len=%d cap=%d %v", len(s[2:]), cap(s[2:]), s[2:]) // vai começar a conta da posição 3 até o fim do meu slice

	/*
		:4 não altera a capacidade do slice, apenas ignora os outros elementos
		2: altera a capacidade do slice, vai ser 2 elementos a menos
	*/

	s = append(s, 240) //vai pegar o tamanho inicial do slice e dobrar a capacidade em um array novo
}
