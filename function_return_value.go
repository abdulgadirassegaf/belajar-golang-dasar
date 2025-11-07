package main

import "fmt"

func getHello(name string) string {
	hello := "Hello " + name
	return hello
}

func main() {
	result := getHello("Gadir")
	fmt.Println(result)

	fmt.Println(getHello("Assegaf"))
	fmt.Println(getHello("Abdul"))
}
