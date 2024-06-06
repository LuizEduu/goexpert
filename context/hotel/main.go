package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background() //inicia um contexto em branco
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)

	defer cancel() // defer vai executar o cancel no final da execução da função bookHotel

	bookHotel(ctx)
}

func bookHotel(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("Hotel booking cancelled. Timeout reached.")
		return
	case <-time.After(6 * time.Second):
		fmt.Println("Hotel booked.")
	}
}
