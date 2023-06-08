package main

import "fmt"

type Crud1[T any] interface {
	Insert(T)
	Update(T)
}

type dummy string

func (d dummy) Insert[T dummy](r T) {
	fmt.Println("hello world")
}

func (d dummy) Update[T dummy](r T) {
	fmt.Println("hello world")
}

type PersonData interface {
	WorkData()
}

type workerData string

func (w workerData) WorkData() {
	fmt.Printf("%s is working\n", w)
}

// func (c GenCategory) insert(prod *GenProduct)Println(a ...any)
// 	c.dummyProduct = append(c.dummyProduct, *prod)
// }

func main() {
	//prod := GenProduct{id: 123}
	//prod.insert(prod)
	//fmt.Println()
	//cat := GenCategory{id: 1, name: "Hello World"}

	// cat.insert(*GenProduct{id: 1, name: "TV", price: 10})
	// cat.insert(*GenProduct{id: 2, name: "TV New", price: 100})

	// cat.insert(&prod)
	// prod.insert(&cat)
}
