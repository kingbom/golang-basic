package main

import (
	"fmt"
)

func main() {
	index := 0
	for true {
		if index > 10 {
			break
		}
		fmt.Printf("For while loop i = %d \n", index)
		index++
	}
}
