package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func Post() {
	user := Person{
		Name:    "Ram",
		Age:     28,
		IsAdult: true,
	}

	// Convert struct → JSON
	jsonData, _ := json.Marshal(user)

	// Send POST request
	resp, err := http.Post(
		"http://localhost:8080/person",
		"application/json",
		bytes.NewBuffer(jsonData),
	)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	//“When this function ends, CLOSE the body”

	defer resp.Body.Close()

	// Read response
	var result map[string]string
	json.NewDecoder(resp.Body).Decode(&result)
	fmt.Println("Response from server:", result)
}
