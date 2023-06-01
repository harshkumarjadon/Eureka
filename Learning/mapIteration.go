package main

import (
	"fmt"
	"time"
)

func MapKeys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) GetAll() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func nestedMaps() {
	var x = map[string]map[string]string{}

	x["fruits"] = map[string]string{}
	x["colors"] = map[string]string{}

	x["fruits"]["a"] = "apple"
	x["fruits"]["b"] = "banana"

	x["colors"]["r"] = "red"
	x["colors"]["b"] = "blue"

	fmt.Println(x)

}

type rateLimit struct {
	Count int
}
type RateLimiter struct {
	cachePeriod time.Duration
	rateLimits  map[string]*rateLimit
}

func nestedMapData() {
	rateLimits := make(map[string]map[string]*rateLimit)
	rateLimits["foo"] = make(map[string]*rateLimit)
	//rateLimits["foo"]["bar"] = "a"

	//	rateLimits := map[string]map[string]*rateLimit{"foo": {"bar": a}}

}

func main() {
	employeeSalary := map[string]int{
		"steve": 12000,
		"jamie": 15000,
		"mike":  9000,
	}
	fmt.Println("Original employee salary", employeeSalary)
	modified := employeeSalary
	modified["mike"] = 18000
	fmt.Println("Employee salary changed", employeeSalary)

	//Data structure
	var m = map[int]string{1: "2", 2: "4", 4: "8"}

	fmt.Println("keys:", MapKeys(m))

	_ = MapKeys[int, string](m)

	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	fmt.Println("list:", lst.GetAll())

	nestedMaps()

}
