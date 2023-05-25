package main

import (
	"fmt"
)

//_____________________________________________________________________

func playWithArrays() {

	// Creating Array of 3 ints
	//		Index Starts From 0 Till len - 1 i.e. [0, len)
	//
	var a [3]int

	fmt.Println("\nArray a Length: ", len(a))
	for index, value := range a {
		fmt.Printf("At Index: %d Value: %d\n", index, value)
	}

	fmt.Println(a[0])
	fmt.Println(a[1])
	fmt.Println(a[len(a)-1])

	// Compilation Error
	// 		invalid argument: array index 3 out of bounds [0:3]
	// fmt.Println( a[ len( a ) ] )

	fmt.Println("\nArray a Length: ", len(a))
	for _, value := range a {
		fmt.Printf("Value: %d\n", value)
	}

	var q [5]int = [5]int{10, 20, 30, 40, 50}
	var r [5]int = [5]int{10, 20, 30}

	fmt.Println("\nArray q Length: ", len(q))
	for index, value := range q {
		fmt.Printf("At Index: %d Value: %d\n", index, value)
	}

	fmt.Println("\nArray r Length: ", len(r))
	for index, value := range r {
		fmt.Printf("At Index: %d Value: %d\n", index, value)
	}

	s := [...]int{10, 20, 100, 111}
	fmt.Println("\nArray s Length: ", len(s))
	fmt.Printf("\nArray s Data Type: %T\n", s)
	for index, value := range s {
		fmt.Printf("At Index: %d Value: %d\n", index, value)
	}

	some := [3]int{100, 200, 300}
	fmt.Println("\nArray s Length: ", len(some))
	fmt.Printf("\nArray s Data Type: %T\n", some)

	someAgain := [...]int{9: 900, 11: 90, 111}
	fmt.Println("\nArray s Length: ", len(someAgain))
	fmt.Printf("\nArray s Data Type: %T\n", someAgain)
	for index, value := range someAgain {
		fmt.Printf("At Index: %d Value: %d\n", index, value)
	}

	aa := [2]int{100, 200}
	bb := [...]int{10, 20}
	cc := [2]int{10, 20}

	fmt.Println("aa > bb > cc", aa, bb, cc)
	fmt.Println(aa == bb, bb == cc, cc == aa)
	dd := [2]int{10, 300}
	fmt.Println("dd >", dd)
	fmt.Println(aa == dd)
}

//_____________________________________________________________________
//_____________________________________________________________________
//_____________________________________________________________________
//_____________________________________________________________________
//_____________________________________________________________________
//_____________________________________________________________________
//_____________________________________________________________________
//_____________________________________________________________________
// EXPERIMENT FOLLOWING CODE, MOMENT DONE, PLEASE RAISE YOUR HAND!!!

func main() {
	fmt.Println("\nFunction : playWithArrays")
	playWithArrays()

	// fmt.Println("\nFunction : ")
	// fmt.Println("\nFunction : ")
	// fmt.Println("\nFunction : ")
	// fmt.Println("\nFunction : ")
	// fmt.Println("\nFunction : ")
	// fmt.Println("\nFunction : ")
	// fmt.Println("\nFunction : ")
	// fmt.Println("\nFunction : ")
	// fmt.Println("\nFunction : ")
	// fmt.Println("\nFunction : ")
	// fmt.Println("\nFunction : ")
	// fmt.Println("\nFunction : ")
	// fmt.Println("\nFunction : ")
	// fmt.Println("\nFunction : ")
	// fmt.Println("\nFunction : ")
	// fmt.Println("\nFunction : ")
	// fmt.Println("\nFunction : ")
	// fmt.Println("\nFunction : ")
	// fmt.Println("\nFunction : ")
}
