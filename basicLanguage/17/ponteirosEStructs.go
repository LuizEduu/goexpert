package main

import "fmt"

type Client struct {
	name string
}

func (c Client) run() {
	fmt.Printf("O cliente %v andou\n", c.name)
}

func main() {
	client := Client{
		name: "Luiz Eduardo",
	}

	client.run()
}
