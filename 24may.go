package main

import (
	"bufio"
	"fmt"
	"os"
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

	// NOTE: Array Size Is Part of Array Type
	// Following Are Three Different Arrays Allocated
	//		In Three Different Memory Locations
	aa := [2]int{10, 20} // Type [2]int Array Of int Of Size 2
	// ... Means Array Size Computed From Number Of Member Elements
	bb := [...]int{10, 20} // Type [2]int
	cc := [2]int{10, 30}   // Type [2]int

	fmt.Printf("Array s Data Type: %T\n", aa)
	fmt.Printf("Array s Data Type: %T\n", bb)
	fmt.Printf("Array s Data Type: %T\n", cc)

	// In Go
	// 		Using == Operator
	//		Array Values Are Compared Of Arrays of Same Type

	// In Java
	// 		Using == Operator
	// 		Array References Are Compared Of Arrays of Same Type
	fmt.Println(aa == bb, aa == cc, bb == cc)
	// true false false

	dd := [3]int{10, 30} // Type [3]int
	fmt.Printf("Array s Data Type: %T\n", dd)
	fmt.Println(dd)

	//Compiler Error
	// 	invalid operation: aa == dd (mismatched types [2]int and [3]int)
	// Two Values Are Equal Only When Types Are Same
	// fmt.Println( aa == dd )

	var someAgain1 [10]int
	fmt.Println("\nArray s Length: ", len(someAgain1))
	fmt.Printf("\nArray s Data Type: %T\n", someAgain1)
	for index, value := range someAgain1 {
		fmt.Printf("At Index: %d Value: %d\n", index, value)
	}

	someAgain2 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("\nArray s Length: ", len(someAgain2))
	fmt.Printf("\nArray s Data Type: %T\n", someAgain2)
	for index, value := range someAgain2 {
		fmt.Printf("At Index: %d Value: %d\n", index, value)
	}

	someAgain3 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("\nArray s Length: ", len(someAgain3))
	fmt.Printf("\nArray s Data Type: %T\n", someAgain3)
	for index, value := range someAgain3 {
		fmt.Printf("At Index: %d Value: %d\n", index, value)
	}

	someAgain4 := [...]int{5: 50, 7: 70}
	fmt.Println("\nArray s Length: ", len(someAgain4))
	fmt.Printf("\nArray s Data Type: %T\n", someAgain4)
	for index, value := range someAgain4 {
		fmt.Printf("At Index: %d Value: %d\n", index, value)
	}

	someAgain5 := [...]int{5: 50, 7: 70, 12: 222}
	fmt.Println("\nArray s Length: ", len(someAgain5))
	fmt.Printf("\nArray s Data Type: %T\n", someAgain5)
	for index, value := range someAgain5 {
		fmt.Printf("At Index: %d Value: %d\n", index, value)
	}

	someAgain := [...]int{99: -1}
	fmt.Println("\nArray s Length: ", len(someAgain))
	fmt.Printf("\nArray s Data Type: %T\n", someAgain)
	for index, value := range someAgain {
		fmt.Printf("At Index: %d Value: %d\n", index, value)
	}
}

//_____________________________________________________________________

func zeroValue(ival int) {
	ival = 0
}

// Pointer Variable
//		Is A Variable Used To Address Of Memory Location

func zeroPointer(iptr *int) {
	*iptr = 0
}

func playWithPointers() {
	var something = 111

	fmt.Println("Something Value: ", something)
	zeroValue(something)
	fmt.Println("Something Value: ", something)

	zeroPointer(&something)
	fmt.Println("Something Value: ", something)
}

// Function : playWithPointers
// Something Value:  111
// Something Value:  111
// Something Value:  0

//_____________________________________________________________________

type Flags uint

const ( // Definining Constants For Status BITs
	FlagUp           Flags = 1 << iota // Is Up  //  1 << 0
	FlagBroadcast                      //  1 << 1
	FlagLoopback                       //  1 << 2
	FlagPointToPoint                   //  1 << 3
	FlagMulticast                      //  1 << 4
)

