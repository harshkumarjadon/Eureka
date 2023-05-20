package main

import (
	"bytes"
	"fmt"
	"math"
	"math/big"
	"strconv"
	"strings"
)

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
	fmt.Println(piAgain1)

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

	fmt.Println(piggyBank == 0.3)

	// Function : playWithFloatingPointTypes
	// false

	// BAD CODE : Buggy Logic
	//		Serious Security Threat :  Validation Can Be Bypassed
	if piggyBank == 0.3 {
		fmt.Println("piggyBank Have Dollers : ", piggyBank)
	} else {
		fmt.Println("piggyBank Is Dead...")
	}

	// GOOD CODE
	if math.Abs(piggyBank-0.3) < 0.000001 {
		fmt.Println("piggyBank Have Dollers : ", piggyBank)
	} else {
		fmt.Println("piggyBank Is Dead...")
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

	var dollar = '$' // Inferred Type Is int32 and Binded With dollar Variable
	fmt.Printf("\nValue %v and Type %T", dollar, dollar)

	//cannot use 900 (untyped int constant) as int8 value in
	//	variable declaration (overflows)
	// var b int8 = 900
	var b int8 = 127
	fmt.Printf("\nValue %v and Type %T", b, b)

	var red, green, blue = 0, 141, 213         // int Type Values In Base 10
	var red1, green1, blue1 = 0x00, 0x8d, 0xd5 // int Type Values In Base 16

	fmt.Printf("\nValue %v and Type %T", red, red)
	fmt.Printf("\nValue %v and Type %T", red1, red1)

	fmt.Printf("\nValue %v and Type %T", green, green)
	fmt.Printf("\nValue %v and Type %T", blue, blue)

	fmt.Printf("\nValue %v and Type %T", green1, green1)
	fmt.Printf("\nValue %v and Type %T", blue1, blue1)

	age := 1415
	marsAge := float64(age) // Type Casting
	fmt.Println(marsAge)

	fmt.Println("String Converted Value: ", strconv.Itoa(age))

	fmt.Println(math.MinInt16)
	fmt.Println(math.MaxInt16)

	fmt.Println(math.MinInt32)
	fmt.Println(math.MaxInt32)
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

	var distanceStart = 41.3e12 // float64
	fmt.Println("Alpha Centauri is", distanceStart, "km away.")

	var distance int64 = 41.3e12 // Exponential Notation  41300000000000
	fmt.Println("Alpha Centauri is", distance, "km away.")

	days := distance / lightSpeed / secondsPerDay
	fmt.Println("That is", days, "days of travel at light speed.")

	bigNumber1 := big.NewInt(888880000000777)
	bigNumber2 := big.NewInt(111)

	fmt.Println(bigNumber1)
	fmt.Println(bigNumber2)
}

//_____________________________________________________________________

// Two Complex Types In Go
//		a + ib Form
//		complex64
//			Means a And b Are float32 Type
//		complex128
//			Means a And b Are float64 Type

func playWithComlexTypes() {
	var x complex128 = complex(1, 2) // 1 + 2i
	var y complex128 = complex(3, 4) // 3 + 4i

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(x + y)
	fmt.Println(x * y)
	fmt.Println(real(x))
	fmt.Println(imag(x))

	fmt.Println(1i * 1i)

	xx := 1 + 2i
	yy := 3 + 4i

	fmt.Println(xx)
	fmt.Println(yy)
	fmt.Println(xx + yy)
	fmt.Println(xx * yy)
	fmt.Println(real(xx))
	fmt.Println(imag(xx))
}

//_____________________________________________________________________

// string Type In Go
//
//	BuiltIn Type In Go
func basename(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:] // Taking String Slice
			break
		}
	}

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i] // Taking String Slice
			break
		}
	}
	return s
}

func basenameAgain(s string) string {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]

	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

func playWithBaseName() {
	fmt.Println(basename("/media/WorkArea/Trainings.Running/Coforge/Progress"))
	fmt.Println(basename("/media/WorkArea/Trainings.Running/Coforge/Progress/GoTypes.go"))

	fmt.Println(basenameAgain("/media/WorkArea/Trainings.Running/Coforge/Progress"))
	fmt.Println(basenameAgain("/media/WorkArea/Trainings.Running/Coforge/Progress/GoTypes.go"))
}

