
package main

import (
	"fmt"
	"math"
)

//_____________________________________________________________________

const (
	width, height 	= 600, 300 				// Canvas Size In Pixels
	cells 			= 100 					// Number Of Grid Cells
	xyrange 		= 30.0 					// Axis Ranges
	xyscale 		= width / 2 / xyrange   // Pixel Per x And y Unit
	zscale 			= height * 0.4 			// Pixel Per z Unit
	angle 			= math.Pi / 6 			// Anle of x And y Axes
)


var sin30, cos30 = math.Sin(angle), math.Cos(angle)

// Function Takes Two Arguments Of int Type
//		Returns Pair Of Values Of float64 Type
func corner( i int, j int ) ( float64, float64 ) {
	// Here x And y Will Be Of float64 Type
	//		1. Type Is Inferred From RHS Expression
	//		2. Inferred Type Is Binded With LHS
	x := xyrange * ( float64( i )  / cells - 0.5 )
	y := xyrange * ( float64( j )  / cells - 0.5 )

	// Here z Will Be Of float64 Type
	//		1. Type Is Inferred From RHS Expression
	//		2. Inferred Type Is Binded With LHS
	z := f(x, y)

	// Here sx And sy Will Be Of float64 Type
	//		1. Type Is Inferred From RHS Expression
	//		2. Inferred Type Is Binded With LHS
	sx := width/2 + ( x - y ) * cos30 * xyscale
	sy := height/2 + ( x + y) * sin30 * xyscale - z* zscale 

	return sx, sy // Returning Tuple Of Two Values
}

// Function Takes Two Arguments Of float64 Type
//		Returns A Value Of float64 Type
// func f(x float64, y float64) float64 {
func f(x, y float64) float64 {
	r := math.Hypot( x, y )
	return math.Sin( r ) / r
}

func generateCanvas() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
	"style='stroke: grey; fill: white; stroke-width: 0.7' "+
	"width='%d' height='%d'>", width, height)

	for  i := 0 ; i < cells ; i++   {
		for  j := 0 ; j < cells ; j++  {
			ax, ay := corner( i + 1, j )
			bx, by := corner( i, j )
			cx, cy := corner( i, j + 1 )
			dx, dy := corner( i + 1, j + 1)


			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}

	fmt.Println( "</svg>")
}

//_____________________________________________________________________
//_____________________________________________________________________
//_____________________________________________________________________
//_____________________________________________________________________
//_____________________________________________________________________

func main() {
	generateCanvas()

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


