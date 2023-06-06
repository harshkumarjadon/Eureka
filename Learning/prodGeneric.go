package main

import (
	"fmt"

)

type crud[T, K any] interface {
	insert(T, K)
	update(int, T, K)
}

type Product struct {
	id    float64
	name  string
	price float64
}

type Category struct {
	id      float64
	name    string
	product []Product
}

func (c Product) insert(cat *Category) {
	cat.product = append(cat.product, c)
}

func (c Product) update(upd int, cat *Category) {
	//cat.id = float64(upd)
	for i, v := range cat.product {
		if v.id == float64(upd) {
			cat.product[i] = c
		}
	}
}

func (c Category) insert(prod *Product) {
	c.product = append(c.product, *prod)
}

func (c Category) update(upd int, cat *Category) {
	cat.id = float64(upd)
	cat.name = "Hellllllo"
}

func main() {
	cat := Category{id: 1, name: "Hello World"}
	prod := Product{id: 123, name: "New Product", price: 1000}

	// cat.insert(*Product{id: 1, name: "TV", price: 10})
	// cat.insert(*Product{id: 2, name: "TV New", price: 100})

	cat.insert(&prod)
	prod.insert(&cat)

	fmt.Println(cat)
	fmt.Println("Hello world!")

	prod1 := Product{id: 123, name: "PPPPP PPPP", price: 9999}
	prod1.update(123, &cat)
	fmt.Println(cat)
}