//_____________________________________________________________________

// func LastIndex(s, substr string) int
// 	LastIndex returns the index of the last instance of substr in s,
// 	or -1 if substr is not present in s.

func playWithLastIndex() {
	fmt.Println(strings.Index("go gopher", "go"))
	fmt.Println(strings.LastIndex("go gopher", "go"))
	fmt.Println(strings.LastIndex("go gopher", "rodent"))
}

// Function : playWithLastIndex
// 0
// 3
// -1

//_____________________________________________________________________

// In C
//		There Is No String Type
//		There Is String Literal " "
//			String Literal Is A ASCII String
// In Go
//		string Type Is Unicode String

func playWithStringType() {
	s := "Hello World!"

	fmt.Println(s)
	fmt.Println("String Length: ", len(s))
	//  Hello World!
	//  String Length:  12

	// String Constructing By Using One Composed Unicode Character
	// 한 Unicode Character Have Unicode Point \uD55C
	s1 := "Hello World!\uD55C"
	fmt.Println(s1)
	fmt.Println("String Length: ", len(s1))

	// String Constructing By Composing Smaller Unicode Character
	// 한 Unicode Character Composed Using Smaller Unicode Characters
	// 		Using These Unicode Characters ᄒ ᅡ   ᆫ
	// Smaller Unicode Character ᄒ ᅡ   ᆫ Having Unicode Points
	// 	 ᄒ   \u1112
	// 	ᅡ   \u1161
	//   ᆫ  \u11AB

	// Hello World!한
	s2 := "Hello World!\u1112\u1161\u11AB"
	fmt.Println(s2)
	fmt.Println("String Length: ", len(s2))
	fmt.Println("\u1112")
	fmt.Println("\u1161")
	fmt.Println("\u11AB")

	//       0123456789
	// s := "Hello World!"
	fmt.Println(s[0], s[7])
	fmt.Println(s[0:5])
	fmt.Println(s[:5])
	fmt.Println(s[7:])
	fmt.Println(s[:])

	fmt.Println("Goodbye " + s[5:])
	// 72 111
	// Hello
	// Hello
	// orld!
	// Hello World!
	// Goodbye  World!

	ss := "Left Foot"
	tt := ss // String Copy

	fmt.Println(tt)
	fmt.Println(ss)

	tt += ", and Right Foot"

	fmt.Println(tt)
	fmt.Println(ss)
	// Left Foot
	// Left Foot
	// Left Foot, and Right Foot
	// Left Foot

	// Compile Time Error:
	//		cannot assign to s[0] (value of type byte)
	// s[0] = "L"
	fmt.Println(s)
}

// Function : playWithStringType
// Hello World!한
// String Length:  15
// Hello World!한
// String Length:  15

//_____________________________________________________________________

// string Package Documentation
// https://pkg.go.dev/strings

func playWithStringsFunctions() {

	fmt.Println("Contains:  ", strings.Contains("test", "es"))
	fmt.Println("Count:     ", strings.Count("test", "t"))
	fmt.Println("HasPrefix: ", strings.HasPrefix("test", "te"))
	fmt.Println("HasSuffix: ", strings.HasSuffix("test", "st"))
	fmt.Println("Index:     ", strings.Index("test", "e"))
	fmt.Println("Join:      ", strings.Join([]string{"a", "b"}, "-"))
	fmt.Println("Repeat:    ", strings.Repeat("a", 5))
	fmt.Println("Replace:   ", strings.Replace("foo", "o", "0", -1))
	fmt.Println("Replace:   ", strings.Replace("foo", "o", "0", 1))
	fmt.Println("Split:     ", strings.Split("a-b-c-d-e", "-"))
	fmt.Println("ToLower:   ", strings.ToLower("TEST"))
	fmt.Println("ToUpper:   ", strings.ToUpper("test"))

}

//_____________________________________________________________________

// String Type Design Choices In C/Go Language

