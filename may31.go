package main

import (
	"fmt"
	"math"
)

//_____________________________________________________________________

// DESIGN PRINCIPLE
//		 DESIGN TOWARDS INTERFACES RATHER THAN CONCERTE CLASSES

// Intefaces Defines What It Is?
type Geometry interface {
	area() float64
	perimeter() float64
	volume() float64
}

// In Java Interface
// inteface Geometry {
// 	double area()
// 	double perimeter()
// }

// Interfaces Confirms Implicitly By Compilers
//
//	Compiler Deduces Rectangle Is Geometry Type
//		Because Rectangle Implements area() And perimeter()
type Rectangle struct {
	width, height float64
}

// In Java Classes Must Explicitly Implement Interfaces
//		All The Members Methods Must Be Implemented
// class Rectangle implements Geometry {
// 	double area() 	   {  	}
// 	double perimeter() { 	}
// }

// Interfaces Confirms Implicitly By Compilers
//
//	Compiler Deduces Rectangle Is Geometry Type
//		Because Rectangle Implements area() And perimeter()
type Circle struct {
	radius float64
}

// Implementation In Concrete
//
//	Why, Where, Which Way, When, How?
//
// Member Methods Of Rectangle
func (r Rectangle) area() float64 {
	return r.width * r.height
}

func (r Rectangle) perimeter() float64 {
	return 2 * (r.width + r.height)
}

func (r Rectangle) origin() (float64, float64) {
	return 0.0, 0.0
}

// Member Methods Of Circle
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (c Circle) origin() (float64, float64) {
	return 0.0, 0.0
}

// Polymorphic Function
//
//	General Functions
func measureGeometry(g Geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimeter())
}

// func measureGeometry( c Circle ) {
// 	fmt.Println( c )
// 	fmt.Println( c.area() )
// 	fmt.Println( c.perimeter() )
// }

// func measureGeometry( r Rectangle ) {
// 	fmt.Println( r )
// 	fmt.Println( r.area() )
// 	fmt.Println( r.perimeter() )
// }

func playWithGeometry() {
	r := Rectangle{width: 10, height: 20}
	c := Circle{radius: 10}

	fmt.Printf("\nRectanlge Area: %f Perimeter: %f", r.area(), r.perimeter())
	// fmt.Printf("\nCircle    Area: %f Perimeter: %f", c.area(), c.perimeter() )
	fmt.Printf("\nCircle    Area: %f Perimeter: %f", c.area()) //, c.perimeter() )

	fmt.Println("\nMagic Happens Here Onwards...")

	measureGeometry(r)
	// Compiler Error:
	// cannot use c (variable of type Circle) as type Geometry
	// in argument to measureGeometry:
	// 		Circle does not implement Geometry (missing perimeter method)

	measureGeometry(c)
}

//_____________________________________________________________________
//_____________________________________________________________________
//_____________________________________________________________________
//_____________________________________________________________________
//_____________________________________________________________________
// EXPERIMENT FOLLOWING CODE, MOMENT DONE, PLEASE RAISE YOUR HAND!!!

func main() {
	fmt.Println("\nFunction : playWithGeometry")
	playWithGeometry()

	// fmt.Println("\nFunction : ")
	// fmt.Println("\nFunction : ")
	// fmt.Println("\nFunction : ")
}
