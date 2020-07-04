package main

import (
	"encoding/json"
	"fmt"

	simplepb "github.com/renishb10/goprotobuf/src/simple"
)

func main() {
	fmt.Println("Hello World")
	s := simplepb.Simple{
		FirstName: "Renish",
		LastName:  "Bhaskaran",
		Age:       31,
	}
	v, _ := json.Marshal(s)
	fmt.Println(string(v))

	fmt.Println("Get FirstName", s.GetFirstName)
}
