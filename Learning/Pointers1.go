package main

import "fmt"

// taking a structure
type Employee struct {

	// taking variables
	name  string
	empid int
}

type ArithOp func(int, int) int

func calculate(fp func(int, int) int) {
	ans := fp(3, 2)
	fmt.Printf("\n%v\n", ans)
}

// This is the same function but uses the type/fp defined above

func calculate1(fp ArithOp) {
	ans := fp(3, 2)
	fmt.Printf("\n%v\n", ans)
}

func Plus(a, b int) int {
	return a + b
}

func Minus(a, b int) int {
	return a - b
}

func Multiply(a, b int) int {
	return a * b
}

type pfunc func(string, int) int

func testCase1(toStr string, toInt int) int {
	fmt.Printf("1st Function Call %s %d\n", toStr, toInt)
	return toInt
}

func testCase2(toStr string, toInt int) int {
	fmt.Printf("2nd Function Call %s %d\n", toStr, toInt)
	return toInt
}

func testCase3(toStr string, toInt int) int {
	fmt.Printf("3rd Function Call %s %d\n", toStr, toInt)
	return toInt
}

type Delete struct {
	instance string
}

type command interface {
	DoLoop()
}

func (dev Delete) DoLoop() {
	fmt.Println("input: delete ", dev.instance)
}

func route(command string) {
	cmd := mainFuncTable[command]
	cmd.DoLoop()
}

var mainFuncTable = make(map[string]command)

// understanding pointers
func pointers() {
	var a int = 10
	var b *int = &a
	var c *int = b
	fmt.Println(&a)
	fmt.Println(a, "\n")
	fmt.Println(b)
	fmt.Println(*b)
	fmt.Println(&b, "\n")
	fmt.Println("", b)
	fmt.Println(*b)
	fmt.Println(&b, "\n")

	fmt.Println(c)
	fmt.Println(&c)
	fmt.Println(*c)
}
func main() {
	pointers()
	//mainFuncTable["delete"] = Delete{"newTestInstance"}
	//cmd := mainFuncTable["delete"]
	//cmd := Delete{"newTestInstance"}
	//cmd.DoLoop()
	//fmt.Println(route("new"))
	// calculate(Plus)
	// calculate(Minus)
	// calculate(Multiply)

	// calculate1(Plus)
	// calculate1(Minus)
	// calculate1(Multiply)

	// funcArray := []pfunc{testCase1, testCase2, testCase3}

	// for n := range funcArray {
	// 	result := funcArray[n]("Test", n)
	// 	fmt.Printf("Run Test Case #%d reference %v result %d\n", n, funcArray[n], result)
	// }

	// creating the instance of the Employee struct type
	emp := Employee{"ABC", 19078}

	// Here, it is the pointer to the struct
	pts := &emp
	fmt.Println(pts)

	// accessing the struct fields(liem employee's name)
	// using a pointer but here we are not using dereferencing explicitly
	fmt.Println(pts.name)

	// same as above by explicitly using dereferencing concept
	// means the result will be the same
	fmt.Println((*pts).name)

	// updating the value of name
	pts.name = "XYZ"
	fmt.Println(pts)
}
