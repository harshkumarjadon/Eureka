
package learnKotlin

fun outsideFunction() {
	var somethingOutside = 100

	// In C/C++/Java/Go
	//		You CANNOT Define Function Inside Function

	// In Kotlin/Swift etc.
	//		You CAN Define Function Inside Function

	// Local Function: Function Defined Inside Function
	fun insideFunction() {
		var somethingInside = 100
		println("\nInside Function : $somethingInside" )
	}

	println("\nOutiside Function : $somethingOutside" )
	insideFunction()

	// Lambda Expression
	var insideLambda = { 
		var somethingInsideLamdba = 100
		println("\nInside Lambda : $somethingInsideLamdba" )
	}

	insideLambda()
}

fun playWithOutsideFunction() {
	outsideFunction()
}

//_____________________________________________________________________

//File : Context 0
var somethingAtFile = 111

class DingDong { // For Context 1 Outside Context 0 
	var somethingDingDong = 100

	// Member Function
	fun doSomething() : Int { // For Context 2 Outside Content 1
		return somethingDingDong + somethingAtFile
	}
}

fun playWithClasses() {
	var dingdong = DingDong()

	var result = dingdong.doSomething()
	println("\nResult : $result ")
}

// Function : playWithClasses
// Result : 211 


//_____________________________________________________________________
// EXPERIMENT FOLLOWING CODE, MOMENT DONE, PLEASE RAISE YOUR HAND!!!

fun main() {
	// println("\nFunction : playWithOutsideFunction")
	// playWithOutsideFunction()

	println("\nFunction : playWithClasses")
	playWithClasses()
}
