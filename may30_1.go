package main

import (
	"bytes"
	"fmt"
	"math"
)

//_____________________________________________________________________

type PointData struct { // Associative Type
	X, Y float64
}

// Go In Relationship With Java
// class PointData {
// 		double X;
// 		double Y;
// }

// Distance Function : Calculates Distance Between Two PointDatas
//
//	Function Taking Two Arguments Of PointData Type
//	Return Type Is float64
func Distance(p, q PointData) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// Distance Method: Function Binded With Type ( Receiver Type )
//		Member Function Of PointData Type
//		Methdo Taking One Argument Of PointData Type
//		Return Type Is float64

//	 p Is Similar To this In C++/Java
//		In C++/Java
//			this Is Implicitly Passed To Every Member Method
//		In Go/Python
//			this Is Explicitly Passed To Every Member Method
//			In Go this Is Called Receiver Object Of Receiver Type
func (p PointData) Distance(q PointData) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// Compilation Error : PointData.Distance redeclared in this block
// Distance Between Three PointDatas PR = PQ + QR
// func (p PointData) Distance( q PointData, r PointData ) float64 {
// 	return math.Hypot( p.X - q.X, p.Y - q.Y ) + math.Hypot( q.X - r.X, q.Y - r.Y )
// }

type Path []PointData

// Path Type Member Distance Method
//
//	Distance Method Is Binded With Type Path
//	Receiver Type Is Path And
//	Receiver Object Is path Of Type Path
func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum = sum + path[i-1].Distance(path[i])
		}
	}
	return sum
}

// Go In Relationship With Java
// class Path {
// 	PointData [] path;
// 	public double Distance() {
// 		// this Is Implicitly Passed
// 		this.path
// 	}
// }

// Go In Relationship With Java
// func (this Path) Distance() float64 {
// 	sum := 0.0
// 	for i := range this {
// 		if i > 0 {
// 			sum = sum + this[ i-1 ].Distance( this[i] )
// 		}
// 	}
// 	return sum
// }

// Go In Relationship With Java
// public double Path::Distance() {
// 	sum := 0.0
// 	for i := range this {
// 		if i > 0 {
// 			sum = sum + this[ i-1 ].Distance( this[i] )
// 		}
// 	}
// 	return sum
// }

// func (modi Path) Distance() float64 {
// 	sum := 0.0
// 	for i := range modi {
// 		if i > 0 {
// 			sum = sum + modi[ i-1 ].Distance( modi[i] )
// 		}
// 	}
// 	return sum
// }

func playWithFunctionsAndMethods() {
	var PointData1 = PointData{10.0, 20.0}
	var PointData2 = PointData{30.0, 40.0}
	var PointData3 = PointData{50.0, 60.0}

	fmt.Println("Distance Between PointDatas: ", PointData1, PointData2)
	fmt.Println("\tDistance: ", Distance(PointData1, PointData2))

	fmt.Println("Distance Between PointDatas: ", PointData1, PointData2)
	//							PointData1 Is Reciever Object Of Receiver Type ( PointData )
	fmt.Println("\tDistance: ", PointData1.Distance(PointData2))

	fmt.Println("Distance Between PointDatas: ", PointData1, PointData2, PointData3)
	// too many arguments in call to PointData1.Distance
	// fmt.Println("\tDistance: ", PointData1.Distance( PointData2, PointData3 ) )
	var path Path = Path{
		PointData{10.0, 20.0}, // PointData A
		PointData{30.0, 40.0}, // PointData B
		PointData{50.0, 60.0}, // PointData C
	}
	//							Invoking Path Type Distance Method
	// Distance AC = AB + BC
	fmt.Println("\tDistance: ", path.Distance())

	var pathAgain Path = Path{
		PointData{10.0, 20.0},  // PointData A
		PointData{30.0, 40.0},  // PointData B
		PointData{50.0, 60.0},  // PointData C
		PointData{60.0, 100.0}, // PointData D
	}
	//							Invoking Path Type Distance Method
	// Distance AD = AB + BC + CD
	fmt.Println("\tDistance: ", pathAgain.Distance())
}

//_____________________________________________________________________

type Double float64

func (this Double) printValue() {
	fmt.Println("Printing Value: ", this)
}

func playWithMemberMethodOfDoubleType() {
	var something Double = 90.90
	// var somethingAgain float64 = 90.90

	something.printValue()
	// somethingAgain.printValue()
}

//_____________________________________________________________________
// EXPERIMENT FOLLOWING CODE, MOMENT DONE, PLEASE RAISE YOUR HAND!!!

type Vertex struct {
	X, Y float64
}

