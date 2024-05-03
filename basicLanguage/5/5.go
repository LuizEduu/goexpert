package main

import "fmt"

/*
Arrays é uma estrutura de dados de tamanho fixo cujo
é alocado na memória os espaços necessários que são definidos na sua declaração
por sua vez ele salva os elementos sequencialmente "um ao lado do outro"
*/

func main() {
	var myArr [3]int

	for i := range myArr {
		myArr[i] = i
	}

	for _, v := range myArr {
		fmt.Println(v)
	}

}
