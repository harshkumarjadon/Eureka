package main

import (
	"fmt"

)

type crud1[T, K any] interface {
	insert(T, K)
	update(int, T, K)
}

type GenProduct struct {
	id    float64
	name  string
	price float64
}

type GenCategory struct {
	id           float64
	name         string
	dummyProduct []GenProduct
}

func (c GenProduct) insert(cat *GenCategory) {
	cat.dummyProduct = append(cat.dummyProduct, c)
}

func (c GenProduct) update(upd int, cat *GenCategory) {
	//cat.id = float64(upd)
	for i, v := range cat.dummyProduct {
		if v.id == float64(upd) {
			cat.dummyProduct[i] = c
		}
	}
}

func (c GenCategory) insert(prod *GenProduct) {
	c.dummyProduct = append(c.dummyProduct, *prod)
}

func (c GenCategory) update(upd int, cat *GenCategory) {
	cat.id = float64(upd)
	cat.name = "Hellllllo"
}

func main() {
	cat := GenCategory{id: 1, name: "Hello World"}
	prod := GenProduct{id: 123, name: "New GenProduct", price: 1000}

	// cat.insert(*GenProduct{id: 1, name: "TV", price: 10})
	// cat.insert(*GenProduct{id: 2, name: "TV New", price: 100})

	cat.insert(&prod)
	prod.insert(&cat)

	fmt.Println(cat)
	fmt.Println("Hello world!")

	prod1 := GenProduct{id: 123, name: "PPPPP PPPP", price: 9999}
	prod1.update(123, &cat)
	fmt.Println(cat)
}
