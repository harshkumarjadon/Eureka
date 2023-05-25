package main

import (
	"fmt"

)

/*
In this example we have a Cat type, which can count its number of legs with its Legs method.
We embed this Cat type into a new type, an OctoCat, and declare that Octocats have five legs. However, although OctoCat defines its own Legs method, which returns 5, when the PrintLegs method is invoked, it returns 4.
This is because PrintLegs is defined on the Cat type. It takes a Cat as its receiver, and so
it dispatches to Cat‘s Legs method. Cat has no knowledge of the type it has been embedded
into, so its method set cannot be altered by embedding.

Thus, we can say that Go’s types, while being open for extension, are closed for modification.

In truth, methods in Go are little more than syntactic sugar around a function with a
predeclared formal parameter, their receiver.
func (c Cat) PrintLegs() {
        fmt.Printf("I have %d legs\n", c.Legs())
}

func PrintLegs(c Cat) {
        fmt.Printf("I have %d legs\n", c.Legs())
}
*/
type Cat struct {
	Name string
}

func (c Cat) Legs() int { return 4 }

func (c Cat) PrintLegs() {
	fmt.Printf("I have %d legs\n", c.Legs())
}

type OctoCat struct {
	Cat
}

func (o OctoCat) Legs() int { return 5 }

func main() {
	var octo OctoCat
	fmt.Println(octo.Legs()) // 5
	octo.PrintLegs()         // I have 4 legs
}
