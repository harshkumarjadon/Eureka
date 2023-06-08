
package main 

import (
	"fmt"
	"time"
	"io"
	"os"
	"net"
	"log"
)

//_____________________________________________________________________

func spinner( delay time.Duration ) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep( delay )
		}
	} // Infinite Loop

	// Will Not Return From This Function
}

func fib( x int ) int {
	if x < 2 {
		return x 
	}
	return fib( x - 1 ) + fib( x - 2 )
}

func playWithFibonacci() {
	// spinner( 100 * time.Millisecond ) // Blocking Function Call
	// Making It As Go Routine/Coroutine
	//		It Runs Concurrently
	go spinner( 100 * time.Millisecond ) // Non Blocking Function Call

	const n = 40
	fibN := fib( n ) // Slow
	fmt.Printf("\nFibonnacci (%d) = %d\n", n, fibN)
}

//_____________________________________________________________________

func doSomething( from string ) {
	for i := 0 ; i < 3 ; i++ {
		fmt.Println( from, " : ", i )
		time.Sleep( time.Second * 2 )
	}
}

func playWithGoRoutines() {
	fmt.Println( time.Now().Format( time.RFC850) ) 

	// doSomething("Direct")
	go doSomething("Direct")

	// doSomething("Once More Called")
	go doSomething("Once More Called")

	messageClosure := func( message string ) {
		fmt.Println( message )
		time.Sleep( time.Second * 3 )
	}

	go messageClosure("Here Goes Message Closure Call...")

	// func( message string ) {
	// 	fmt.Println( message )
	// 	time.Sleep( time.Second * 3 )
	// }("Here Goes Closure Call...")

	go func( message string ) {
		fmt.Println( message )
		time.Sleep( time.Second * 3 )
	}("Here Goes Messag Again Closure Call...")

	time.Sleep( time.Second * 10 )
	fmt.Println("Done")
	fmt.Println( time.Now().Format(time.RFC850))
}

//_____________________________________________________________________

// Polymorphic Function
func mustCopy( dst io.Writer, src io.Reader ) {
	if _, err := io.Copy( dst, src ); err != nil {
		log.Fatal( err )
	}
}

func tcpClient() {
	// localhost ???
	connection, err := net.Dial("tcp", "localhost:8000")

	if err != nil {
		log.Fatal( err )
	}
	defer connection.Close()

	// mustCopy( os.Stdout, connection )
	go mustCopy( os.Stdout, connection )
	mustCopy( connection, os.Stdin )
	
	// connection.Close()
}

func playWithTCPClient() {
	fmt.Println("\nFunction: tcpClient")
	tcpClient()
}

//_____________________________________________________________________
//_____________________________________________________________________
//_____________________________________________________________________
//_____________________________________________________________________
//_____________________________________________________________________
//_____________________________________________________________________
//_____________________________________________________________________
// EXPERIMENT FOLLOWING CODE, MOMENT DONE, PLEASE RAISE YOUR HAND!!!

func main() {
	fmt.Println("\nFunction : playWithFibonacci")
	playWithFibonacci()

	fmt.Println("\nFunction : playWithGoRoutines")
	playWithGoRoutines()

	fmt.Println("\nFunction : playWithTCPClient")
	playWithTCPClient()

	// fmt.Println("\nFunction : ")
	// fmt.Println("\nFunction : ")
	// fmt.Println("\nFunction : ")
	// fmt.Println("\nFunction : ")
	// fmt.Println("\nFunction : ")
}

