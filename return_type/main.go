package main

import (
	"fmt"
)

func main() {
	fmt.Printf("sum  %d+%d = %d \n", 2, 3, sum(2, 3))

	a, b := swap(1, 0)
	fmt.Printf("swap %d,%d => %d,%d\n", 1, 0, a, b)

	x, y := swapNamingReturn(10, 20)
	fmt.Printf("swap spacific naming return %d,%d => %d,%d \n", 10, 20, x, y)
}

func sum(a int, b int) int {
	return a + b
}

func swap(a int, b int) (int, int) {
	return b, a
}

func swapNamingReturn(a int, b int) (x int, y int) {
	y = a
	x = b
	return x, y
}
