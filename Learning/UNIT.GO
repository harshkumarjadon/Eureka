package main

import "fmt"

// var FlagUp uint = 1
// var FlagBroadcast uint = 10
// var FlagLoopback uint = 100
// var FlagPointToPoint uint = 1000
// var FlagMulticast uint = 10000

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

func main() {
	var a uint = 1  /* 60 = 0011 1100 */
	var b uint = 10 /* 13 = 0000 1101 */
	var c uint = 100
	// var d uint = 1000
	// var e uint = 10000

	var v Flags = FlagMulticast | FlagUp
	fmt.Printf("%b %d\n", v, v)
	fmt.Printf("%b %t\n", v, IsUp(v))

	TurnDown(&v)
	fmt.Printf("%b %t\n", v, IsUp(v))

	SetBroadcast(&v)
	fmt.Printf("%b %t\n", v, IsUp(v))
	fmt.Printf("%b %t\n", v, IsCast(v))

	c = a & b /* 12 = 0000 1100 */
	fmt.Printf("Line 1 - Value of c is %d\n", c)

	c = a | b /* 61 = 0011 1101 */
	fmt.Printf("Line 2 - Value of c is %d\n", c)

	c = a ^ b /* 49 = 0011 0001 */
	fmt.Printf("Line 3 - Value of c is %d\n", c)

	c = a << 2 /* 240 = 1111 0000 */
	fmt.Printf("Line 4 - Value of c is %d\n", c)

	c = a >> 2 /* 15 = 0000 1111 */
	fmt.Printf("Line 5 - Value of c is %d\n", c)
}
