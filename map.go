package main

import "fmt"

func main() {
	// var person map[string]string = map[string]string{}
	// person["name"]="Abdul Gadir"
	// person["address"]="Cirebon"

	person := map[string]string{
		"name":    "Abdul Gadir",
		"address": "Cirebon",
	}
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person)

	book := make(map[string]string)
	book["title"] = "Buku Golang"
	book["author"] = "Ading"
	book["ups"] = "Salaaaah"

	fmt.Println(book)

	delete(book, "ups")

	fmt.Println(book)
}
