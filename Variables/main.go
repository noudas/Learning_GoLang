package main

import "fmt"

// Variables can be declared at the package level (outside any function)
// These variables have package-level scope, meaning they are accessible throughout the package.

var pacote string = "IM OUTSIDE!"

// If the var is attributed to a variable but not assigned a value
// The value must be added only inside the codeblock

var valorfora string

func main() {

	// The primitive variable types in Go are string, int, and bool.
	// Bools are declared as "true" or "false".

	// A way to declare a variable exclusive to Go is using the shorthand ":=", or Gopher
	// which automatically infers the type of your variable based on the value assigned.
	// They can only be used for a new variable; use only "=" to reassign a value.
	// The Gopher only works within a code block; outside of a function, it won’t work.

	x := "String" // x is a string
	y := 10       // y is an int
	z := true     // z is a bool

	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("y: %v, %T\n", y, y)
	fmt.Printf("z: %v, %T\n\n", z, z)

	// Another way to declare a variable is using the keyword "var"
	// The format is var VariableName VariableType = Value

	var istringui string = "String"
	var isinti int = 10
	var isbooli bool = true

	fmt.Printf("istringui: %v, %T\n", istringui, istringui)
	fmt.Printf("isinti: %v, %T\n", isinti, isinti)
	fmt.Printf("isbooli: %v, %T\n\n", isbooli, isbooli)

	// You can create Bulk Variables by a couple of methods
	// Using the Gopher
	// They can be declared with or withdout a value
	// with or withdout the type
	// inside the codeblock

	a, b, c := "aaaaaa", 10, false

	fmt.Printf("a: %v, %T\n", a, a)
	fmt.Printf("b: %v, %T\n", b, b)
	fmt.Printf("c: %v, %T\n\n", c, c)

	// by using var
	// They can be declared with or withdout a value
	// with or withdout the type
	// in or outside the codeblock

	var (
		my   string = "name is"
		name int    = 15
		is   bool   = false
	)

	fmt.Printf("my: %v, %T\n", my, my)
	fmt.Printf("name: %v, %T\n", name, name)
	fmt.Printf("is: %v, %T\n\n", is, is)

	valorfora = "Fora porem dentro do codeblock\n"
	fmt.Println(valorfora)

	fmt.Printf("pacote: %v, %T\n", pacote, pacote)
}
