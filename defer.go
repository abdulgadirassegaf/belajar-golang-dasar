package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplicaton() {
	defer logging()
	fmt.Println("Run application")
}

func main() {
	runApplicaton()
}
