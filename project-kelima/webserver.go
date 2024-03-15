package main

import (
	"net/http"
	"fmt"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("ini dari fmt println")
		w.Write([]byte("hello world"))
	})

	http.ListenAndServe("localhost:8080", nil)
}