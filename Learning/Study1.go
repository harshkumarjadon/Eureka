package main

import (
	"fmt"
)

// Type address, auto increment onc ethe address and value are mapped
type address struct {
	hostname string
	port     int
}

func copyData() {
	a := []string{"a", "b", "c", "d", "e"}
	// var a1 []string
	// a1 = append(a, "1", "2")
	var dest = make([]string, len(a))
	fmt.Println(copy(dest, a))
	fmt.Println(a)
	fmt.Println(dest)
}

func underscoreValues() {
	// calling the function
	// function returns two values which are assigned to mul and blank identifier
	mul, _ := mul_div(105, 7)

	// only using the mul variable
	fmt.Println("105 x 7 = ", mul)
}

// function returning two values of integer type
func mul_div(n1 int, n2 int) (int, int) {
	// returning the values
	return n1 * n2, n1 / n2
}

// understanding pointers
func pointers() {
	var a int = 10
	var b *int = &a
	var c *int = b
	fmt.Println(&a)
	fmt.Println(a, "\n")
	fmt.Println(b)
	fmt.Println(*b)
	fmt.Println(&b, "\n")
	fmt.Println(c)
	fmt.Println(&c)
	fmt.Println(*c)
}

func helloPointer() {
	fmt.Println("Hello World")
}

type HelloFunc func(string)

func SayHello(to string) {
	fmt.Printf("Hello, %s!\n", to)
}

func removeArray() {
	s := []int{5, 6, 7, 8, 9}
	fmt.Println("Array items", s)
	fmt.Println("Preserve array - removing 1 ", remove(s, 2))       // "[5 6 8 9]"
	fmt.Println("Don't Preserve array - removing 1 ", remove(s, 2)) // "[5 6 8 9]"
}
func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}
func removeDonotpreserve(slice []int, i int) []int {
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}

func main() {
	//copyData()
	//pointers()
	//underscoreValues()
	//Function as pointer
	removeArray()
	// pFunc := helloPointer
	// pFunc()

	// //Use function signature
	// var hf HelloFunc
	// hf = SayHello
	// hf("world, as function signature")
	// //count increases automatically once you add the data in maps
	// hits := make(map[address]int)
	// hits[address{"golang.org", 443}]++
	// hits[address{"golang.org", 443}]++
	// fmt.Println(hits)

}
