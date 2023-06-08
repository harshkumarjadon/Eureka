package main

import "fmt"

type Doer[T any] interface {
	Do()
	*T
}

type Thing struct {
	used int
}

func (t *Thing) Do() {
	t.used++
	fmt.Println("Doing my thing.")
}

func Append[PT Doer[T], T any](list []T, constructor func() PT) []T {
	v := constructor()
	v.Do()
	return append(list, *v)
}

func main() {
	constructor := func() *Thing { return &Thing{} }
	s := make([]Thing, 0)
	s = Append(s, constructor)
	fmt.Println(s)
}
