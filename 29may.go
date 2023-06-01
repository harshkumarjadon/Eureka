// In C/C++/Java

package main

import (
	"fmt"

)

//_____________________________________________________________________
// EXPERIMENT FOLLOWING CODE, MOMENT DONE, PLEASE RAISE YOUR HAND!!!

// In C/C++/Java
//		Function Can Return ONLY ONE Value
// In Go Lang
//		Function Can Return More Than One Value

// Function Takes Two int Arguemnts
//
//	And Return Tuple Of Two int Values
func division(dividend, divisor int) (int, int) {
	if divisor != 0 {
		quotient := dividend / divisor
		reminder := dividend % divisor
		return quotient, reminder
	} else {
		fmt.Println("Error : Division By Zero!")
	}
	return 0, 0
}

// Labelled Tuple
func divisionAgain(dividend, divisor int) (quotient int, remainder int) {
	if divisor != 0 {
		quotient := dividend / divisor
		reminder := dividend % divisor
		// return
		return quotient, reminder
	} else {
		fmt.Println("Error : Division By Zero!")
	}
	return 0, 0
}

func playWithDivision() {
	q, r := division(10, 3)
	fmt.Printf("\n Quotient : %d Remainder : %d", q, r)

	qq, rr := divisionAgain(10, 3)
	fmt.Printf("\n Quotient : %d Remainder : %d", qq, rr)
}

//_____________________________________________________________________
// EXPERIMENT FOLLOWING CODE, MOMENT DONE, PLEASE RAISE YOUR HAND!!!

func sumFunction(x, y int) int { return x + y }
func subFunction(x, y int) int { return x - y }

// Closure/Lambda Expression
var sumClosureOutside = func(x, y int) int { return x + y }

func playWithClosures() {
	var x, y = 30, 10
	var result int

	result = sumFunction(x, y)
	fmt.Println("\tResult : ", result)

	result = subFunction(x, y)
	fmt.Println("\tResult : ", result)

	// Anonymous Functions : Functions Without Name
	// Closure/Lambda Expression
	sumClosure := func(x, y int) int { return x + y }
	subClosure := func(x, y int) int { return x - y }

	result = sumClosure(x, y)
	fmt.Println("\tResult : ", result)

	result = subClosure(x, y)
	fmt.Println("\tResult : ", result)

	result = sumClosureOutside(x, y)
	fmt.Println("\tResult : ", result)
}

//_____________________________________________________________________

func playWithClosuresCaptures() { // Enclosing/Outside Context
	var start = 0
	var result int

	//Closure/Lambda Expression
	incrementBy7 := func() int { // Enclosed/Inside Context
		// Enclosed/Inside Context Captures Enclosing/Outside Context
		//		Hence start Variable Is Accessible Inside incrementBy7 Closure
		start = start + 7
		return start
	}

	incrementBy10 := func() int { // Enclosed/Inside Context
		// Enclosed/Inside Context Captures Enclosing/Outside Context
		//		Hence start Variable Is Accessible Inside incrementBy10 Closure
		start = start + 10
		return start
	}

	result = incrementBy7()
	fmt.Println("incrementBy7 : ", result)

	result = incrementBy7()
	fmt.Println("incrementBy7 : ", result)

	result = incrementBy7()
	fmt.Println("incrementBy7 : ", result)

	result = incrementBy10()
	fmt.Println("incrementBy10 : ", result)

	result = incrementBy10()
	fmt.Println("incrementBy10 : ", result)
}

//_____________________________________________________________________

func outsideFunction() {
	somethingOutside := 100

	// In C/C++/Java/Go
	//		You CANNOT Define Function Inside Function
	// syntax error: unexpected insideFunction, expecting (
	// func insideFunction() {
	// 	somethingInside := 100
	// 	fmt.Println("\nInside Function : ", somethingInside )
	// }

	fmt.Println("\nOutiside Function : ", somethingOutside)
	// insideFunction()
}

func playWithOutsideFunction() {
	outsideFunction()
}

var SIZE = 10

func playWithArrayAndPointers() {
	// Array Name Is Pointer To Starting Element Of Array
	//		numbers Is Pointer To First Element Of Continuous Array
	var numbers = [10]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	//pointNumbers * []int

	// fmt.Println("\nPrinting Values Using Array Index\n")
	// for i := 0; i < SIZE; i++ {
	// 	fmt.Println("\t%d", numbers[i])
	// }

	//pointNumbers = numbers
	fmt.Println("\nPrinting Values Using Pointer To Array\n")
	for i := 0; i < SIZE; i++ {
		fmt.Println("\t%d", *(&numbers[i])) // + i))
		fmt.Println("\t%d", (&numbers[i])) // + i))
		fmt.Println("\t%d", numbers)        // + i))
	}
}

