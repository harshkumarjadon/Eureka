1. Create Hello Directory
2. cd Hello
3. Create main.go File In Hello Directory
4. Write Following Code In main.go

    1
    2 package main
    3
    4 //  Package
    5 import "fmt"
    6 import "rsc.io/quote"
    7
    8 func Hello() string {
    9     return quote.Hello()
   10 }
   11
   12 func main() {
   13     fmt.Println("Hello World!!!!")
   14     Hello()
   15 }
   16

5. To Make Directory As go  Module Run Following Command
	 go mod init example.com/hello

6. It Will Create go.mod File

7. To Resolve Dependencies Run Following Command
	 
   go mod tidy

8. To Run Code
	 go run .
	 go run main.go
