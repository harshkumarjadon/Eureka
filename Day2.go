package main

import "fmt"

var pow = []int{1, 2, 3, 11, 4, 5, 6, 7, 8}
var cc string

func main() {
	// for i, r := range pow {
	// 	fmt.Println(i, " > ", r*r)
	// }

	for _, value := range pow {
		fmt.Printf("%d ", value)
	}
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
		//cc = 1 << uint(i)
		fmt.Print(cc, " ", i, " : ", pow[i], " - ", 1<<uint(i), " >\n ")
	}

	for i := range pow {
		//cc = 1 << uint(i)
		fmt.Print(cc, " ", i, " : ", pow[i], " - ", 1<<uint(i), " >\n ")
	}
	// fmt.Println("new program")

	// for i := 0; i < len(pow); i++ {
	// 	element := &pow[i]
	// 	element.Val = pow[i]
	// }
}
