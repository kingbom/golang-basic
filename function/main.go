package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello function")
	fnNoParameter()
	fnHasParameter("Test function")
}

func fnNoParameter() {
	fmt.Println("function no parameter")
}

func fnHasParameter(massage string) {
	fmt.Printf("function %s\n", massage)
}
