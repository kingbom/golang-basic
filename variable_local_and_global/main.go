package main

import (
	"fmt"
)

var countGloble int = 0

func main() {
	var countLocal int = 0
	fmt.Println("########## Counting local ...")
	countLocal++
	fmt.Printf("counting local %d \n", countLocal)
	countLocal++
	fmt.Printf("counting local %d \n", countLocal)
	countLocal++
	fmt.Printf("counting local %d \n", countLocal)

	countRunGloble()
}

func countRunGloble() {
	fmt.Println("########## Counting globle ...")
	countGloble++
	fmt.Printf("counting globle %d \n", countGloble)
	countGloble++
	fmt.Printf("counting globle %d \n", countGloble)
	countGloble++
	fmt.Printf("counting globle %d \n", countGloble)
}