func IsUp(v Flags) bool { return v&FlagUp == FlagUp } // & Is Bitwise AND
func TurnDown(v *Flags) { *v &^= FlagUp }             // &^ AND NOT Operator
// This Is A Bit Clear Operator
func SetBroadcast(v *Flags) { *v |= FlagBroadcast } // |  Bitwise OR Operator
func IsCast(v Flags) bool   { return v&(FlagBroadcast|FlagMulticast) != 0 }

func playWithFlags() {
	var v Flags = FlagMulticast | FlagUp
	fmt.Printf("%b %t\n", v, IsUp(v))

	TurnDown(&v)
	fmt.Printf("%b %t\n", v, IsUp(v))

	SetBroadcast(&v)
	fmt.Printf("%b %t\n", v, IsUp(v))
	fmt.Printf("%b %t\n", v, IsCast(v))
}

// Function : playWithFlags
// 10001 true
// 10000 false
// 10010 false
// 10010 true

//_____________________________________________________________________

// In C/C++/Java
//		By Default Arrays Are Pass By Reference

// In Go Lang
//		By Default Arrays Are Pass By Value

func doChange(a [10]int) {
	for i, _ := range a {
		a[i] = 99
	}
	fmt.Println("\nInside doChange Function...")
}

func doChangeAgain(a *[10]int) {
	for i, _ := range a {
		a[i] = 99
	}
	fmt.Println("\nInside doChangeAgain Function...")
}

func doChangeOnceAgain(slice []int) {
	// Iterating Over Slice
	for i, _ := range slice {
		slice[i] = 99
	}
	fmt.Println("\nInside doChangeAgain Function...")
}

func printArray(a [10]int) {
	for i, _ := range a {
		fmt.Printf("\nAt Index: %d Value: %d", i, a[i])
	}
}

func playWithArray() {
	var numbers [10]int = [10]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	fmt.Println("\nArray Before doChange Called...")
	printArray(numbers)
	doChange(numbers) // By Default Passing Arrays As Pass By Value
	fmt.Println("\nArray After doChange Called...")
	printArray(numbers)

	var numbers1 [10]int = [10]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	fmt.Println("\nArray Before doChangeAgain Called...")
	printArray(numbers1)
	doChangeAgain(&numbers1) // Passing Arrays As Pass By Reference
	fmt.Println("\nArray After doChangeAgain Called...")
	printArray(numbers1)

	var numbers2 [10]int = [10]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	fmt.Println("\nArray Before doChangeOnceAgain Called...")
	printArray(numbers2)
	doChangeOnceAgain(numbers2[:]) // Passing Arrays Slice
	// Pass By Reference
	fmt.Println("\nArray After doChangeOnceAgain Called...")
	printArray(numbers2)
}

//_____________________________________________________________________

