package main

import "fmt"

func main() {
	// fmt.Println prints the provided string to the console.
	// It returns the number of bytes written and any write error encountered.
	// <nil> is the same as Null
	numerodebytes, errors := fmt.Println("Hello World!")
	fmt.Println(numerodebytes, errors)

	// You can also use "_" as a way to create an "discart string".

	_, errinhos := fmt.Println("_ test")
	fmt.Println(errinhos)
}
