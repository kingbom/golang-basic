package main

import (
	"fmt"
)

func main() {
	//Explicit Declaration
	//Can change value
	var tmp1 int = 0
	var tmp2 string = "Hello"
	var tmp3 bool = true
	tmp3 = false

	//Implicit Declaration
	//Can change value
	tmp4 := "Xxx"
	tmp5 := 0
	tmp6 := true
	tmp6 = true

	//Can't change final value
	const tmp7 = 10

	fmt.Println(tmp1)
	fmt.Println(tmp2)
	fmt.Println(tmp3)
	fmt.Println(tmp4)
	fmt.Println(tmp5)
	fmt.Println(tmp6)
	fmt.Println(tmp7)
}
