
package main

//	Package
import "fmt"
import "rsc.io/quote"

func Hello() string {
	return quote.Hello()
}

func main() {
	fmt.Println("Hello World!!!!")
	Hello()
}


