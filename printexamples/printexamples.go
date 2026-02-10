package printexamples

import "fmt"

func PrintExamples() {
	age := 21
	name := "sankalp"
	height := 6.3434132

	// print using print line
	// adds space after each arguement
	fmt.Println("Name: ", name, "Age: ", age, "Height: ", height)
	// adds new line character
	fmt.Println(" hi ")

	// print using formatter
	// can be done in single line as well
	fmt.Printf("Name: %s\n", name)
	fmt.Printf("Age: %d", age)
	fmt.Printf("Height: %.3f\n", height)
	fmt.Printf("oops !! forgot the new line character after age\n")
	fmt.Printf("type of hight is %T\n", height)

}
