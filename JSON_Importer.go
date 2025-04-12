package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Username string `json:"username"`
	Role     string `json:"role"`
}

func main() {
	// var jsonHardCodedVal json I dont know how to define it ?? I used :=
	// jsonHardCodedVal := `{"username":"Alice","role":"Developer"},{"username":"Sam","role":"Administrator"},{"username":"Lisa","role":"SRE Engineer"}`
	// dont know if I have like a list of json array or something like I mentioned above

	// jsonHardCodedVal := `{"username":"Alice","role":"Developer"}`
	jsonHardCodedVal := `[{"username":"Alice","role":"Developer"},{"username":"Sam","role":"Administrator"},{"username":"Lisa","role":"SRE Engineer"}]`

	var v []User
	err := json.Unmarshal([]byte(jsonHardCodedVal), &v)

	if err != nil {
		panic(err)
	}

	for _, val := range v {
		fmt.Println(val.Username)
		fmt.Println(val.Role)
	}

}
