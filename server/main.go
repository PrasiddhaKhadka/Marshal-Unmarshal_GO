package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/person", personHandler)

	fmt.Println("Server running on :8080")

	http.ListenAndServe(":8085", nil)
}
