package main

import (
	"fmt"
)

func main() {
	// slice()
	// removeLeading()
	// removeTrailing()
	removeSpecifiIndex()
}

func slice() {
	var numbers = make([]int, 0, 5)
	numbers = append(numbers, 1)
	numbers = append(numbers, 2)
	showIndex(numbers)

	var numbers2 []int
	numbers2 = append(numbers2, 1)
	numbers2 = append(numbers2, 2)
	showIndex(numbers2)

	var numbers3 []int = []int{1, 2, 3}
	showIndex(numbers3)
}

func removeLeading() {
	fmt.Println("removeLeading ..")
	var numbers []int = []int{1, 2, 3}
	showIndex(numbers)
	numbers = numbers[1:len(numbers)]
	showIndex(numbers)
	numbers = numbers[1:len(numbers)]
	showIndex(numbers)

}

func removeTrailing() {
	fmt.Println("removeTrailing ..")
	var numbers []int = []int{1, 2, 3}
	showIndex(numbers)
	numbers = numbers[0 : len(numbers)-1]
	showIndex(numbers)
}

func removeSpecifiIndex() {
	fmt.Println("removeSpecifiIndex ..")
	var numbers []int = []int{1, 2, 3, 4, 5, 6}
	showIndex(numbers)
	numbers = removeIndex(numbers, 2)
	showIndex(numbers)
}

func showIndex(a []int) {
	//len = length of array
	//cap = capacity of array
	fmt.Printf("len=%d , cap=%d slice=%v \n", len(a), cap(a), a)
}

func removeIndex(a []int, index int) []int {
	return append(a[:index], a[index+1:]...)
}
