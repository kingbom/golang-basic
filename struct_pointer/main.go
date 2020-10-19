package main

import "fmt"

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
	fmt.Println("show product before update ...")
	show(p)
	update(&p)
	fmt.Println("show product after update  ...")
	show(p)
}

func show(p product) {
	fmt.Println("show product : ", p)
}

func update(p *product) {
	p.price = p.price + 100
	p.stock = 10
}
