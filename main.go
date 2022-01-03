package main

import "fmt"

func main() {
	var e Employee

	e.SetName("Rafli")
	e.SetAddress("Cambridge, Desa 10 RT/RW 100")
	e.SetAge("19")

	fmt.Println("Your Name: ", e.GetName())
	fmt.Println("Your Address: ", e.GetAddress())
	fmt.Println("Your Age: ", e.GetAge())
}
