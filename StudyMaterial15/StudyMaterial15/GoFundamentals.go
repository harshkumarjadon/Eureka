
package main

// import "fmt"
// import "os"
import (
	"fmt"
	"os"
	"strings"
	"bufio"
)

//_______________________________________________
// PLEASE RAISE YOUR HAND, MOMENT YOU ARE DONE!!!

func helloWorld() {
	fmt.Println("Hello World!!!!")
}

//_______________________________________________
// PLEASE RAISE YOUR HAND, MOMENT YOU ARE DONE!!!

func playWithCommandLineArguments() {
	var s, sep string

	// At Index 0 Means Program Name Is Included From Arguments List
	for i := 0 ; i < len( os.Args ) ; i++ {

	// At Index 1 Means Program Name Is Excluded From Arguments List
	// for i := 1 ; i < len( os.Args ) ; i++ {
		s += sep + os.Args[i]
		sep = " "
	}

	fmt.Println("Command Line Arguments Supplied")
	fmt.Println(s)
}

//_______________________________________________
// PLEASE RAISE YOUR HAND, MOMENT YOU ARE DONE!!!

func playWithCommandLineArgumentsAgain() {
	// Following Both Statements Are Equivalent
	// var s, sep string
	s, sep := "", ""

	// For Range Loop	
	fmt.Println("Understanding range Operator")
	for index, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
		fmt.Println("Index, Value: ", index, arg)
	}

	s = ""
	// To Ignore index Value Use _ 
	// 		i.e. Variable You Don't Want To Use It 
	//				Than Name It As _ (Underscore)
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}

	for index := range os.Args[1:] {
		fmt.Println("Index: ", index)	
	}

	fmt.Println("Command Line Arguments Supplied")
	fmt.Println(s)
}

//_______________________________________________
// PLEASE RAISE YOUR HAND, MOMENT YOU ARE DONE!!!

func playWithCommandLineArgumentsOnceAgain() {
	s, sep := "", " "

	s = strings.Join( os.Args[ 1: ], sep )
	fmt.Println("Command Line Arguments Supplied")
	fmt.Println(s)
}

//_______________________________________________
// PLEASE RAISE YOUR HAND, MOMENT YOU ARE DONE!!!

func playWithUserInput() {
	// input Is Scanner
	input := bufio.NewScanner( os.Stdin )	
	// Asking Scanner Object To Scan
	//		From Standard Input ( Stdin )
	//			Stdin By Default Poining To Keyboard
	//		Scan Will Stop Reading From Stdin Moment It Gets EOF
	//			EOF Character Type Using Ctrl+D [ In Unix ]
	//			EOF Character Type Using Ctrl+Z [ In Unix ]	
	for input.Scan() {
		fmt.Println( input.Text() )
	}
}

//_______________________________________________
// PLEASE RAISE YOUR HAND, MOMENT YOU ARE DONE!!!

func playWithUserInputAndCountDuplicateLines() {
	// counts Is A Map of Key and Value
	//		Key Is string And Value Is int Type
	counts := make( map[string]int )
	// input Is Scanner
	input := bufio.NewScanner( os.Stdin )
	// Fequency Of Duplicates Words/Lines
	for input.Scan() {
		// Storing Word/Line As Key
		// Value As Word/Line Count
		counts[ input.Text() ]++
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t", n, line)
		}
	}
}


//_______________________________________________
//_______________________________________________
//_______________________________________________
//_______________________________________________
//_______________________________________________
//_______________________________________________
//_______________________________________________
//_______________________________________________
//_______________________________________________
//_______________________________________________

func main() {
	fmt.Println("\nFunction : helloWorld")
	helloWorld()

	fmt.Println("\nFunction : playWithCommandLineArguments")
	playWithCommandLineArguments()

	fmt.Println("\nFunction : playWithCommandLineArgumentsAgain")
	playWithCommandLineArgumentsAgain()

	fmt.Println("\nFunction : playWithCommandLineArgumentsOnceAgain")
	playWithCommandLineArgumentsOnceAgain()

	fmt.Println("\nFunction : playWithUserInput")
	playWithUserInput()

	// fmt.Println("\nFunction : playWithUserInputAndCountDuplicateLines")
	// playWithUserInputAndCountDuplicateLines()

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
	// fmt.Println("\nFunction : ")
	// fmt.Println("\nFunction : ")
}