func playWithArraySlices() {
	months := [...]string{
		0: "", "Jan", "Feb", "Mar", "Apr", "May", "Jun",
		"Jul", "Aug", "Sep", "Oct", "Nov", "Dec",
	}

	fmt.Println("Length of Months: ", len(months))
	fmt.Println("Months : ", months)
	for index, value := range months {
		fmt.Println(index, value)
	}
	// Creating Three Slices From Underlining One Array months
	//		Slices Are Data Structure Storing Following Information
	//  _____________________________
	// 	| Pointers To Start Location |
	// 	______________________________
	// 	| Length Of Slice            |
	// 	______________________________
	// 	| Capacity Of Slice          |
	// 	______________________________

	// All Following Slices Can Used To Used To Access Same Array
	// 		Slice Creation Syntax
	//			ArrayName[ StartIndex : EndIndex ]
	//			StartIndex and EndIndex Both Optional
	//				StartIndex Is Omitted Than It's 0
	//				EndIndex Is Omitted Than It's len( Array )

	quater1 := months[1:4]
	quater2 := months[4:7]
	summer := months[2:7]

	fmt.Println("Quater1 : ", quater1)
	fmt.Println("Quater2 : ", quater2)
	fmt.Println("Summer : ", summer)

	for _, s := range summer {
		for _, q := range quater2 {
			if s == q {
				fmt.Printf("	%s Appears In Both\n", s)
			}
		}
	}

	quater1[0] = "January"
	quater1[2] = "March"
	quater2[0] = "April"
	summer[4] = "JUNE"

	fmt.Println("Months Array After Assignment In Slice: ", months)
	for index, value := range months {
		fmt.Println(index, value)
	}

	// Creating Slice For Underlining Array months
	//		Following Slice Will Access Full Array From Index 0 To len( months )
	//				StartIndex Is Omitted Than It's 0
	//				EndIndex Is Omitted Than It's len( Array )

	monthsSlice := months[:]
	//  __________________________________
	// 	| Pointers To Start Location : 0 |
	//  __________________________________
	// 	| Length Of Slice = 13           |
	//  __________________________________
	// 	| Capacity Of Slice = 13         |
	//  __________________________________
	fmt.Println("Months Full Slice: ", monthsSlice)

	firstSixMonths := months[:7]
	lastSixMonths := months[7:]

	fmt.Println("Months First Six Months: ", firstSixMonths)
	fmt.Println("Months Last  Six Months: ", lastSixMonths)

	fmt.Println("Months First Six Months Length: ", len(firstSixMonths))
	fmt.Println("Months Last  Six Months Length: ", len(lastSixMonths))

	fmt.Println("Months First Six Months Length: ", cap(firstSixMonths))
	fmt.Println("Months Last  Six Months Length: ", cap(lastSixMonths))
}

//_____________________________________________________________________________________
//
// 								GO SLICES CONCEPTS
//_____________________________________________________________________________________

// Slices represent variable-length sequences whose elements all have the same type.
// 	A slice type is written []T, where the elements have type T;
// 		it looks like an array type without a size.

// A slice is a lightweight data structure that gives access to a subsequence
// (or perhaps all) of the elements of an array, which is known as the
// slice’s underlying array.

// 	A slice has three components: a pointer, a length, and a capacity.
// 		The pointer points to the first element of the array that is reachable
//		through the slice,
// 			which is not necessarily the array’s first element.
// 		The length is the number of slice elements; it can’t exceed the capacity,
// 			which is usually the number of elements between
// 			the start of the slice and the end of the underlying array.
// 		The built-in functions len and cap return those values.

// The slice operator s[ i : j ], where 0 ≤ i ≤ j ≤ cap(s),
// 	Creates a new slice that refers to elements i through j-1 of the sequence s,
// 	which may be an array variable, a pointer to an array, or another slice.
// 	The resulting slice has j-i elements.
// 	If i is omitted,it’s 0,and if j isomitted, it’s len(s).

// 	Thus the slice months[1:13] refers to the whole range of valid months,
// 	as does the slice months[1:]; the slice months[:] refers to the whole array.

//_____________________________________________________________________

func reverse(slice []int) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

