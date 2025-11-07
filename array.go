package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Abdul"
	names[1] = "Gadir"
	names[2] = "Assegaf"

	fmt.Println(names)
	fmt.Println(names[0])

	var values = [...]int{
		82, 71, 59, 100, 32,
	}
	fmt.Println(values)
}
