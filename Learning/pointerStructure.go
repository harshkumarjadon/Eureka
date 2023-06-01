package main

import (
	"fmt"

)

type PeronData struct {
	Name string
	Age  int
}

func ValueReceiver(p PeronData) {
	p.Name = "John"
	fmt.Println("Inside ValueReceiver : Name - Age", p.Name, p.Age)
}
func PointerReceiver(p *PeronData) {
	p.Age = 24
	fmt.Println("Inside PointerReceiver model: Name - Age ", p.Name, p.Age)
}
func main() {
	p := PeronData{"Tom", 28}
	p1 := &PeronData{"Patric", 68}
	ValueReceiver(p)
	fmt.Println("Inside Main after value receiver : Name, Age ", p.Name, p.Age)
	PointerReceiver(p1)
	fmt.Println("Inside Main after value receiver : Name, Age ", p1.Name, p1.Age)
}
