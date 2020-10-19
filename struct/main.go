package main

import "fmt"

//short-cut key for new tsy
type product struct {
	name  string
	price int
	stock int
}

func main() {
	var p product
	p.name = "Book"
	p.price = 100
	p.stock = 20
	show(p)
}

func show(p product) {
	fmt.Println("show product : ", p)
	fmt.Println("show product name : ", p.name)
	fmt.Println("show product price : ", p.price)
	fmt.Println("show product stock : ", p.stock)
}
