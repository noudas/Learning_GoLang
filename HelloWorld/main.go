/*
Every Go file must belong to a "package".
The "main package" is special because is a standalone executable program.
*/
package main

/*
Import the standard library for Go
The "fmt package" provides I/O functions,
including formatted output (Print, Println, Printf, etc.) and formatted input (Scan, Scanln, Scanf, etc.).
*/
import "fmt"

/*
Creates the main function, imports the "fmt" library into the code to make the "Println"
*/
func main() {
	// Where your programme starts
	fmt.Println("Hello World!")
	// Where your programme finishes
}
