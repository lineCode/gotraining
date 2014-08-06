// Sample program to show how to create values from exported types with
// embedded unexported types.
package main

import (
	"fmt"

	"github.com/ArdanStudios/gotraining/05-packaging/example6/animals"
)

// main is the entry point for the application.
func main() {
	/// Create a value of type Dog from the animals package.
	dog := animals.Dog{
		BarkStrength: 10,
	}

	// Set the exported fields from the unexported
	// animal inner type.
	dog.Name = "Chole"
	dog.Age = 1

	fmt.Printf("Counter: %#v\n", dog)
}
