
package main

import (
	"fmt"

)

//_____________________________________________________________________

func playWithArrays() {

	// Creating Array of 3 ints
	//		Index Starts From 0 Till len - 1 i.e. [0, len)
	//		
	var a[3]int

	fmt.Println("\nArray a Length: ", len(a) )
	for index, value := range a {
		fmt.Printf("At Index: %d Value: %d\n", index, value)
	}

	fmt.Println( a[0] )
	fmt.Println( a[1] )
	fmt.Println( a[ len( a ) - 1 ] )

	// Compilation Error
	// 		invalid argument: array index 3 out of bounds [0:3]
	// fmt.Println( a[ len( a ) ] )	
	
	fmt.Println("\nArray a Length: ", len(a) )
	for _, value := range a {
		fmt.Printf("Value: %d\n", value)
	}

	var q [5]int = [5]int { 10, 20, 30, 40, 50 }
	var r [5]int = [5]int { 10, 20, 30  }

	fmt.Println("\nArray q Length: ", len(q) )
	for index, value := range q {
		fmt.Printf("At Index: %d Value: %d\n", index, value)
	}

	fmt.Println("\nArray r Length: ", len(r) )
	for index, value := range r {
		fmt.Printf("At Index: %d Value: %d\n", index, value)
	}

	s := [...]int { 10, 20, 100, 111 }
	fmt.Println("\nArray s Length: ", len(s) )
	fmt.Printf("\nArray s Data Type: %T\n", s )
	for index, value := range s {
		fmt.Printf("At Index: %d Value: %d\n", index, value)
	}

	some := [3]int { 100, 200, 300 }
	fmt.Println("\nArray s Length: ", len(some) )
	fmt.Printf("\nArray s Data Type: %T\n", some )

	// NOTE: Array Size Is Part of Array Type
	// Following Are Three Different Arrays Allocated 
	//		In Three Different Memory Locations
	aa := [2]int{ 10, 20 }   // Type [2]int Array Of int Of Size 2
	// ... Means Array Size Computed From Number Of Member Elements
	bb := [...]int{ 10, 20 } // Type [2]int
	cc := [2]int{ 10, 30 }   // Type [2]int 

	fmt.Printf("Array s Data Type: %T\n", aa )
	fmt.Printf("Array s Data Type: %T\n", bb )
	fmt.Printf("Array s Data Type: %T\n", cc )

	// In Go
	// 		Using == Operator 
	//		Array Values Are Compared Of Arrays of Same Type

	// In Java
	// 		Using == Operator 
	// 		Array References Are Compared Of Arrays of Same Type
	fmt.Println( aa == bb, aa == cc, bb == cc )
	// true false false

	dd := [3]int{ 10, 30 }  // Type [3]int
	fmt.Printf("Array s Data Type: %T\n", dd )
	fmt.Println( dd )
	
	//Compiler Error
	// 	invalid operation: aa == dd (mismatched types [2]int and [3]int)
	// Two Values Are Equal Only When Types Are Same
	// fmt.Println( aa == dd )

	var someAgain1 [10]int
	fmt.Println("\nArray s Length: ", len(someAgain1) )
	fmt.Printf("\nArray s Data Type: %T\n", someAgain1 )
	for index, value := range someAgain1 {
		fmt.Printf("At Index: %d Value: %d\n", index, value)
	}

	someAgain2 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("\nArray s Length: ", len(someAgain2) )
	fmt.Printf("\nArray s Data Type: %T\n", someAgain2 )
	for index, value := range someAgain2 {
		fmt.Printf("At Index: %d Value: %d\n", index, value)
	}

	someAgain3 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("\nArray s Length: ", len(someAgain3) )
	fmt.Printf("\nArray s Data Type: %T\n", someAgain3 )
	for index, value := range someAgain3 {
		fmt.Printf("At Index: %d Value: %d\n", index, value)
	}

	someAgain4 := [...]int{ 5: 50, 7: 70 }
	fmt.Println("\nArray s Length: ", len(someAgain4) )
	fmt.Printf("\nArray s Data Type: %T\n", someAgain4 )
	for index, value := range someAgain4 {
		fmt.Printf("At Index: %d Value: %d\n", index, value)
	}

	someAgain5 := [...]int{ 5: 50, 7: 70, 12: 222 }
	fmt.Println("\nArray s Length: ", len(someAgain5) )
	fmt.Printf("\nArray s Data Type: %T\n", someAgain5 )
	for index, value := range someAgain5 {
		fmt.Printf("At Index: %d Value: %d\n", index, value)
	}

	someAgain := [...]int{ 99 : -1 }
	fmt.Println("\nArray s Length: ", len(someAgain) )
	fmt.Printf("\nArray s Data Type: %T\n", someAgain )
	for index, value := range someAgain {
		fmt.Printf("At Index: %d Value: %d\n", index, value)
	}
}

