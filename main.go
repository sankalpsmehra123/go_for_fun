package main // main package entry package of go lang

import (
	"fmt"
)

func main() { // main method is the entry point in the program should always be in main package.
	fmt.Println("hello go for fun!!!")

	// calling methods form different class
	//myutil.PrintMessage("hello go for fun!!! from myutil.PrintMessage() method")

	// variables
	var name string = "Sankalp"
	var version = "version latest"
	var money int = 50000
	var currency float64 = 97.53
	var istrue bool = false
	istrue = true // allowed
	fmt.Println(name)
	fmt.Println(version)
	fmt.Println(money)
	fmt.Println("Currency in $ is: ", currency)
	fmt.Println(istrue)

	// another way to declare a variable
	person := "Prince" // used when a data is returned from a fuction for which type is not known and we want to save it in a variable.
	fmt.Println(person)

	// constants cannot change value later
	const pi = 3.141
	// pi = 3.142 // not allowed
	fmt.Println(pi)

}
