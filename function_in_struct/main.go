package main

import "fmt"

type product struct {
	name  string
	price int
	stock int
}

func (p product) clear() product {
	p.price = 0
	p.stock = 0
	return p
}

func (p product) setDiscount(d int) product {
	p.price = p.price - d
	return p
}

func main() {
	var p product
	p.name = "Book"
	p.price = 100
	p.stock = 20
	fmt.Println("show product ...")
	show(p)
	p = p.setDiscount(10)
	fmt.Println("show product after discount...")
	show(p)
	p = p.clear()
	fmt.Println("show product after clear  ...")
	show(p)
}

func show(p product) {
	fmt.Println("show product : ", p)
}