//_____________________________________________________________________

func zeroValue( ival int ) {
	ival = 0
}

// Pointer Variable
//		Is A Variable Used To Address Of Memory Location

func zeroPointer( iptr *int ) {
	*iptr = 0
}

func playWithPointers() {
	var something = 111

	fmt.Println("Something Value: ", something)
	zeroValue( something )
	fmt.Println("Something Value: ", something)

	zeroPointer( &something )
	fmt.Println("Something Value: ", something)
}

// Function : playWithPointers
// Something Value:  111
// Something Value:  111
// Something Value:  0

//_____________________________________________________________________

type Flags uint

const ( // Definining Constants For Status BITs
	FlagUp Flags = 1 << iota  // Is Up  //  1 << 0
	FlagBroadcast 						//  1 << 1 
	FlagLoopback  						//  1 << 2
	FlagPointToPoint  					//  1 << 3
	FlagMulticast 						//  1 << 4
)

func IsUp(v Flags) bool 		{ return v & FlagUp == FlagUp }  // & Is Bitwise AND
func TurnDown(v * Flags) 		{ *v &^= FlagUp } 	// &^ AND NOT Operator
											  		// This Is A Bit Clear Operator
func SetBroadcast( v * Flags ) 	{ *v |= FlagBroadcast } // |  Bitwise OR Operator
func IsCast(v Flags) bool 		{ return v & (FlagBroadcast | FlagMulticast) != 0 }

