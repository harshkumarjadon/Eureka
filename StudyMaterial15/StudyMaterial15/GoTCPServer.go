

package main 

import (
	"fmt"
	"time"
	"io"
	"net"
	"log"
)

//_____________________________________________________________________
//_____________________________________________________________________

// c Is Passed By Value To This Function
func handleConnection( c net.Conn ) {
	// c Is Local To This Function
	defer c.Close() // Will Close Local Variable c 

	for { // Blocking Created Using Busyiness.. i.e. Busy Wait...
		_, err := io.WriteString( c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep( 1 * time.Second )
	}
}

func tcpServer() {
	listener, err := net.Listen("tcp", "localhost:8000")

	if err != nil {
		log.Fatal( err )
	}

	for {
		fmt.Printf("#")
		connection, err := listener.Accept() // Blocking Calls
		if err != nil {
			log.Print( err ) // Connection Aborted
			continue
		}
		fmt.Printf(".")

	//  In Java
	//	Create Thread
	// 	Thread connectedThread = new Thread( new Runnable( 
			// handleConnection( connection ) {
			// 	Handle Connection Body
			// } );
	//  connectedThread.start();

	 // handleConnection( connection ) 	// Blocking Function Created Due To Busy
	 go handleConnection( connection )  // Making Non Blocking Call Using Go Routine
	}
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

	fmt.Println("\nFunction : tcpServer")
	tcpServer()

	// fmt.Println("\nFunction : ")
	// fmt.Println("\nFunction : ")
	// fmt.Println("\nFunction : ")
	// fmt.Println("\nFunction : ")
}

