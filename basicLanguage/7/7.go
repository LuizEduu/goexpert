package main

import "fmt"

func main() {
	/* salarios := map[string]int{"Luiz": 10000, "João": 7200}
	//fmt.Println(salarios["Luiz"])
	delete(salarios, "João")
	salarios["Jo"] = 5200
	fmt.Println(salarios) */

	salario := make(map[string]float32)

	salario["Luiz"] = 4.200

	fmt.Println(salario)
}