func slicesEqual(x []int, y []int) bool {
	if len(x) != len(y) {
		return false
	}

	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

func playWitArrayAndSlices() {
	// numbers Is A Array With Size Calcuated By Compiler
	//		Based On Number Of Members In Initialisation
	numbersArray := [...]int{10, 20, 30, 40, 50, 60, 70, 80, 90}

	fmt.Println("\nArray :", numbersArray)
	reverse(numbersArray[:])
	fmt.Println("\nArray :", numbersArray)

	// numbersSlice Is A Slice With Size Calcuated By Compiler
	//		Based On Number Of Members In Initialisation
	//		And One Background Array Will Be Allocated With These Members
	//		numbersSlice Will Be Pointing To Background Array
	numbersSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("\nSlice :", numbersSlice)
	fmt.Println("\nSlice Length and Capacity:", len(numbersSlice), cap(numbersSlice))
	fmt.Println("\nSlice :", numbersSlice[:5])
	fmt.Println("\nSlice :", numbersSlice[5:])

	reverse(numbersSlice[:5])
	reverse(numbersSlice[5:])
	fmt.Println("\nSlice :", numbersSlice)

	aa := [2]int{10, 20} // Type [2]int Array Of int Of Size 2
	// ... Means Array Size Computed From Number Of Member Elements
	bb := [...]int{10, 20} // Type [2]int
	cc := [2]int{10, 30}   // Type [2]int

	fmt.Printf("Array s Data Type: %T\n", aa)
	fmt.Printf("Array s Data Type: %T\n", bb)
	fmt.Printf("Array s Data Type: %T\n", cc)

	// In Go
	// 		Using == Operator
	//		Array Values Are Compared Of Arrays of Same Type

	// In Java
	// 		Using == Operator
	// 		Array References Are Compared Of Arrays of Same Type
	fmt.Println(aa == bb, aa == cc, bb == cc)
	// true false false

	numbersArray1 := [...]int{10, 20, 30, 40, 50, 60, 70, 80, 90}
	numbersArray2 := [...]int{10, 20, 30, 40, 50, 60, 70, 80, 90}
	fmt.Println("Arrays Equality: ", numbersArray1 == numbersArray2)

	slice1 := numbersArray1[2:6]
	slice2 := numbersArray1[2:6]
	slice3 := numbersArray1[5:8]
	// Compiler Time Error
	// 		invalid operation: slice1 == slice2 (slice can only be compared to nil)

	// In GO Slices Are Reference Type
	//		References Can't Be Compared With == Operator
	// fmt.Println("Slices Equality: ", slice1 == slice2 )
	fmt.Println("Slices Equality: ", slicesEqual(slice1, slice2))
	fmt.Println("Slices Equality: ", slicesEqual(slice1, slice3))
}

//_____________________________________________________________________

func playWithSlicesFunctions() {

	// The built-in function make creates a slice of a specified element type,
	// length, and capacity. The capacity argument may be omitted, in which
	// case the capacity equals the length.
	// 		make([]T, len)
	// 		make([]T, len, cap) // same as make([]T, cap)[:len]

	s := make([]string, 3) // Background Array of Size 3 and Capacity 3 Allocated
	fmt.Println("Empty Slice : \n", s)
	fmt.Printf("Empty Slice Type : %T\n", s)
	fmt.Printf("Length: %d Capacity: %d", len(s), cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("\nSlice : ", s)
	fmt.Println("\nSlice Element s[2] : ", s[2])
	fmt.Printf("Length: %d Capacity: %d", len(s), cap(s))

	s = append(s, "d") // Background Array Of Size 6 Allocated
	fmt.Printf("Length: %d Capacity: %d", len(s), cap(s))

	fmt.Println("\nSlice : ", s)
	s = append(s, "e", "f")

	fmt.Println("\nSlice : ", s) // Background Array Of Size 6 Allocated
	fmt.Printf("\nLength: %d Capacity: %d", len(s), cap(s))

	s = append(s, "x")
	fmt.Println("\nSlice : ", s) // Background Array Of Size 12 Allocated
	fmt.Printf("\nLength: %d Capacity: %d", len(s), cap(s))

	c := make([]string, len(s))
	c = s // Reference Assignment i.e. It's Shallow Copy

	fmt.Println()
	fmt.Println("s Slice : ", s)
	fmt.Println("c Slice : ", c)
	// s Slice :  [a b c d e f x]
	// c Slice :  [a b c d e f x]

	s[0] = "Hello!"

	fmt.Println("s Slice : ", s)
	fmt.Println("c Slice : ", c)
	// s Slice :  [Hello! b c d e f x]
	// c Slice :  [Hello! b c d e f x]

	s[0] = "a"
	fmt.Println("s Slice : ", s)
	fmt.Println("c Slice : ", c)
	// s Slice :  [a b c d e f x]
	// c Slice :  [a b c d e f x]

	ss := make([]string, 0)
	fmt.Println(">>> ss Slice: ", ss)
	// fmt.Println(">>>>ss Slice Element ss[0]: ",  ss[0] )

	ss = append(ss, "AA", "BB", "CC", "DD", "EE", "FF")
	cc := make([]string, len(ss))

	fmt.Println("ss Slice: ", ss)
	fmt.Println("cc Slice: ", cc)

	copy(cc, ss) // Deep Copy
	fmt.Println("ss Slice: ", ss)
	fmt.Println("cc Slice: ", cc)

	ss[0] = "XXX"
	fmt.Println("ss Slice: ", ss)
	fmt.Println("cc Slice: ", cc)

	tt := []string{"Ding", "Dong", "Ming", "Mong"}
	fmt.Println("tt Slice: ", tt)
	fmt.Printf("tt Slice Type : %T", tt) // tt Slice Type : []string
}

//_____________________________________________________________________

func playWith2DimentionalArray() {
	// Creating 2D Fixed Size Array
	var a = [5][2]int{{0, 0}, {1, 2}, {2, 4}, {3, 6}, {4, 8}}
	var matrix = [3][2]int{{10, 20}, {30, 40}, {50, 60}}

	fmt.Println(a)
	fmt.Println(matrix)
}

func playWith2DimentionalDynamicSizeArray() {
	twoD := make([][]int, 3)

	for i := 0; i < 3; i++ {
		innerLen := i + 1

		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("2 D Dynamic Array : ", twoD)
}

//_____________________________________________________________________

func appendInt(x []int, y int) []int {
	var z []int

	zlen := len(x) + 1
	if zlen <= cap(x) {
		z = x[:zlen]
	} else { // Underlining Array Is Full
		// When Array Is Full, Allocate A New Array
		//	To Append Additional Data
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x) // Pool Size Increased
		}

		z = make([]int, zlen, zcap)
		copy(z, x)
	}

	z[len(x)] = y
	return z
}

func playWithAppendInt() {
	var x, y []int

	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf(" %d Capacity=%d \t %v \n", i, cap(y), y)
		x = y
	}
}