// Function
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Function
func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// Member Methods Of Type Vertex
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Member Methods Of Type Vertex
// Receiver Object v Pass By Reference
func (v *Vertex) ScaleBy(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// Vertex Value:  {10 20}
// Vertex Scaled Value:  {100 200}

// Receiver Object v Pass By Value
func (v Vertex) ScaleByAgain(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// Vertex Value:  {10 20}
// Vertex Scaled Value:  {10 20}

func playWithVertex() {
	v := Vertex{3, 4}
	fmt.Println("Vertex Value: ", v)
	Scale(&v, 10)
	fmt.Println("Vertex Scaled Value: ", v)
	fmt.Println("Vertex Absolute Value: ", Abs(v))

	vAgain1 := Vertex{10, 20}
	fmt.Println("Vertex Value: ", vAgain1)
	vAgain1.ScaleByAgain(10)
	fmt.Println("Vertex Scaled Value: ", vAgain1)
	fmt.Println("Vertex Absolute Value: ", vAgain1.Abs())

	vAgain2 := Vertex{10, 20}
	PointDataer := &vAgain2
	fmt.Println("Vertex Value: ", vAgain2)
	PointDataer.ScaleBy(10)
	fmt.Println("Vertex Scaled Value: ", vAgain2)
	fmt.Println("Vertex Absolute Value: ", vAgain2.Abs())

	vAgain21 := Vertex{10, 20}
	PointDataerAgain := &vAgain21
	fmt.Println("Vertex Value: ", vAgain21)
	// Compiler Will Converts Following Line Of Code To ( *PointDataerAgain ).ScaleByAgain( 10 )
	PointDataerAgain.ScaleByAgain(10)
	fmt.Println("Vertex Scaled Value: ", vAgain21)
	fmt.Println("Vertex Absolute Value: ", vAgain21.Abs())

	vAgain3 := Vertex{10, 20}
	fmt.Println("Vertex Value: ", vAgain3)
	// even though vAgain3 is a value and not a PointDataer,
	// The method with the PointDataer receiver is called automatically.
	// Compiler Will Converts Following Line Of Code To ( &vAgain3 ).ScaleBy( 10 )
	vAgain3.ScaleBy(10)
	fmt.Println("Vertex Scaled Value: ", vAgain3)
	fmt.Println("Vertex Absolute Value: ", vAgain3.Abs())
}

// Choosing a value or PointDataer receiver
// There are two reasons to use a PointDataer receiver.

// The first is so that the method can modify the value
//		that its receiver PointDatas to.

// The second is to avoid copying the value on each method call.
//		This can be more efficient if the receiver is a large struct, for example.
// 		In this example, both Scale and Abs are methods with receiver type *Vertex,
// 			even though the Abs method needn't modify its receiver.

// In general, all methods on a given type should have either value or
// PointDataer receivers, but not a mixture of both.

//_____________________________________________________________________

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

// Function
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// Function
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

// Member Method On Type Celsius
func (c Celsius) String() string {
	return fmt.Sprintf("%g ℃ ", c)
}

// Member Method On Type Fahrenheit
func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g ℉", f)
}

func playWithDegreesAndCentigrades() {
	freezingF := CToF(FreezingC)
	boilingF := CToF(BoilingC)

	fmt.Println("Freezing and Boiling In Fahrenheit: ",
		freezingF.String(), boilingF.String())

	freezingC := FToC(freezingF)
	boilingC := FToC(boilingF)

	fmt.Println("Freezing and Boiling In Centigrade: ",
		freezingC.String(), boilingC.String())

	// var farenheitValue Fahrenheit = 90.90
	// var centigradeValue Celsius = 100.100

	// Compiler Error:
	// invalid operation: farenheitValue + centigradeValue
	// (mismatched types Fahrenheit and Celsius)
	//var totalValue = farenheitValue + centigradeValue
	//	fmt.Println(totalValue)
}

//_____________________________________________________________________

// HOME WORK
// HOME WORK
// HOME WORK

// An IntSet is a set of small non-negative integers.
//
//	Its zero value represents the empty set.
type IntSet struct {
	words []uint64
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func playWithIntSet() {
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String()) // "{1 9 144}"

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String()) // "{9 42}"

	x.UnionWith(&y)
	fmt.Println(x.String()) // "{1 9 42 144}"

	fmt.Println(x.Has(9), x.Has(123)) // "true false"

	// Output:
	// {1 9 144}
	// {9 42}
	// {1 9 42 144}
	// true false
}

func playWithIntSetAgain() {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(42)

	fmt.Println(&x)         // "{1 9 42 144}"
	fmt.Println(x.String()) // "{1 9 42 144}"
	fmt.Println(x)          // "{[4398046511618 0 65536]}"

	// Output:
	// {1 9 42 144}
	// {1 9 42 144}
	// {[4398046511618 0 65536]}
}

//_____________________________________________________________________
// EXPERIMENT FOLLOWING CODE, MOMENT DONE, PLEASE RAISE YOUR HAND!!!

func main() {
	// fmt.Println("\nFunction : playWithFunctionsAndMethods")
	// playWithFunctionsAndMethods()

	// fmt.Println("\nFunction : playWithMemberMethodOfDoubleType")
	// playWithMemberMethodOfDoubleType()

	// fmt.Println("\nFunction : playWithVertex")
	// playWithVertex()

	// fmt.Println("\nFunction : playWithDegreesAndCentigrades")
	// playWithDegreesAndCentigrades()

	fmt.Println("\nFunction : playWithIntSet")
	playWithIntSet()

	fmt.Println("\nFunction : playWithIntSetAgain")
	playWithIntSetAgain()

	// fmt.Println("\nFunction : ")
	// fmt.Println("\nFunction : ")
	// fmt.Println("\nFunction : ")
	// fmt.Println("\nFunction : ")
	// fmt.Println("\nFunction : ")
}
