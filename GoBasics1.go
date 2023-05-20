
package main

// import "fmt"
// import "math/rand"

import (
	"fmt"
	"math/rand"
	"os"
)

//_______________________________________________

func playWithValues() {
	fmt.Println("go" + "lang")
	fmt.Println("1 + 1 =", 1 + 1 )

	// 7.0 Type Is float64 i.e. double in C
	fmt.Println("7.0 + 3.0 =", 7.0 + 3.0 )

	fmt.Println( true && false )
	fmt.Println( true || false )
}

//_______________________________________________


func playWithVariables() {
	// Explicitly Type Annotation Specified
	var aa string 
	aa = "Good Morning"

	// Compiler Will Infer Type Implicitly
	// 1. Infer The Type of RHS Value
	// 2. Bind The Inferred Type From RHS to LHS
	// i.e. 
	// 1. "Initial Value" Inferred Type Will Be string
	// 2. string Type Binded With a L H S
	var a = "Initial Value"
	fmt.Println(a)
	fmt.Println(aa)

	// ./GoBasics.go:41:15: syntax error: unexpected newline, expecting type
	// var something

	// b and c Will Be Defined Of int Type
	//		1 Will Be Initialised To b
	//      2 Will Be Initialised To c
	var b, c int = 1, 2
	fmt.Println(b, c)

	// Inferred and Binded Type In bool
	var flag = true 
	fmt.Println(flag)

	// Inferred Type From RHS and Binded Type To LHS Is string
	var something = "Apple"
	fmt.Println( something )

	// Inferred Type From RHS and Binded Type To LHS Is string
	something1 := "Apple"
	fmt.Println( something1 )
}


//_______________________________________________

// Format Specifiers
// Reference:
// 		https://pkg.go.dev/fmt
func playWithPrintf() {
	fmt.Printf("My weight on the surface of Mars is %v lbs,", 149.0*0.3783)
	fmt.Printf(" and I would be %v years old.\n", 41*365/687)

	fmt.Printf("My weight on the surface of %v is %v lbs.\n", "Earth", 149.0)

	fmt.Printf("%-15v $%4v\n", "SpaceX", 94)
	fmt.Printf("%-15v $%4v\n", "Virgin Galactic", 100)
}

//_______________________________________________

func playWithVariablesAgain() {
	var distance = 5600
	var speed = 100

	// Defining Multiple Variables
	var (
		distance1 = 5600
		speed1 	  = 100
	)

	var distance2, speed2 = 5600, 1000

	fmt.Println(distance, speed)
	fmt.Println(distance1, speed1)
	fmt.Println(distance2, speed2)

	var randomNumber = rand.Intn(10) + 1
	fmt.Println( randomNumber )
}

//_______________________________________________

func playWithControlStructures() {
	var command = "go east"

	// if-else Construct In Go Lang
	if command == "go east" {
		fmt.Println("Going East...")
	} else if command == "go west" {
		fmt.Println("Going West...")
	} else {
		fmt.Println("Somewhere Else Going...")
	}

	var year = 2100
	var leapYear = year % 400 == 0 || ( year % 4 == 0 && year % 100 != 0 )

	if leapYear {
		fmt.Println("It's Leap Year: ", leapYear)
	} else {
		fmt.Println("It's Not Leap Year :", year)
	}
}

//_______________________________________________

var something1 = "Apple"

// Used Only Inside Local Scope
// something2 := "Apple"

func playWithScopes() {
	var something = "Some Value"
	// GoBasics.go:134:6: something redeclared in this block 
	//		Within A Same Scope You Can't Redeclare
	// var something = "Something Once Again"
	fmt.Println("something:", something)

	if something == "Some Value" {
		// Declared Inside Local Scope
		//		i.e. If Block Scope
		var something = "Ding Dong" 
		fmt.Println("something:", something)
	} else {
		fmt.Println("something Else:", something)
	}

	fmt.Println("something:", something)

	if something == "Some Value" {
		// Declared Inside Local Scope
		//		i.e. If Block Scope
		//		Identifiers With Same Name
		//		Declared Inside Local Scope Takes Precedence Over Global Scope
		something := "Ding Dong"
		fmt.Println("something:", something)
	} else {
		fmt.Println("something Else:", something)
	}
	fmt.Println("something:", something)

	var s, sep = "", ""
	for index, argument := range os.Args[1:] {
		// index And argument Are Local Variables 
		// i.e. Accessible Inside Local Scope Only
		s += sep + argument
		sep = " "
		fmt.Println("Index, Value: ", index, argument)
		// index And argument Will Not Exists After The Scope Ends i.e. Closing }
	}
	// fmt.Println("Index, Value: ", index, argument)
}

// DESIGN PRINCIPLE
//		DEFINE AND USE INSIDE LOCAL SOOPE ONLY
//		LOCAL STATE MUST BE ENCAPSULATED/LOCALISED

//_______________________________________________

func playWithScopesAgain() {
	
	// count Is Local To for Loop
	for count := 5 ; count > 0 ; count -- {
		fmt.Println( count )
	}

	// num Is Local To if-else 
	if num := rand.Intn( 3 ) ; num == 0 {
		fmt.Println("Random Number : ", num ) 
	} else if num == 1 {
		fmt.Println("Random Number Is One: ", num ) 
	} else {
		fmt.Println("Else Part...")
	}

	// num Is Local To switch
	switch  num := rand.Intn( 3 ) ; num {
	case 0:
		fmt.Println("Random Number Is Zero: ", num ) 
	case 1:
		fmt.Println("Random Number Is One: ", num ) 
	case 2:
		fmt.Println("Random Number Is Two: ", num ) 	
	default:
		fmt.Println("Default Case...", num)	
	}
}


//_______________________________________________
//_______________________________________________
//_______________________________________________
//_______________________________________________
//_______________________________________________

func main() {
	fmt.Println("\nFunction : playWithValues")
	playWithValues()

	fmt.Println("\nFunction : playWithVariables")
	playWithVariables()

	fmt.Println("\nFunction : playWithPrintf")
	playWithPrintf()

	fmt.Println("\nFunction : playWithVariablesAgain")
	playWithVariablesAgain()

	fmt.Println("\nFunction : playWithControlStructures")
	playWithControlStructures()

	fmt.Println("\nFunction : playWithScopes")
	playWithScopes()

	fmt.Println("\nFunction : playWithScopesAgain")
	playWithScopesAgain()

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