// In C Language
// There Is No String Type In C
// String Value Denoted By ""
// String Is Sequence Of ASCII Characters Stored In char Type Array
// In C
//		String Follows ASCII Coding
//		String Ends With '\0' Null ASCII Character In Memory

// In Go Language
// A string is an immutable sequence of bytes.
// Strings may contain arbitrary data, including bytes with value 0,
// but usually they contain human-readable text.

// Text strings are conventionally interpreted as UTF-8-encoded
// sequences of Unicode code points (runes)
// It Stores Unicode Characters

// Internal Structure Of string Type In Go Langauge

// type _string struct {
// 		elements *byte // underlying bytes
// 		len      int   // number of bytes
// }

// The built-in len function returns the number of bytes (not runes)
// in a string, and the index operation s[i] retrieves the i-th
// byte of string s, where 0 ≤ i < len(s).

// The i-th byte of a string is not necessarily the i-th character
// of a string, because the UTF-8 encoding of a non-ASCII code point
// requires two or more bytes.

//___________________________________________________________________
// EXPERIMENT FOLLOWING CODE, MOMENT DONE, PLEASE RAISE YOUR HAND!!!

func insertCommas(s string) string {
	n := len(s)

	if n <= 3 {
		return s
	}
	return insertCommas(s[:n-3]) + "," + s[n-3:]
}

func playWithInsertCommas() {
	fmt.Println(insertCommas("12345678"))
	fmt.Println(insertCommas("999"))
	fmt.Println(insertCommas("6789999"))
	fmt.Println(insertCommas("89389438493849384398439483"))
	fmt.Println(insertCommas("123456788888888888"))
}

//___________________________________________________________________
// EXPERIMENT FOLLOWING CODE, MOMENT DONE, PLEASE RAISE YOUR HAND!!!

// type byte
// type byte = uint8
//		Unsigned int8 i.e. Range Is [ 0, 255 ]
// byte is an alias for uint8 and is equivalent to uint8 in all ways.
// It is used, by convention, to distinguish byte values from
// 8-bit unsigned integer values.

// Package bytes implements functions for the manipulation
//	of byte slices.
// It is analogous to the facilities of the strings package.

// func (b *Buffer) WriteByte(c byte) error
//  WriteByte appends the byte c to the buffer, growing the buffer as needed.
//	The returned error is always nil, but is included to match
//	bufio.Writer's WriteByte.
//  If the buffer becomes too large, WriteByte will panic with ErrTooLarge.

// func (*Buffer) WriteString ¶
// func (b *Buffer) WriteString(s string) (n int, err error)
//  WriteString appends the contents of s to the buffer,
//	growing the buffer as needed.
//	The return value n is the length of s; err is always nil.
//  If the buffer becomes too large, WriteString will panic with ErrTooLarge.

// func Fprintf(w io.Writer, format string, a ...any) (n int, err error)
//  Fprintf formats according to a format specifier and writes to w.
//	It returns the number of bytes written and any write error encountered.

func intsToString(values []int) string {
	var buffer bytes.Buffer
	buffer.WriteByte('[')

	// range Operator Will Return List Of Tuples
	//		Each Tuple Having Two Members (index, value)
	for index, value := range values {
		fmt.Printf("\nAt index: %d value: %v", index, value)

		if index > 0 {
			buffer.WriteString(", ")
		}
		fmt.Fprintf(&buffer, "%d", value)
	}
	buffer.WriteByte(']')
	return buffer.String()
}

func playWithIntToString() {
	fmt.Println("\nString : ", intsToString([]int{10, 20, 30}))

	fmt.Println("\nString : ",
		intsToString([]int{10, 20, 30, 99, 100, 101, 8000}))

	fmt.Println("\nString : ",
		intsToString([]int{1000, 2000, 3000, 9000, 7000, 6000, 8000}))
}

//_____________________________________________________________________
// EXPERIMENT FOLLOWING CODE, MOMENT DONE, PLEASE RAISE YOUR HAND!!!

