package main

import "fmt"

type PersonalDetails struct {
	name, surname string
}

func main() {
	john := PersonalDetails{name: "John", surname: "Doe"}
	fmt.Printf("Hello %+v,", john)
}
