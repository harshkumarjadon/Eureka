
package main 

import (
	"fmt"
	// "time"
	"io"
	"os"
	"net"
	"log"
)

//_____________________________________________________________________
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
	fmt.Println("\nFunction : playWithTCPClient")
	playWithTCPClient()

	// fmt.Println("\nFunction : ")
	// fmt.Println("\nFunction : ")
	// fmt.Println("\nFunction : ")
	// fmt.Println("\nFunction : ")
	// fmt.Println("\nFunction : ")
}

