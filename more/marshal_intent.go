package more

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Email    string `json:"email"`
	Password string `json:"-"`
	Address  string `json:"address,omitempty"`
}

func BeautifyJson() {

	user := User{
		Email:    "cvy@gmail.com",
		Password: "123456789",
		Address:  "Nepal",
	}

	fmt.Println(user)

	// BEAUTIFY ONLY FOR LOGGING NOT FOR API :
	data, err := json.MarshalIndent(user, "", "")

	if err != nil {
		fmt.Println(err)
	}
	// RETURNS THE []BYTE :> SLICE OF THE BYTES
	fmt.Println(data)

	// RETURNS THE DATA
	fmt.Println(string(data))

	// THEN TO UNMARHALL
	var userData User
	err = json.Unmarshal(data, &user)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(userData)

}
