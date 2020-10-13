package main

import (
	"fmt"
)

func main() {
	forLoop()
	forWhileLoop()

}

func forLoop() {
	for i := 0; i <= 10; i++ {
		fmt.Printf("For loop i = %d \n", i)
	}
}

func forWhileLoop() {
	index := 0
	for index <= 5 {
		fmt.Printf("For while loop i = %d \n", index)
		index++
	}
}
