package myutil

import "fmt"

func PrintMessage(message string) {
	fmt.Println(message)
}

func PublicMyUtilFunction() {
	fmt.Println("This function starts with capital letter thus can be used outside this package")
}

func privateMyUtilFunction() {
	fmt.Println("This function starts with small letter thus can NOT be used outside this package")
}