//_____________________________________________________________________

// Polymorphic Function
//		Based On Variable Number Of Arguments

func summation(numbers ...int) int {
	fmt.Print(numbers)
	total := 0

	for _, number := range numbers {
		total = total + number
	}

	return total
}

func playWithSummation() {
	var result int

	result = summation()
	fmt.Println("\n\t Result : ", result)

	result = summation(10, 20, 30, 40, 50)
	fmt.Println("\n\t Result : ", result)

	result = summation(10, 20, 30, 40, 50, 60, 70, 80, 90, 100)
	fmt.Println("\n\t Result : ", result)

	numbers := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	result = summation(numbers...)
	fmt.Println("\n\t Result : ", result)
}

//_____________________________________________________________________
// HOME WORK

func appendSlice(x []int, y ...int) []int {
	var z []int

	zlen := len(x) + len(y)
	if zlen <= cap(x) {
		z = x[:zlen]
	} else { // Underlining Array Is Full
		// When Array Is Full, Allocate A New Array
		//	To Append Additional Data
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x) // Pool Size Increased
		}

		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	copy(z[len(x):], y)
	return z
}

func playWithAppendSlices() {
	var x, y []int

	for i := 0; i < 10; i++ {
		y = appendSlice(x, i)
		fmt.Printf(" %d Capacity=%d \t %v \n", i, cap(y), y)
		x = y
	}
	fmt.Println(x)

	x = appendSlice(x, 10, 20, 30)
	fmt.Println(x)

	x = appendSlice(x, 11, 22, 33, 44, 55, 66)
	fmt.Println(x)

	numebrs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	x = appendSlice(x, numebrs...)
	fmt.Println(x)
}

//___________________________________________________________________
// HOME WORK

