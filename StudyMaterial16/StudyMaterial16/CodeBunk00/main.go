// 2011  go mod init example.com/codebunk
// 2012  ls
// 2013  clear
// 2014  ls
// 2015  touch main.go
// 2016  ls -l
// 2017  ls
// 2018  go run .
// 2019  clear
// 2020  ls
// 2021  cat main.go 

// Ways To Start Server Code
// go run .
// go run main.go
// go run example.com/codebunk

// Ways To Get URL Data
//	curl http://localhost:4000/

package main

import (
	"fmt"
	"net/http"
	"log"
)

func HelloWeb() string {
	return "Starting Web Programming...."
}

func home( w http.ResponseWriter, r * http.Request ) {
	// In Go By Default 
	//		"/" Pattern Act Like A Catch-All Pattern
	// To Restrict "/" Pattern To Domain Name Only
	//		Use Following Condition
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write( []byte("Hello World!!!! Welcome To Web!!!") )
}

func codeView( w http.ResponseWriter, r * http.Request ) {
	helloWorldCode := "\tfunc main() {\n\t\t fmt.Println(\"Hello World\") \n\t}\n"
	w.Write( []byte("Displaying Code !!!! \n") )
	w.Write( []byte(helloWorldCode) )
}

func codeCreate( w http.ResponseWriter, r * http.Request ) {
// curl -i -X PUT http://localhost:4000/code/create
// curl -i -X GET http://localhost:4000/code/create
// curl -i -X POST http://localhost:4000/code/create

	if r.Method != "POST" {
		w.WriteHeader( 405 ) // Method Not Allowed
		w.Write( []byte("Displaying Code !!!! \n") )
		return
	}

	w.Write( []byte(">>>>> Method Not Allowed <<<< \n") )
}

func main() {
	helloWeb := HelloWeb()
	fmt.Println( helloWeb )

	// Creating Routing Multiplexer
	mux := http.NewServeMux()

	// Adding Routes With Path And Action
	// 		Path Is "/" and Action Is home
	// 		Path Is "/code/view" 	and Action Is codeView
	// 		Path Is "/code/create" 	and Action Is codeCreate
	mux.HandleFunc( "/", home )
	mux.HandleFunc( "/code/view", 	codeView )
	mux.HandleFunc( "/code/create", codeCreate )

	log.Println("Starting Server On Port: 4000")
	err := http.ListenAndServe(":4000", mux )

	log.Fatal( err )
}

