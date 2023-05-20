package main

import (
	"fmt"
	"math"
	"math/big"
	"strconv"

)

func types() {
	var a = 10
	var aa int = 100
	fmt.Printf("\nValues %v and type %T", a, a)
	fmt.Printf("\nValues %v and type %T", aa, aa)

	var f = 10.0
	var ff float32 = 100.99
	var fff float64 = 999.9999

	fmt.Printf("\nValues %v and type %T", f, f)
	fmt.Printf("\nValues %v and type %T", ff, ff)
	fmt.Printf("\nValues %v and type %T", fff, fff)

	var str, str1, str2 = "hello", "Good", "Morning"
	fmt.Printf("\nValues %v and type %T", str, str)
	fmt.Printf("\nValues %v and type %T", str1, str1)
	fmt.Printf("\nValues %v and type %T", str2, str2)

}


//_____________________________________________________________________

// In Go float32 And float64 Type
//		float32 Is Same As float in C/C++/Java
//		float64 Is Same As double in C/C++/Java

func playWithTypes() {
	// Inferred Type From RHS Is float64 i.e. double In C
	// 		float64 Binded With LHS 
	//		Hence pi Type Is float64	
	pi := 3.1415

	// Inferred Type From RHS Is float64 i.e. double In C
	// 		float64 Binded With LHS 
	//		Hence piValue Type Is float64	
	var piValue = 3.1415 
	var piAgain float64 = 3.1415 

	// All Above Declarations Are Equivalent
	fmt.Println(pi, piValue, piAgain)


	// To Make float32 Explicitly Annotate It!
	//		float32 Is Same As float in C/C++/Java
	var piAgain1 float32 = 3.1415 
	fmt.Println( piAgain1 )


	third := 1.0 / 3
	fmt.Println(third)
	fmt.Printf("%v\n", third)
	fmt.Printf("%f\n", third)
	fmt.Printf("%.3f\n", third)
	fmt.Printf("%4.2f\n", third)
}

//_____________________________________________________________________

func playWithFloatingPointTypes() {
	// Type Is float64
	piggyBank := 0.1
	piggyBank = piggyBank + 0.2

	fmt.Println( piggyBank == 0.3 )

	// Function : playWithFloatingPointTypes
	// false

	// BAD CODE : Buggy Logic
	//		Serious Security Threat :  Validation Can Be Bypassed
	if piggyBank == 0.3 {
		fmt.Println( "piggyBank Have Dollers : ", piggyBank )		
	} else {
		fmt.Println( "piggyBank Is Dead..." )		
	}

	// GOOD CODE 
	if math.Abs( piggyBank - 0.3 ) < 0.000001  {
		fmt.Println( "piggyBank Have Dollers : ", piggyBank )		
	} else {
		fmt.Println( "piggyBank Is Dead..." )		
	}
}

// DESIGN PRINCIPLE AND PRACTICE
//		Never Compare Floating Point Values With Equality/Not Equality Operator ==, !=
//		Correct Mathematics For Comparing Floating Points As Follows
// 			| floatingValue1 - floatingValue2 | < ToleranceValue

//_____________________________________________________________________

// Architecture-Independent Integer Types In Go Language

// Integer Types 	Range

// int8 			–128 to 127  							8-bit (one byte)
// uint8			0 to 255
// int16 			–32,768 to 32,767  						16-bit (two byte)
// uint16 			0 to 65535
// int32 			–2,147,483,648 to 2,147,483,647  		32-bit (four bytes)
// uint32 			0 to 4,294,967,295
// int64 			–9,223,372,036,854,775,808 to 9,223,372,036,854,775,807
// uint64  		0 to 18,446,744,073,709,551,615  		64-bit (eight bytes)

func playWithTypesAgain() {
	var a = 10 // Inferred Type of 10 Is int And Binded To a Variable Type int
	var aa int = 100

	fmt.Printf("\nValue %v and Type %T", a, a)
	fmt.Printf("\nValue %v and Type %T", aa, aa)

	var f = 90.90
	var ff float64 = 90.90
	var fff float32 = 90.90

	fmt.Printf("\nValue %v and Type %T", f, f)
	fmt.Printf("\nValue %v and Type %T", ff, ff)
	fmt.Printf("\nValue %v and Type %T", fff, fff)

	var greeting = "Good Morning!"
	fmt.Printf("\nValue %v and Type %T", greeting, greeting)

	var dollar = '$'  // Inferred Type Is int32 and Binded With dollar Variable
	fmt.Printf("\nValue %v and Type %T", dollar, dollar) 

	//cannot use 900 (untyped int constant) as int8 value in 
	//	variable declaration (overflows)
	// var b int8 = 900
	var b int8 = 127
	fmt.Printf("\nValue %v and Type %T", b, b)

	var red, green, blue  		= 0, 141, 213 		 // int Type Values In Base 10 
	var red1, green1, blue1  	= 0x00, 0x8d, 0xd5   // int Type Values In Base 16

	fmt.Printf("\nValue %v and Type %T", red, red)
	fmt.Printf("\nValue %v and Type %T", red1, red1)

	fmt.Printf("\nValue %v and Type %T", green, green)
	fmt.Printf("\nValue %v and Type %T", blue, blue)

	fmt.Printf("\nValue %v and Type %T", green1, green1)
	fmt.Printf("\nValue %v and Type %T", blue1, blue1)

	age := 1415
	marsAge := float64( age ) // Type Casting
	fmt.Println( marsAge )

	fmt.Println("string conversion", strconv.Itoa(age))

	fmt.Println( math.MinInt16 )
	fmt.Println( math.MaxInt16 )

	fmt.Println( math.MinInt32 )
	fmt.Println( math.MaxInt32 )
}

// On Unix-based operating systems, time is represented as the number 
// of seconds since January 1, 1970 UTC (Coordinated Universal Time). 
// In the year 2038, the number of seconds since January 1, 1970 will 
// exceed two billion, the capacity of an int32.

// Thankfully, int64 can support dates well beyond 2038. This is a 
// situation where int32 or int simply won’t do. Only the int64 and uint64 
// integer types are able to store numbers well beyond two billion on 
// all platforms.

//_____________________________________________________________________

func playWithLargeNumbers() {
	const lightSpeed = 299792 // km/s
	const secondsPerDay = 86400

	var distanceStart = 41.3e12  // float64
	fmt.Println("Alpha Centauri is", distanceStart, "km away.")

	var distance int64 = 41.3e12 // Exponential Notation  41300000000000
	fmt.Println("Alpha Centauri is", distance, "km away.")

	days := distance / lightSpeed / secondsPerDay
	fmt.Println("That is", days, "days of travel at light speed.")

	bigNumber1 := big.NewInt( 888880000000777 )
	bigNumber2 := big.NewInt( 111 )

	fmt.Println( bigNumber1 )
	fmt.Println( bigNumber2 )
}

//_____________________________________________________________________
//_____________________________________________________________________
//_____________________________________________________________________
//_____________________________________________________________________
//_____________________________________________________________________
//_____________________________________________________________________
//_____________________________________________________________________

func main() {
	fmt.Println("\nFunction : playWithTypes")
	playWithTypes()

	fmt.Println("\nFunction : playWithFloatingPointTypes")
	playWithFloatingPointTypes()

	fmt.Println("\nFunction : playWithTypesAgain")
	playWithTypesAgain()

	fmt.Println("\nFunction : playWithLargeNumbers")
	playWithLargeNumbers()
	
	type()

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
	// fmt.Println("\nFunction : ")
}