func filterEmpty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func filterEmptyAgain(strings []string) []string {
	out := strings[:0] // Zero Length Slice Of Original
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

func playWithNonEmpty() {
	data := []string{"one", "", "two", "ding", "", "ting"}
	fmt.Printf("%q \n", data)
	filteredData := filterEmpty(data)
	fmt.Printf("%q \n", filteredData)

	dataAgain := []string{"111", "", "", "9999", "", "Data", ""}
	fmt.Printf("%q \n", dataAgain)
	filteredDataAgain := filterEmptyAgain(dataAgain)
	fmt.Printf("%q \n", filteredDataAgain)
}

//_____________________________________________________________________
// MAPS in GO LANGUAGE
//		Set Of (Key, Value) Tuples
//

func playWithMaps() {
	// Creating Map With
	//		Key Of string Type and Value Of int Type
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("Map: ", m)

	v1 := m["k1"]
	fmt.Println("Map Key Value: ", v1)

	fmt.Println("Map Length :", len(m))
	delete(m, "k2")
	fmt.Println("Map: ", m)

	value, status := m["k1"]
	fmt.Println("Map Key Value And Status: ", value, status)

	value1, status1 := m["k3"]
	fmt.Println("Map Key Value And Status: ", value1, status1)

	mm := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("Map: ", mm)
}

//_____________________________________________________________________
// HOME WORK

func removeDuplicates() {
	seen := make(map[string]bool) // a set of strings
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		line := input.Text()
		if !seen[line] {
			seen[line] = true
			// fmt.Println(line)
		}
	}

	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "removeDuplicates: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(seen)
}

// Function : removeDuplicates
// Ding Dong
// Ding Dong
// Ding Dong
// Ting Tong
// Ting Tong
// Zing Zong
//  ^D

//_____________________________________________________________________
//
//							GO STRUCT
//_____________________________________________________________________

// Creating New Types Using type Keyword

type Circle struct {
	// Name Starting With Captials Are Exported Members
	X, Y, Radius int
}

type Wheel struct {
	X, Y, Radius, Spokes int
}

func playWithCircleAndWheel() {
	// c Is Instance Of Type Circle
	var c Circle
	c.X = 10
	c.Y = 20
	c.Radius = 50
	fmt.Println("\nCircle :", c) // Circle : {10 20 50}

	// w Is Instance Of Type Wheel
	var w Wheel
	w.X = 11
	w.Y = 22
	w.Radius = 55
	w.Spokes = 24
	fmt.Println("\nWheel :", w) // Wheel : {11 22 24 0}
}

type poi struct {
	X, Y int
}
type Cir struct {
	// Name Starting With Captials Are Exported Members
	point  poi
	Radius int
}

type Whe struct {
	circle Cir
	Spokes int
}

func playCir() {
	// c Is Instance Of Type Circle
	var x poi
	x.X = 10
	x.Y = 20

	var c Cir
	c.point.X = 10
	c.point.Y = 20
	c.Radius = 50
	fmt.Println("\nCircle :", c) // Circle : {10 20 50}

	// w Is Instance Of Type Wheel
	var w Whe
	w.circle.point.X = 11
	w.circle.point.Y = 22
	w.circle.Radius = 55
	w.Spokes = 24
	fmt.Println("\nWheel :", w) // Wheel : {11 22 24 0}
}

type Point struct {
	X int
	Y int
}

func ScalePoint(point Point, factor int) Point {
	var newX = point.X * factor
	var newY = point.Y * factor

	return Point{newX, newY}
	// return Point{ point.X * factor, point.Y * factor }
}

func AddPoint(point1 Point, point2 Point) Point {
	var newX = point1.X + point2.X
	var newY = point1.Y + point2.Y
	return Point{newX, newY}
	// return Point{ point1.X + point2.X , point1.Y + point2.Y }
}

