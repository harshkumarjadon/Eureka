
package main

import (
	"fmt"
	"math"
	"image/color"
)

//_____________________________________________________________________

// GO FOLLOWS DUCK TYPING
//		DUCK TEST
//			If It Walks Like A Duck
//			AND
//			It Quacks Like A Duck
//			Then It Must Be Duck

// DESIGN PRINCIPLE
//		 DESIGN TOWARDS INTERFACES RATHER THAN CONCERTE CLASSES

// Intefaces Defines What It Is?
type Geometry interface {
	area() 			float64
	perimeter() 	float64
}

// In Java Interface
// inteface Geometry {
// 	double area()
// 	double perimeter()
// }

// Interfaces Confirms Implicitly By Compilers
//		Compiler Deduces Rectangle Is Geometry Type
//			Because Rectangle Implements area() And perimeter()
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
//		Compiler Deduces Rectangle Is Geometry Type
//			Because Rectangle Implements area() And perimeter()
type Circle struct {
	radius float64
}

// Implementation In Concrete
//		Why, Where, Which Way, When, How?
// Member Methods Of Rectangle
func ( r Rectangle ) area() float64 {
	return r.width * r.height
}

func ( r Rectangle ) perimeter() float64 {
	return 2 * ( r.width + r.height )
}

func ( r Rectangle ) origin() ( float64, float64 ) {
	return 0.0, 0.0
}

// Member Methods Of Circle
func ( c Circle ) area() float64 {
	return math.Pi * c.radius * c.radius
}

func ( c Circle ) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func ( c Circle ) origin() ( float64, float64 ) {
	return 0.0, 0.0
}

// Polymorphic Function
//		Using Mechanism: By Passing Intefrace Type As Argument
//		INTERFACE DRIVEN DESIGN
func measureGeometry( g Geometry ) {
	fmt.Println( g )
	fmt.Println( g.area() )
	fmt.Println( g.perimeter() )
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
	r := Rectangle{ width : 10, height : 20 }
	c := Circle{ radius : 10 }

	fmt.Printf("\nRectanlge Area: %f Perimeter: %f", r.area(), r.perimeter() )
	// fmt.Printf("\nCircle    Area: %f Perimeter: %f", c.area(), c.perimeter() )
	fmt.Printf("\nCircle    Area: %f Perimeter: %f", c.area() ) //, c.perimeter() )

	fmt.Println("\nMagic Happens Here Onwards...")

	measureGeometry( r )
	
	// Compiler Error: 
	// cannot use c (variable of type Circle) as type Geometry 
	// in argument to measureGeometry:
	// 		Circle does not implement Geometry (missing perimeter method)
	measureGeometry( c )
}

//_____________________________________________________________________
// IN GO TYPE EMBEDDING

type Point struct {
	X, Y float64
}

// ColoredPoint Type Composed Of Point Type
// Composing Type : ColoredPoint
type ColoredPoint struct {
	// Composed Type: Point
	Point 			// Embedding Point Type Inside ColoredPoint
	Color color.RGBA
}

// Member Method For Type Point
func ( this Point ) Distance( other Point ) float64 {
	fmt.Println("Point Distance Method Called...")
	return math.Hypot( this.X - other.X, this.Y - other.Y )
}

func ( this Point ) ScaleBy( scaleFactor float64 ) Point {
	this.X = this.X * scaleFactor
	this.Y = this.Y * scaleFactor

	return Point{ this.X, this.Y }
}

// Member Method For Type ColoredPoint
func ( this ColoredPoint ) Distance( other ColoredPoint ) float64 {
	fmt.Println("ColoredPoint Distance Method Called...")
	return math.Hypot( this.X - other.X, this.Y - other.Y )
}

func playWithColoredPoint() {
	var point1 = Point{ 55, 50 }
	var point2 = Point{ 66, 60 }

	fmt.Println("Distance Between Points: ", point1.Distance( point2 ) )
	fmt.Println("Actual Point1          : ", point1 )
	fmt.Println("Scaled Point1          : ", point1.ScaleBy( 10.0 ) )
	fmt.Println("Actual Point2          : ", point2 )
	fmt.Println("Scaled Point2          : ", point2.ScaleBy( 10.0 ) )

	red 	:= color.RGBA{ 255,   0, 0, 255 }
	blue 	:= color.RGBA{ 	 0, 255, 0, 255 }

	var coloredPoint1  = ColoredPoint{ Point{ 10, 20 }, red  }
	var coloredPoint2  = ColoredPoint{ Point{ 11, 22 }, blue }

	fmt.Println("ColoredPoint Point1: ", coloredPoint1)
	fmt.Println("ColoredPoint Point2: ", coloredPoint2)

	// Compile Time Error:
	//	cannot use coloredPoint2 (variable of type ColoredPoint) 
	// 	as type Point in argument to coloredPoint1.Distance
	// var distanceBetweenColoredPoint = coloredPoint1.Distance( coloredPoint2 )
	// var distanceBetweenColoredPoint = coloredPoint1.Distance( coloredPoint2.Point )
	var distanceBetweenColoredPoint = coloredPoint1.Distance( coloredPoint2 )

	fmt.Println("Distance Between Points: ", distanceBetweenColoredPoint )

	fmt.Println("Actual ColoredPoint1          : ", coloredPoint1 )
	fmt.Println("Scaled ColoredPoint1          : ", coloredPoint1.ScaleBy( 10.0 ) )
	fmt.Println("Actual ColoredPoint2          : ", coloredPoint2 )
	fmt.Println("Scaled ColoredPoint2          : ", coloredPoint2.ScaleBy( 10.0 ) )

}

// We can call methods of the embedded Point field using a 
// receiver of type ColoredPoint, even though ColoredPoint 
// has no declared methods:

//_____________________________________________________________________
//_____________________________________________________________________
//_____________________________________________________________________
//_____________________________________________________________________
//_____________________________________________________________________
//_____________________________________________________________________
// EXPERIMENT FOLLOWING CODE, MOMENT DONE, PLEASE RAISE YOUR HAND!!!

func main() {
	fmt.Println("\nFunction : playWithGeometry")
	playWithGeometry()

	fmt.Println("\nFunction : playWithColoredPoint")
	playWithColoredPoint()

	// fmt.Println("\nFunction : ")
	// fmt.Println("\nFunction : ")
	// fmt.Println("\nFunction : ")
	// fmt.Println("\nFunction : ")
	// fmt.Println("\nFunction : ")
	// fmt.Println("\nFunction : ")
	// fmt.Println("\nFunction : ")
	// fmt.Println("\nFunction : ")
}

