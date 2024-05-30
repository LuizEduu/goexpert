package main

import (
	car "cursogo/Car"
)

//utilizando interfaces vazias
/* func Sum(m map[string]interface{}) interface{} {
	var sum int

	for _, v := range m {
		sum += v.(int)
	}

	return sum
}

func main() {

	m := map[string]interface{}{"a": 5, "b": 10}

	result := Sum(m)

	fmt.Println(result)
}
*/

func compare[T comparable](a, b T) bool {
	if a == b {
		return true
	}

	return false
}

func main() {
	/* m := map[string]int{"a": 5, "b": 10}
	m2 := map[string]float64{"a": 20.2, "b": 450.24}

	result := matematica.Sum(m)
	result2 := matematica.Sum(m2)

	fmt.Println(result)
	fmt.Println(result2) */

	car := car.Car{Marca: "Fiat"}

	car.Run()

}
