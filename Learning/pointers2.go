package main

import (
	"fmt"
)

func pointers() {
	var a int = 20
	var ptr *int = &a
	fmt.Println("Address of a", &a)
	fmt.Println("value of a", a)
	fmt.Println("value of ptr", *ptr)
	fmt.Println("Address of a stored in ptr", ptr)
	fmt.Println("Address of ptr", &ptr)
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
	FlagUp Flags = 1 << iota // Is Up
	FlagBroadcast
	FlagLoopback
	FlagPointToPoint
	FlagMulticast
)

func IsUp(v Flags) bool { return v&FlagUp == FlagUp } // & Is Bitwise AND
func TurnDown(v *Flags) { *v &^= FlagUp }             // &^ AND NOT Operator
// This Is A Bit Clear Operator
func SetBroadcast(v *Flags) { *v |= FlagBroadcast } // |   Bitwise OR Operator
func IsCast(v Flags) bool   { return v&(FlagBroadcast|FlagMulticast) != 0 }

func playWithFlags() {
	fmt.Println("FlagUp %b", FlagUp)
	fmt.Println("FlagBroadcast %b", FlagBroadcast)
	fmt.Println("FlagLoopback %b", FlagLoopback)
	fmt.Println("FlagPointToPoint %b", FlagPointToPoint)
	fmt.Println("FlagMulticast %b", FlagMulticast)

	fmt.Printf("FlagUp %b\n", FlagUp)
	fmt.Printf("FlagBroadcast %b \n", FlagBroadcast)
	fmt.Printf("FlagLoopback %b \n", FlagLoopback)
	fmt.Printf("FlagPointToPoint %b \n", FlagPointToPoint)
	fmt.Printf("FlagMulticast %b \n", FlagMulticast)

	var v Flags = FlagMulticast | FlagUp
	fmt.Printf("%b %d\n", v, v)
	fmt.Printf("%b %t\n", v, IsUp(v))

	TurnDown(&v)
	fmt.Printf("%b %t\n", v, IsUp(v))

	SetBroadcast(&v)
	fmt.Printf("%b %t\n", v, IsUp(v))
	fmt.Printf("%b %t\n", v, IsCast(v))
}

func main() {
	pointers()

	fmt.Println("\nFunction : playWithPointers")
	playWithPointers()

	fmt.Println("\nFunction : playWithFlags")
	playWithFlags()
}
