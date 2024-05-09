package main

import "fmt"

/*
	Struct não são classes, GO não é orientada a objetos
*/

type Address struct {
	City   string
	State  string
	Number int32
}

type Client struct {
	Name   string
	Age    int32
	Active bool
	Address
}

func (client *Client) Deactivate() { // * indica que o meu valor é um ponteiro que vai direcionar para o endereço em memória de quem chamar o método e alterar o Active para false
	client.Active = false
}

func main() {
	person := Client{
		Name:   "Luiz",
		Age:    26,
		Active: true,
	}

	fmt.Println(person.Active)
	person.Deactivate()
	fmt.Println(person.Active)

	/* person.City = "Recife"
	fmt.Println(person.City) */
}
