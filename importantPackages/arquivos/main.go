package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	f, err := os.Create("arquivo.txt")

	defer f.Close()

	if err != nil {
		panic(err)
	}

	length, err := f.Write([]byte("Escrevendo no arquivo2222..."))

	if err != nil {
		panic(err)
	}

	fmt.Printf("escrita no arquivo com sucesso bytes %d \n", length)

	file2, err := os.Open("arquivo.txt")

	defer file2.Close()

	reader := bufio.NewReader(file2)
	buffer := make([]byte, 4)

	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}

		fmt.Println(string(buffer[:n]))
	}

}