//_____________________________________________________________________
// EXPERIMENT FOLLOWING CODE, MOMENT DONE, PLEASE RAISE YOUR HAND!!!

//	Higher Order Functions
//		Functions Which Takes Functions As Arguments
//			And/Or
//		Can Returns Function

// addition, substraction and multiplication Have Function Type func (int, int) int
// Function Type
//
//	func (int, int) int
func addition(x, y int) int       { return x + y }
func substraction(x, y int) int   { return x - y }
func multiplication(x, y int) int { return x * y }

// Can I Say calculator Function Is Polymorphic?
// Polymoprhism
//		Using Mechanism By Passing Function/Behviour To Function/Behaviour

// Higher Order Functions
//
//	Functions Which Takes Functions As Arguments
//	Here calculator Is Higher Order Function
//		Because It Takes Function As Argument
func calculator(x, y int, operation func(int, int) int) int {
	return operation(x, y)
}

func playWithCalculator() {
	a := 40
	b := 10
	var result = 0

	result = calculator(a, b, addition)
	fmt.Println("\tResult : ", result)

	result = calculator(a, b, substraction)
	fmt.Println("\tResult : ", result)

	result = calculator(a, b, multiplication)
	fmt.Println("\tResult : ", result)

	// What Is The Type Of doSomething ???
	// Function Type: func( int, int, func(int, int) int ) int
	doSomething := calculator
	result = doSomething(a, b, multiplication)
	fmt.Println("\tResult : ", result)

	// var doSomething1 int
	// Compilationn Error : cannot use calculator
	//		(value of type func(x int, y int, operation func(int, int) int) int)
	//  	as type int in assignment
	// doSomething1 = calculator
	var doSomething1 func(int, int, func(int, int) int) int
	doSomething1 = calculator
	result = doSomething1(a, b, multiplication)
	fmt.Println("\tResult : ", result)
}

//_____________________________________________________________________
// EXPERIMENT FOLLOWING CODE, MOMENT DONE, PLEASE RAISE YOUR HAND!!!

//	Higher Order Functions
//	adderFunction Is Higher Order Function
//		Can Returns A Function Of Function Type
//		i.e. Can Return Function Of Type func(int) int

func adderFunction() func(int) int {
	sum := 0
	totalClosure := func(x int) int {
		sum += x
		return sum
	}

	return totalClosure
}

func playWithAdder() {
	// Compilation Error
	//		cannot use adderFunction() (value of type func(int) int) as
	//		type int in assignment
	// 	var doSomething int
	var doSomething func(int) int

	doSomething = adderFunction()
	for i := 0; i < 10; i++ {
		fmt.Println(doSomething(i))
	}

	// What Is The Data Type Of doSomethingAgain
	// Type of doSomethingAgain Is Function
	//			Function Which Takes int Argument and Return int
	// Function Type: func(int) int
	doSomethingAgain := adderFunction()
	for i := 0; i < 10; i++ {
		fmt.Println(doSomethingAgain(i))
	}

}

//_____________________________________________________________________
//_____________________________________________________________________
// EXPERIMENT FOLLOWING CODE, MOMENT DONE, PLEASE RAISE YOUR HAND!!!

func main() {
	// fmt.Println("\nFunction : playWithDivision")
	// playWithDivision()

	// fmt.Println("\nFunction : playWithClosures")
	// playWithClosures()

	// fmt.Println("\nFunction : playWithClosuresCaptures")
	// playWithClosuresCaptures()

	// fmt.Println("\nFunction : playWithOutsideFunction")
	// playWithOutsideFunction()

	fmt.Println("\nFunction : playWithArrayAndPointers")
	playWithArrayAndPointers()

	// fmt.Println("\nFunction : playWithCalculator")
	// playWithCalculator()

	// fmt.Println("\nFunction : playWithAdder()")
	// playWithAdder()
	// fmt.Println("\nFunction : ")
	// fmt.Println("\nFunction : ")
	// fmt.Println("\nFunction : ")
	// fmt.Println("\nFunction : ")
	// fmt.Println("\nFunction : ")
	// fmt.Println("\nFunction : ")
}