func playWitStringUnicode() {
	const something = `⌘`

	fmt.Printf("\nPlain String: %s", something)
	fmt.Printf("\nLength Of %s: %d", something, len(something))

	fmt.Println()

	fmt.Printf("\nQuoted String: ")
	fmt.Printf("%q", something)
	fmt.Println()

	fmt.Printf("\nHex Bytes: ")
	for i := 0; i < len(something); i++ {
		fmt.Printf(" %x ", something[i])
	}

	fmt.Println()

	var ding = "Ding"
	var dingdong = "Ding Dong"
	fmt.Printf("\nLength Of %s: %d", ding, len(ding))
	fmt.Printf("\nLength Of %s: %d", dingdong, len(dingdong))

	fmt.Println()

	var ding1 = "Ding⌘"
	var dingdong1 = "Ding Dong⌘"
	fmt.Printf("\nLength Of %s: %d", ding1, len(ding1))
	fmt.Printf("\nLength Of %s: %d", dingdong1, len(dingdong1))

	fmt.Println()
}

//_____________________________________________________________________
// EXPERIMENT FOLLOWING CODE, MOMENT DONE, PLEASE RAISE YOUR HAND!!!

func playWitStringUnicodeMore() {
	somethingAgain := "Hello, 한瘔"
	somethingAgain1 := "Hello, 한瘔"

	if somethingAgain == somethingAgain1 {
		fmt.Println("Matched")
	}
	fmt.Println(strings.Compare(somethingAgain, somethingAgain1))
	fmt.Println(somethingAgain)
	fmt.Println(len(somethingAgain))
	fmt.Printf("Hex Bytes: ")
	for i := 0; i < len(somethingAgain); i++ {
		fmt.Printf(" %x ", somethingAgain[i])
	}

	fmt.Println()

	for index, character := range somethingAgain {
		fmt.Printf("\n%d\t %q \t %x", index, character, character)
	}

	characterCount := 0
	for _, _ = range somethingAgain {
		characterCount++
	}

	fmt.Println("\nUnicode Character Count : ", characterCount)

	characterCount = 0
	for range somethingAgain {
		characterCount++
	}

	fmt.Println("\nUnicode Character Count : ", characterCount)

	x := 123

	// y Is string Type
	y := fmt.Sprintf("%d", x)
	fmt.Println(y)

	fmt.Println("String Conv " + strconv.Itoa(x))
	fmt.Println()
	pp := "79"
	fmt.Println(strconv.Atoi("8970"))
	fmt.Println(strconv.Atoi(pp))
}

//_____________________________________________________________________
//_____________________________________________________________________
//_____________________________________________________________________
//_____________________________________________________________________
//_____________________________________________________________________
//_____________________________________________________________________
//_____________________________________________________________________
//_____________________________________________________________________
// EXPERIMENT FOLLOWING CODE, MOMENT DONE, PLEASE RAISE YOUR HAND!!!

func main() {
	// fmt.Println("\nFunction : playWithTypes")
	// playWithTypes()

	// fmt.Println("\nFunction : playWithFloatingPointTypes")
	// playWithFloatingPointTypes()

	// fmt.Println("\nFunction : playWithTypesAgain")
	// playWithTypesAgain()

	// fmt.Println("\nFunction : playWithLargeNumbers")
	// playWithLargeNumbers()

	// fmt.Println("\nFunction : playWithComlexTypes")
	// playWithComlexTypes()

	// fmt.Println("\nFunction : playWithBaseName")
	// playWithBaseName()

	// fmt.Println("\nFunction : playWithLastIndex")
	// playWithLastIndex()

	// fmt.Println("\nFunction : playWithStringType")
	// playWithStringType()

	// fmt.Println("\nFunction : playWithInsertCommas")
	// playWithInsertCommas()

	// fmt.Println("\nFunction : playWithIntToString")
	// playWithIntToString()

	// fmt.Println("\nFunction : playWitStringUnicode")
	// playWitStringUnicode()

	fmt.Println("\nFunction : playWitStringUnicodeMore")
	playWitStringUnicodeMore()

	// fmt.Println("\nFunction : ")
	// fmt.Println("\nFunction : ")
	// fmt.Println("\nFunction : ")
	// fmt.Println("\nFunction : ")
	// fmt.Println("\nFunction : ")
	// fmt.Println("\nFunction : ")
	// fmt.Println("\nFunction : ")
}