func playWithPointType() {
	point1 := Point{10, 20}
	point2 := Point{100, 200}
	point3 := Point{10, 20}

	fmt.Println(" point1 : ", point1)
	fmt.Println(" point2 : ", point2)
	fmt.Println(" point3 : ", point3)

	fmt.Println(point1.X == point2.X && point1.Y == point2.Y)
	fmt.Println(point1.X == point3.X && point1.Y == point3.Y)

	fmt.Println(" point1 == point2 ", point1 == point2)
	fmt.Println(" point1 == point3 ", point1 == point3)

	fmt.Println("point1 * 5  : ", ScalePoint(point1, 5))
	fmt.Println("point2 * 10 : ", ScalePoint(point2, 10))
	fmt.Println("point3 * 2  : ", ScalePoint(point3, 2))

	point4 := AddPoint(point1, point2)
	fmt.Println(" point4 : ", point4)

	point5 := AddPoint(point1, point3)
	fmt.Println(" point5 : ", point5)
}

//_____________________________________________________________________

// Refactored Code : Created New Type Point
type Point2 struct {
	X, Y int
}

type Circle2 struct {
	// Center Point2
	Point2 // Type Or Structure Embedding
	Radius int
}

type Wheel2 struct {
	// Circle Circle2
	Circle2 // Type Or Structure Embedding
	Spokes  int
}

func playWithCircleAndWheel2() {
	// c Is Instance Of Type Circle
	var c Circle2
	c.Point2.X = 10
	c.Point2.Y = 20
	c.Radius = 50
	fmt.Println("\nCircle :", c) // Circle : {{10 20} 50}

	// w Is Instance Of Type Wheel
	var w Wheel2
	w.Circle2.Point2.X = 11
	w.Circle2.Point2.Y = 22
	w.Circle2.Radius = 55
	w.Spokes = 24
	fmt.Println("\nWheel :", w) // Wheel : {{{11 22} 55} 24}
}

//_____________________________________________________________________
//_____________________________________________________________________
//_____________________________________________________________________
//_____________________________________________________________________
//_____________________________________________________________________
//_____________________________________________________________________
// EXPERIMENT FOLLOWING CODE, MOMENT DONE, PLEASE RAISE YOUR HAND!!!

func main() {
	// fmt.Println("\nFunction : playWithArrays")
	// playWithArrays()

	// fmt.Println("\nFunction : playWithPointers")
	// playWithPointers()

	// fmt.Println("\nFunction : playWithFlags")
	// playWithFlags()

	// fmt.Println("\nFunction : playWithArray")
	// playWithArray()

	// fmt.Println("\nFunction : playWithArraySlices")
	// playWithArraySlices()

	// fmt.Println("\nFunction : playWitArrayAndSlices")
	// playWitArrayAndSlices()

	// fmt.Println("\nFunction : playWithSlicesFunctions")
	// playWithSlicesFunctions()

	// fmt.Println("\nFunction : playWith2DimentionalArray")
	// playWith2DimentionalArray()

	// fmt.Println("\nFunction : playWith2DimentionalDynamicSizeArray")
	// playWith2DimentionalDynamicSizeArray()

	// fmt.Println("\nFunction : playWithAppendInt")
	// playWithAppendInt()

	// fmt.Println("\nFunction : playWithSummation")
	// playWithSummation()

	// fmt.Println("\nFunction : playWithMaps")
	// playWithMaps()

	// fmt.Println("\nFunction : removeDuplicates")
	// removeDuplicates()

	// fmt.Println("\nFunction : playWithCircleAndWheel")
	// playWithCircleAndWheel()

	// fmt.Println("\nFunction : playWithCircleAndWheel")
	// playCir()

	// fmt.Println("\nFunction : playWithPointType")
	// playWithPointType()

	// fmt.Println("\nFunction : playWithCircleAndWheel2")
	playWithCircleAndWheel2()
	// fmt.Println("\nFunction : ")
	// fmt.Println("\nFunction : ")
	// fmt.Println("\nFunction : ")
}
