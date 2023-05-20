package main

import (
	"fmt"
	"math/rand"
	"os"
)

func playValues() {
	fmt.Println("go" + "lang")
	fmt.Println("1 + 1", 1+1)
	fmt.Println("7.0 + 3.0", 7.0+3.0)
	fmt.Println(true && false)
	fmt.Println(true || false)
}

func variables() {
	var a = "initial values"
	var aa string
	aa = "Good morning"
	fmt.Println(a)
	fmt.Println(aa)
	//Throws error as the type is not assigned
	//var something

	var a1, b int = 1, 2
	fmt.Println("a1 ", a1, " b ", b)
	var flag = false
	fmt.Println("flag ", flag)

	var something = "hello"
	fmt.Println(something)
	something1 := "New world"
	fmt.Println(something1)

}

func displayPrintf() {

}

func variables2() {
	// var distance = 5000
	// var speed = 100

	// var (
	// 	distance1 = 5000
	// 	speed1    = 400
	// )

	// var distance2, speed2 = 5000, 1000
	var randomNumber = rand.Intn(10) + 1
	fmt.Println(randomNumber)
}

func control() {
	var command = "go east"

	if command == "go east" {
		fmt.Println("go east")
	} else if command == "go west" {
		fmt.Println("Go west")
	} else {
		fmt.Println("else part")
	}

	var year = 2100
	var leapYear = year%400 == 0 || (year%4 == 0 && year%100 != 0)

	if leapYear {
		fmt.Println("Leap year")

	} else {
		fmt.Println("Not leap year")
	}
}

func scopes() {
	var something = "hello"

	if something == "hello" {
		var something = "new SCope"
		fmt.Println("something", something)
	} else {
		fmt.Println("something else", something)
	}
	fmt.Println("something outside", something)

	var s, sep = "", ""
	for index, argument := range os.Args[1:] {
		s += sep + argument
		sep = " "
		fmt.Println("Index, Values ", index, argument)
	}
}

func intType() {
	pi := 3.14
	var piValue = 3.14
	var piAgain float64 = 3.14
	fmt.Println(pi, piValue, piAgain)
	var pifloat float32 = 3.14
	fmt.Println(pifloat)
}

func main() {
	//playValues()
	//variables()
	//variables2()
	//control()
	//scopes()
	intType()
}
