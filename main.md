package main

import (
"encoding/json"
"fmt"
)

type Person struct {
Name string `json:"name,omitempty"`
Age string `json:"age"`
Password string `json:"-"`
IsAdult bool `json:"isAdult"`
}

func main() {

    person := Person{
    	Name:     "John",
    	Age:      "33",
    	Password: "Password",
    	IsAdult:  true,
    }

    fmt.Println("We are learning JSON!! ")

    fmt.Println(person)

    // CONVERTING THE PERSON INTO JSON ENCODING (MARSHALL) :-> GO OBJECT -> JSON
    data, err := json.Marshal(person)
    if err != nil {
    	fmt.Println(fmt.Errorf("error while marshalling: %w", err))
    } else {
    	// slice of byte
    	// [ 123 34 110 97 109 ...... ]
    	// 123 :-> {.            // (ASCII!!!!)
    	// 34 :-> "              // (ASCII !!!! )
    	fmt.Println("DATA: ", data)

    	// BYTES TO STRING !!!!
    	fmt.Println("DATA: ", string(data))

    }

    //  UNMARHSALLING :-> JSON BUFFER TO GO OBJECT

    var personData Person

    err = json.Unmarshal(data, &personData)
    if err != nil {
    	fmt.Println(fmt.Errorf("error while unmarshilling: %w", err))
    }
    fmt.Println("DATA: ", personData)

}