func playWithFlags() {
	var v Flags = FlagMulticast | FlagUp
	fmt.Printf("%b %t\n", v, IsUp(v))

	TurnDown( &v )
	fmt.Printf("%b %t\n", v, IsUp(v))

	SetBroadcast( &v )
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

func doChange( a [10]int ) {
	for i, _ := range a {
		a[i] = 99
	}
	fmt.Println("\nInside doChange Function...")
}

func doChangeAgain( a *[10]int ) {
	for i, _ := range a {
		a[i] = 99
	}
	fmt.Println("\nInside doChangeAgain Function...")
}

func doChangeOnceAgain( slice []int ) {
	// Iterating Over Slice
	for i, _ := range slice {
		slice[i] = 99
	}
	fmt.Println("\nInside doChangeAgain Function...")
}

func printArray( a[10]int ) {
	for i, _ := range a {
		fmt.Printf("\nAt Index: %d Value: %d", i, a[i] )
	}
}

func playWithArray() {
	var numbers [10]int = [10]int{ 10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	fmt.Println("\nArray Before doChange Called...")
	printArray( numbers )
	doChange( numbers )  // By Default Passing Arrays As Pass By Value
	fmt.Println("\nArray After doChange Called...")
	printArray( numbers )

	var numbers1 [10]int = [10]int{ 10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	fmt.Println("\nArray Before doChangeAgain Called...")
	printArray( numbers1 )
	doChangeAgain( &numbers1 ) // Passing Arrays As Pass By Reference
	fmt.Println("\nArray After doChangeAgain Called...")
	printArray( numbers1 )

	var numbers2 [10]int = [10]int{ 10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	fmt.Println("\nArray Before doChangeOnceAgain Called...")
	printArray( numbers2 )
	doChangeOnceAgain( numbers2[ : ] ) // Passing Arrays Slice
	// Pass By Reference
	fmt.Println("\nArray After doChangeOnceAgain Called...")
	printArray( numbers2 )
}

//_____________________________________________________________________

func playWithArraySlices() {
	months := [...]string {
		0 : "", "Jan", "Feb", "Mar", "Apr", "May", "Jun", 
			    "Jul", "Aug", "Sep", "Oct", "Nov", "Dec", 
	}

	fmt.Println( "Length of Months: ", len( months ) )
	fmt.Println( "Months : ", months )
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

	quater1 := months[ 1 : 4 ]  
	quater2 := months[ 4 : 7 ]
	summer  := months[ 2 : 7 ]

	fmt.Println( "Quater1 : ", quater1 )
	fmt.Println( "Quater2 : ", quater2 )
	fmt.Println( "Summer : ", summer )	

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
	
	fmt.Println( "Months Array After Assignment In Slice: ", months )
	for index, value := range months {
		fmt.Println(index, value)
	}

	// Creating Slice For Underlining Array months
	//		Following Slice Will Access Full Array From Index 0 To len( months )
	//				StartIndex Is Omitted Than It's 0
	//				EndIndex Is Omitted Than It's len( Array )

	monthsSlice  := months[ : ]
	//  __________________________________
	// 	| Pointers To Start Location : 0 |
	//  __________________________________
	// 	| Length Of Slice = 13           |  
	//  __________________________________
	// 	| Capacity Of Slice = 13         |
	//  __________________________________
	fmt.Println( "Months Full Slice: ", monthsSlice )

	firstSixMonths := months[ : 7 ]
	lastSixMonths := months[ 7 : ]

	fmt.Println( "Months First Six Months: ", firstSixMonths )
	fmt.Println( "Months Last  Six Months: ", lastSixMonths )

	fmt.Println( "Months First Six Months Length: ", len( firstSixMonths ) )
	fmt.Println( "Months Last  Six Months Length: ", len( lastSixMonths )  )

	fmt.Println( "Months First Six Months Length: ", cap( firstSixMonths ) )
	fmt.Println( "Months Last  Six Months Length: ", cap( lastSixMonths )  )
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

func reverse( slice []int ) {
	for i, j := 0, len( slice ) - 1 ; i < j ; i, j = i + 1, j - 1 {
		slice[i], slice[j] = slice[j], slice[i]
	} 
}

func playWitArrayAndSlices() {
	// numbers Is A Array With Size Calcuated By Compiler
	//		Based On Number Of Members In Initialisation
	numbersArray := [...]int{ 10, 20, 30, 40, 50, 60, 70, 80, 90 }

	fmt.Println("\nArray :", numbersArray )
	reverse( numbersArray[ : ] )
	fmt.Println("\nArray :", numbersArray )

	// numbersSlice Is A Slice With Size Calcuated By Compiler
	//		Based On Number Of Members In Initialisation
	//		And One Background Array Will Be Allocated With These Members
	//		numbersSlice Will Be Pointing To Background Array
	numbersSlice := []int{ 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("\nSlice :", numbersSlice )
	fmt.Println("\nSlice Length and Capacity:", len( numbersSlice ), cap( numbersSlice ))
	fmt.Println("\nSlice :", numbersSlice[ : 5 ] )
	fmt.Println("\nSlice :", numbersSlice[ 5 :  ] )

	reverse( numbersSlice[ : 5 ] )
	reverse( numbersSlice[ 5 : ] )
	fmt.Println("\nSlice :", numbersSlice )
}

//_____________________________________________________________________


//_____________________________________________________________________
//_____________________________________________________________________
//_____________________________________________________________________
//_____________________________________________________________________
//_____________________________________________________________________
// EXPERIMENT FOLLOWING CODE, MOMENT DONE, PLEASE RAISE YOUR HAND!!!

func main() {
	fmt.Println("\nFunction : playWithArrays")
	playWithArrays()

	fmt.Println("\nFunction : playWithPointers")
	playWithPointers()

	fmt.Println("\nFunction : playWithFlags")
	playWithFlags()

	fmt.Println("\nFunction : playWithArray")
	playWithArray()

	fmt.Println("\nFunction : playWithArraySlices")
	playWithArraySlices()

	fmt.Println("\nFunction : playWitArrayAndSlices")
	playWitArrayAndSlices()

	// fmt.Println("\nFunction : ")
	// fmt.Println("\nFunction : ")
	// fmt.Println("\nFunction : ")
	// fmt.Println("\nFunction : ")
	// fmt.Println("\nFunction : ")
}
