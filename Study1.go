package main

import (
	"fmt"
)

func copyData() {
	a := []string{"a", "b", "c", "d", "e"}
	// var a1 []string
	// a1 = append(a, "1", "2")
	var dest = make([]string, len(a))
	fmt.Println(copy(dest, a))
	fmt.Println(a)
	fmt.Println(dest)
}

func main() {
	copyData()
}
