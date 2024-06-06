package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		log.Println("Request started")

		defer log.Println("Request finished")

		select {
		case <-time.After(5 * time.Second):
			log.Println("Request process sucessfull")

			w.Write([]byte("Request process sucessfull"))

		case <-ctx.Done():
			log.Println("Request aborted")
		}
	})
	http.ListenAndServe(":8080", nil)
}
