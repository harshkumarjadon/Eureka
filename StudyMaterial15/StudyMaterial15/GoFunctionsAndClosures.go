
// In C/C++/Java

package main

import (
	"fmt"
	"math"
	"os"
	"net/http"
)

//_____________________________________________________________________
// EXPERIMENT FOLLOWING CODE, MOMENT DONE, PLEASE RAISE YOUR HAND!!!

// In C/C++/Java
//		Function Can Return ONLY ONE Value
// In Go Lang
//		Function Can Return More Than One Value

// Function Takes Two int Arguemnts
//		And Return Tuple Of Two int Values
func division( dividend, divisor int ) (int, int) {
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
func divisionAgain( dividend, divisor int ) ( quotient int, remainder int) {
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
	q, r := division( 10, 3 )
	fmt.Printf("\n Quotient : %d Remainder : %d", q, r )

	qq, rr := divisionAgain( 10, 3 )
	fmt.Printf("\n Quotient : %d Remainder : %d", qq, rr )
}

//_____________________________________________________________________
// EXPERIMENT FOLLOWING CODE, MOMENT DONE, PLEASE RAISE YOUR HAND!!!

func sumFunction(x, y int ) int { return x + y }
func subFunction(x, y int ) int { return x - y }
// Closure/Lambda Expression
var sumClosureOutside = func(x, y int) int { return x + y }

func playWithClosures() {
	var x, y = 30, 10
	var result int

	result = sumFunction(x, y)
	fmt.Println("\tResult : ", result )

	result = subFunction(x, y)
	fmt.Println("\tResult : ", result )

	// Anonymous Functions : Functions Without Name
	// Closure/Lambda Expression
	sumClosure := func(x, y int) int { return x + y }
	subClosure := func(x, y int) int { return x - y }

	result = sumClosure(x, y)
	fmt.Println("\tResult : ", result )

	result = subClosure(x, y)
	fmt.Println("\tResult : ", result )

	result = sumClosureOutside(x, y)
	fmt.Println("\tResult : ", result )
}

//_____________________________________________________________________
// EXPERIMENT FOLLOWING CODE, MOMENT DONE, PLEASE RAISE YOUR HAND!!!

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
// EXPERIMENT FOLLOWING CODE, MOMENT DONE, PLEASE RAISE YOUR HAND!!!

func outsideFunction() {
	somethingOutside := 100

	// In C/C++/Java/Go
	//		You CANNOT Define Function Inside Function
	// syntax error: unexpected insideFunction, expecting (
	// func insideFunction() {
	// 	somethingInside := 100
	// 	fmt.Println("\nInside Function : ", somethingInside )
	// }

	fmt.Println("\nOutiside Function : ", somethingOutside )
	// insideFunction()
}

func playWithOutsideFunction() {
	outsideFunction()
}

//_____________________________________________________________________
// EXPERIMENT FOLLOWING CODE, MOMENT DONE, PLEASE RAISE YOUR HAND!!!

//	Higher Order Functions
//		Functions Which Takes Functions As Arguments
//			And/Or
//		Can Returns Function

// addition, substraction and multiplication Have Function Type func (int, int) int
// Function Type
//		func (int, int) int
func addition( x, y int ) int 		{ return x + y }
func substraction( x, y int ) int 	{ return x - y }
func multiplication( x, y int ) int	{ return x * y }

// Can I Say calculator Function Is Polymorphic?
// Function Polymoprhism
//		Using Mechanism : By Passing Function To Function
//						  OR
//						  By Passing Behviour To Behaviour

//	Higher Order Functions
//		Functions Which Takes Functions As Arguments
//		Here calculator Is Higher Order Function
//			Because It Takes Function As Argument
func calculator( x, y int, operation func (int, int) int ) int {
	return operation(x, y)
}

func playWithCalculator() {
	a := 40
	b := 10
	var result = 0

	result = calculator(a, b, addition) 		// Configuring As addition
	fmt.Println("\tResult : ", result)

	result = calculator(a, b, substraction) 	// Configuring As substraction
	fmt.Println("\tResult : ", result)

	result = calculator(a, b, multiplication) 	// Configuring As multiplication
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
	var doSomething1 func( int, int, func(int, int) int ) int 
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
	// sum Local Variable
    sum := 0
    // totalClosure Local Variable
    //		Pointer To Function/Closure
    totalClosure := func(x int) int {
        sum += x
        return sum
    }
    return totalClosure

    // Go Compiler Will Escape/Move Local Variable sum To Heap
} // Come Out Of Enclosing Function

// Escape Analysis In Go
//      Local Variable Will Escape/Move To Heap 
//          Moment Programmer Returns Address Of Local Variable

// In Go Language : Go Compiler Will Generate Following Code
//      It Will Replace return &b; Line With Following Code
//      For All Local Variables Who's Refrences Are Returned
//      1. Allocate Memory At Heap
//      int * addressAtHeap = (int *) malloc( sizeof(int) );
//      2. Copy/Move/Escpae The Value Of sum To Heap
//      *addressAtHeap = sum;  
//      3. The Address Of Value At Heap i.e. addressAtHeap
//      	Used By totalClosure Lambda	As Follows
//	    	
//			totalClosure := func(x int) int {
//	        	*addressAtHeap += x
//	        	return *addressAtHeap
//	    	}

func playWithAdder() {
// Compilation Error
//		cannot use adderFunction() (value of type func(int) int) as 
//		type int in assignment
// 	var doSomething int
	var doSomething func(int) int
    doSomething = adderFunction()
    for i := 0; i < 10; i++ {
        fmt.Println( doSomething(i) )
    }
	// What Is The Data Type Of doSomethingAgain
	// Type of doSomethingAgain Is Function 
	//			Function Which Takes int Argument and Return int
	// Function Type: func(int) int
    doSomethingAgain := adderFunction()
    for i := 0; i < 10; i++ {
        fmt.Println( doSomethingAgain(i) )
    }
}

//_____________________________________________________________________
// EXPERIMENT FOLLOWING CODE, MOMENT DONE, PLEASE RAISE YOUR HAND!!!

func Sphere() func( radius float64 ) float64 {
	volumeClosure := func( radius float64 ) float64 {
		volume := 4 / 3 * math.Pi * radius * radius * radius
		return volume
	}

	return volumeClosure
}

func playWithFunctionsAndClosures() {
	// What Is Data Type Of sphereVolume Variable???
	// 		Function Type : func( float64 ) float64
	sphereVolume := Sphere()
	var sphereVolumeAgain func(float64) float64 = Sphere()

	fmt.Println("Volume Of Sphere: ", sphereVolume( 5 ) )
	fmt.Println("Volume Of Sphere: ", sphereVolumeAgain( 5 ) )
}

//_____________________________________________________________________
// HOME WORK

// Simulate Family Scenario Of Cooking Of Food By Cook
//		Based On Request and Receipies Given By Family Members

// type Food struct 	{ }
// type Recipe struct 	{ }
// type Cook struct 	{ } 
// type Request struct  { }
// type Desire struct   { }

// func cooking(cook Cook, recipe Recipe, request Request) Food {
// 	// Cook Knows These Recipes
// 	pasta 	:== func() string { fmt.Println("Cooking Pasta...") }
// 	biryani :== func() string { fmt.Println("Cooking Biryani...") }

// 	// askLadyOfHouse
// 	request, recipe = ladyOfHouse()
	
// 	// askManOfHouse
// 	request, recipe = manOfHouse()

// 	// askKidsOfHouse
// 	request, recipe = ladyOfHouse()

// 	if request == pasta {
// 		food := pasta()
// 		return food
// 	}

// 	if request == biryani {
// 		food := biryani()
// 		return food
// 	}

// 	if (request == somethingElse ) {
// 		food := recipe()
// 		return  food
// 	}
// }

// func ladyOfHouse( desire Desire ) ( Request, func() Recipe ) {
// 	// Recipe 01
// 	muttonBiryani := func() string { 
// 		// Step By Step Process For Cooking Mutton Biryani
// 		fmt.Println("Cooking Muttong Briyani...") 
// 	}
	
// 	// Recipe 01
// 	dalMakhni := func() string { 
// 		// Step By Step Process For Cooking Dal Makhini
// 		fmt.Println("Cooking Dal Makhani...") 
// 	}

// 	// Recipe 03
// 	shahiPaneer := func() string { 
// 		// Step By Step Process For Cooking Shahi Paneer
// 		fmt.Println("Cooking Sahai Paneer...") 
// 	}

// 	if desire == "Shahi Paneer" {
// 		return request, shahiPaneer
// 	} else desire == "Mutton Biryani" {
// 		return request, muttonBiryani
// 	} else {
// 		return request, dalMakhni
// 	}
// }

// func manOfHouse( desire Desire ) ( Request, func() Recipe) {
// 	chilliPaneer := func() string { fmt.Println("Cooking Chilli Paneer...") }

// 	return request, chilliPaneer
// }

//_____________________________________________________________________

func firstFunction() {
	fmt.Println("\tFirst Function Called...")
}

func secondFunction() {
	fmt.Println("\tSecond Function Called...")
}

func playWithFunctionsDefer() {
	// firstFunction()
	// secondFunction()
	// Function : playWithFunctions
	// First Function Called...
	// Second Function Called...

	// defer Keyword Delays Execution Of Statement
	//		Till Exit Of Enclosing Function
	defer firstFunction()
	secondFunction()
	// Function : playWithFunctions
	// Second Function Called...
	// First Function Called...

	fmt.Println("Exiting Enclosing Function...")
	defer firstFunction()
	defer secondFunction()
// Function : playWithFunctions
// Exiting Enclosing Function...
// 	Second Function Called...
// 	First Function Called...

//  defer Statement Will Be Called Just Befor Exit From This Function
}

func playWithFunctionDeferUseCase() {
	// Every File/Resource Opened/Used By Programmer Must Be Closeds
	response, err := http.Get("https://joeshaw.org")
	if err != nil {
	    fmt.Println("Error :", err )
	}
	defer response.Body.Close() // Good Programming Practice

	// Your Can Use Response response Variable 
	//		Till Closing Bracket Of Enclosing Function

	file, err := os.Open("/home/joeshaw/notes.txt")
	if err != nil {
	    fmt.Println("Error :", err )
	}

	// func (f *File) Close() error
	// Close closes the File, rendering it unusable for I/O. 
	// On files that support SetDeadline, any pending I/O operations will be 
	// canceled and return immediately with an ErrClosed error. 
	// Close will return an error if it has already been called.

	defer file.Close() // Ignorring The error Returned By Close() Function

    defer func() { // Writing defer Closure
        cerr := file.Close()
        if cerr != nil {
     		fmt.Println("Error :", cerr )
        }
    }() // Invoking The Closure

	//  defer Statement Will Be Called Just Befor Exit From This Function
}

// / The behavior of defer statements is straightforward and predictable. 
// There are three simple rules:
// 	1. A deferred function’s arguments are evaluated when the defer statement 
//		is evaluated.
//  2. Deferred function calls are executed in Last In First Out order 
// 		after the Surrounding/Enclosing function returns.
// 3. Deferred functions may read and assign to the returning function’s 
//		named return values.

//_____________________________________________________________________
//_____________________________________________________________________
//_____________________________________________________________________
//_____________________________________________________________________
//_____________________________________________________________________
//_____________________________________________________________________
//_____________________________________________________________________
// EXPERIMENT FOLLOWING CODE, MOMENT DONE, PLEASE RAISE YOUR HAND!!!

func main() {
	fmt.Println("\nFunction : playWithDivision")
	playWithDivision()

	fmt.Println("\nFunction : playWithClosures")
	playWithClosures()

	fmt.Println("\nFunction : playWithClosuresCaptures")
	playWithClosuresCaptures()

	fmt.Println("\nFunction : playWithOutsideFunction")
	playWithOutsideFunction()

	fmt.Println("\nFunction : playWithCalculator")
	playWithCalculator()

	fmt.Println("\nFunction : playWithFunctionsAndClosures")
	playWithFunctionsAndClosures()

	fmt.Println("\nFunction : playWithFunctionsDefer")
	playWithFunctionsDefer()

	fmt.Println("\nFunction : playWithFunctionDeferUseCase")
	playWithFunctionDeferUseCase()

	// fmt.Println("\nFunction : ")
	// fmt.Println("\nFunction : ")
	// fmt.Println("\nFunction : ")
	// fmt.Println("\nFunction : ")
	// fmt.Println("\nFunction : ")
	// fmt.Println("\nFunction : ")
	// fmt.Println("\nFunction : ")
	// fmt.Println("\nFunction : ")
}
