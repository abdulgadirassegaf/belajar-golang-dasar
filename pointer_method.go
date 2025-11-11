package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	gadir := Man{"Gadir"}
	gadir.Married()

	fmt.Println(gadir.Name)
}
