package main

import "fmt"

var globalVar = "it's considered best practice to use vars outside of functions"
var globalStringVar string // declares type, initializes with the zero value of that type
var rawStringLiteral = `A raw string literal is different from text interpolation in javascript.
I can do stuff like "\m/"`

func main() {
	fmt.Println("hello")

	// declares there is a variable shortDeclaration and assigns it a value.
	// it is considered best practice to use this inside functions, as it scopes the variable assignment
	shortDeclaration := "':='"
	fmt.Println(shortDeclaration + "is the short declaration operator.")

	// Go allows you to create your own 'types'
	type intType int
	var exampleIntType intType
	var exampleInt int

	exampleIntType = 42
	exampleInt = 43

	// fmt.Printf("%T", exampleIntType) // => main.intType

	fmt.Println(exampleInt)

	// Now, if I were to make a normal int and try to assign that to the variable exampleIntType,
	//	go will prevent this.  It must be exlicitly converted.  (Not casting, that isn't a thing in go.)
	exampleInt = int(exampleIntType) // => 42

}
