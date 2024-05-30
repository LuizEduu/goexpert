package car

import "fmt"

type Car struct {
	Marca string
}

func (c Car) Run() {
	fmt.Println("Andou.....")
}
