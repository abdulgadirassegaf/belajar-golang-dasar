package main

import "fmt"

func main() {
	const (
		firstName = "Abdul Gadir"
		lastName  = "Assegaf"
	)
	fmt.Println(firstName)
	fmt.Println(lastName)
	// error
	// firstName = "Siapa"
	// lastName = "Siapa ini"
}
