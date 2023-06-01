package main

import "fmt"

type Number interface {
	int64 | float64
}

// T is a type parameter that is used like normal type inside the function
// any is a constraint on type i.e T has to implement "any" interface
func reverse[T any](s []T) []T {
	l := len(s)
	r := make([]T, l)

	for i, ele := range s {
		r[l-i-1] = ele
	}
	return r
}

// ForEach on slice, that will execute a function on each element of slice.
func ForEach[T any](s []T, f func(ele T, i int, s []T)) {
	for i, ele := range s {
		f(ele, i, s)
	}
}

type Person interface {
	Work()
}

type worker string

func (w worker) Work() {
	fmt.Printf("%s is working\n", w)
}

func DoWork(things []worker) {
	for _, v := range things {
		v.Work()
	}
}

func DoWorkGeneric[T Person](things []T) {
	for _, v := range things {
		v.Work()
	}
}

func main() {
	/*
		// // Initialize a map for the integer values

		// ints := map[string]int64{
		// 	"first":  34,
		// 	"second": 12,
		// }

		// // Initialize a map for the float values
		// floats := map[string]float64{
		// 	"first":  35.98,
		// 	"second": 26.99,
		// }

		// fmt.Printf("Non-Generic Sums: %v and %v\n",
		// 	SumInts(ints),
		// 	SumFloats(floats))

		// fmt.Printf("Generic Sums: %v and %v\n",
		// 	SumIntsOrFloats[string, int64](ints),
		// 	SumIntsOrFloats[string, float64](floats))

		// fmt.Printf("Generic Sums, type parameters inferred: %v and %v\n",
		// 	SumIntsOrFloats(ints),
		// 	SumIntsOrFloats(floats))

		// fmt.Printf("Generic Sums with Constraint: %v and %v\n",
		// 	SumNumbers(ints),
		// 	SumNumbers(floats))

		// // without passing type
		// fmt.Println(reverse([]int{1, 2, 3, 4, 5}))

		// // passing type
		// fmt.Println(reverse[int]([]int{1, 2, 3, 4, 5}))

		// s := []int{1, 2, 3, 4, 5}
		// ForEach(s, func(ele int, i int, s []int) {
		// 	fmt.Printf("ele at %d is %d\n", i, ele)
		// })
	*/

	var abc=
	var a, b, c worker
	a = "A"
	b = "B"
	c = "C"
	DoWork([]worker{a, b, c})
	DoWork([]worker{a})


	var ga, gb, gc worker
	ga = "Generic A"
	gb = "Generic B"
	gc = "Generic C"
	DoWorkGeneric([]worker{ga, gb, gc})
	DoWorkGeneric([]worker{ga})

}

// SumInts adds together the values of m.
func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

// SumFloats adds together the values of m.
func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

// SumIntsOrFloats sums the values of map m. It supports both floats and integers
// as map values.
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

// SumNumbers sums the values of map m. Its supports both integers
// and floats as map values.
func SumNumbers[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}
