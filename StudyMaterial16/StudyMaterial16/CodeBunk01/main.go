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
	// "golang.org/x/net/html"
	"log"
	"strconv"
)

// _________________________________________________________________

func HelloWeb() string {
	return "Starting Web Programming...."
}

// _________________________________________________________________

// func ListenAndServe(addr string, handler Handler) error

// type Handler interface {
// 		ServeHTTP(ResponseWriter, *Request)
// }

// type ResponseWriter interface {
// 	// Header returns the header map that will be sent by
// 	// WriteHeader. The Header map also is the mechanism with which
// 	// Handlers can set HTTP trailers.
// 	//
// 	// Changing the header map after a call to WriteHeader (or
// 	// Write) has no effect unless the HTTP status code was of the
// 	// 1xx class or the modified headers are trailers.
// 	//
// 	// There are two ways to set Trailers. The preferred way is to
// 	// predeclare in the headers which trailers you will later
// 	// send by setting the "Trailer" header to the names of the
// 	// trailer keys which will come later. In this case, those
// 	// keys of the Header map are treated as if they were
// 	// trailers. See the example. The second way, for trailer
// 	// keys not known to the Handler until after the first Write,
// 	// is to prefix the Header map keys with the TrailerPrefix
// 	// constant value. See TrailerPrefix.
// 	//
// 	// To suppress automatic response headers (such as "Date"), set
// 	// their value to nil.
// 	Header() Header

// 	// Write writes the data to the connection as part of an HTTP reply.
// 	//
// 	// If WriteHeader has not yet been called, Write calls
// 	// WriteHeader(http.StatusOK) before writing the data. If the Header
// 	// does not contain a Content-Type line, Write adds a Content-Type set
// 	// to the result of passing the initial 512 bytes of written data to
// 	// DetectContentType. Additionally, if the total size of all written
// 	// data is under a few KB and there are no Flush calls, the
// 	// Content-Length header is added automatically.
// 	//
// 	// Depending on the HTTP protocol version and the client, calling
// 	// Write or WriteHeader may prevent future reads on the
// 	// Request.Body. For HTTP/1.x requests, handlers should read any
// 	// needed request body data before writing the response. Once the
// 	// headers have been flushed (due to either an explicit Flusher.Flush
// 	// call or writing enough data to trigger a flush), the request body
// 	// may be unavailable. For HTTP/2 requests, the Go HTTP server permits
// 	// handlers to continue to read the request body while concurrently
// 	// writing the response. However, such behavior may not be supported
// 	// by all HTTP/2 clients. Handlers should read before writing if
// 	// possible to maximize compatibility.
// 	Write([]byte) (int, error)

// 	// WriteHeader sends an HTTP response header with the provided
// 	// status code.
// 	//
// 	// If WriteHeader is not called explicitly, the first call to Write
// 	// will trigger an implicit WriteHeader(http.StatusOK).
// 	// Thus explicit calls to WriteHeader are mainly used to
// 	// send error codes or 1xx informational responses.
// 	//
// 	// The provided code must be a valid HTTP 1xx-5xx status code.
// 	// Any number of 1xx headers may be written, followed by at most
// 	// one 2xx-5xx header. 1xx headers are sent immediately, but 2xx-5xx
// 	// headers may be buffered. Use the Flusher interface to send
// 	// buffered data. The header map is cleared when 2xx-5xx headers are
// 	// sent, but not with 1xx headers.
// 	//
// 	// The server will automatically send a 100 (Continue) header
// 	// on the first read from the request body if the request has
// 	// an "Expect: 100-continue" header.
// 	WriteHeader(statusCode int)
// }

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

// # Fixed Path and Subtree Patterns

// Go’s servemux supports two different types of URL patterns: 
// fixed paths and subtree paths. 
// Fixed paths don’t end with a trailing slash, whereas 
// Subtree paths do end with a trailing slash.

// Fixed Path Example :
//		"/snippet/view" and "/snippet/create" — are Fixed paths

// Subtree Path Example :
//		/static/ Will Behave Like /static/*
//		/ 
// 		The "/" pattern is acting like a catch-all. 
//		The pattern essentially means match a single slash, 
//		followed by anything (or nothing at all).

// _________________________________________________________________
// Attempt 01
// func codeView( w http.ResponseWriter, r * http.Request ) {
// 	helloWorldCode := "\tfunc main() {\n\t\t fmt.Println(\"Hello World\") \n\t}\n"
// 	w.Write( []byte("Displaying Code !!!! \n") )
// 	w.Write( []byte(helloWorldCode) )
// }

// _________________________________________________________________
// Attempt 02: URL Query Strings
//		/code/view?id=288

// TEST URL
	// curl http://localhost:4000/code/view?id=8999
	// curl http://localhost:4000/code/view
	// curl http://localhost:4000/code/view?id
	// curl http://localhost:4000/code/view?id=1
	// curl http://localhost:4000/code/view?id=-10
	// curl http://localhost:4000/code/view?id=000
	// curl http://localhost:4000/code/view?id=DingDong

func codeView( w http.ResponseWriter, r * http.Request ) {
 	id, err := strconv.Atoi( r.URL.Query().Get("id") )

 	// Validation Of Query String Parameters
	if err != nil || id < 1 {
		http.NotFound( w, r )
		return
	}

	fmt.Fprintf( w, "Displaying Code with ID %d\n", id)
	w.Write( []byte("Displaying Code !!!! \n") )
}

// _________________________________________________________________

// The Interface http.ResponseWriter parameter provides methods 
// for assembling a HTTP response and sending it to the user, and 
// the *http.Request parameter is a pointer to a struct which holds 
// information about the current request (like the HTTP method and 
// the URL being requested).

// Testing URLs
// curl -i -X PUT http://localhost:4000/code/create
// curl -i -X GET http://localhost:4000/code/create
// curl -i -X POST http://localhost:4000/code/create

//______________________________________________________________
// Attemp 01
// func codeCreate( w http.ResponseWriter, r * http.Request ) {
// 	if r.Method != "POST" {
// 		// If it's not, use the w.WriteHeader() method to send a 405 status
// 		// code and the w.Write() method to write a "Method Not Allowed"
// 		// response body. We then return from the function so that the
// 		// subsequent code is not executed.
// 		w.Header().Set("Allow", "POST")
// 		// w.WriteHeader( 405 ) // Method Not Allowed
// 		w.Write( []byte("Displaying Code !!!! \n") )
// 		return
// 	}

// 	w.Write( []byte(">>>>> Method Not Allowed <<<< \n") )
// }

//______________________________________________________________
// Attemp 02: Alternative To WriteHeader Method

// func codeCreate( w http.ResponseWriter, r * http.Request ) {
// 	if r.Method != "POST" {
// 		// If it's not, use the w.WriteHeader() method to send a 405 status
// 		// code and the w.Write() method to write a "Method Not Allowed"
// 		// response body. We then return from the function so that the
// 		// subsequent code is not executed.
// 		w.Header().Set("Allow", "POST")
// 		// w.WriteHeader( 405 ) // Method Not Allowed

// 		// Alternative To WriteHeader Method
// 		// 		Use the http.Error() function to send a 405 
// 		//		status code and "Method Not Allowed" string as 
// 		//		the response body.
// 		http.Error( w, "Method Not Allowed", 405)
// 		w.Write( []byte("Displaying Code !!!! \n") )
// 		return
// 	}

// 	w.Write( []byte(">>>>> Method Not Allowed <<<< \n") )
// }

//______________________________________________________________
// Attempt 03 : Using net/http Constants 

func codeCreate( w http.ResponseWriter, r * http.Request ) {
	if r.Method !=  http.MethodPost { //"POST" { Hard Coding Values Is Bad Programming
		// If it's not, use the w.WriteHeader() method to send a 405 status
		// code and the w.Write() method to write a "Method Not Allowed"
		// response body. We then return from the function so that the
		// subsequent code is not executed.
		w.Header().Set("Allow", http.MethodPost)
		// w.WriteHeader( 405 ) // Method Not Allowed

		// Alternative To WriteHeader Method
		// 		Use the http.Error() function to send a 405 
		//		status code and "Method Not Allowed" string as 
		//		the response body.
		http.Error( w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	
	w.Write( []byte("Displaying Code !!!! \n") )
}

//______________________________________________________________
// Attempt 04 : Using net/http Constants 

// func codeCreate( w http.ResponseWriter, r * http.Request ) {
// 	if r.Method !=  http.MethodPost { //"POST" { Hard Coding Values Is Bad Programming
// 		// If it's not, use the w.WriteHeader() method to send a 405 status
// 		// code and the w.Write() method to write a "Method Not Allowed"
// 		// response body. We then return from the function so that the
// 		// subsequent code is not executed.
// 		w.Header().Set("Allow", http.MethodPost)

// 		// Set a new cache-control header. If an existing "Cache-Control" header exists
// 		// it will be overwritten.
// 		w.Header().Set("Cache-Control", "public, max-age=31536000")

// 		// In contrast, the Add() method appends a new "Cache-Control" header and can
// 		// be called multiple times.
// 		w.Header().Add("Cache-Control", "public")
// 		w.Header().Add("Cache-Control", "max-age=31536000")

// 		// Delete all values for the "Cache-Control" header.
// 		w.Header().Del("Cache-Control")

// 		// Retrieve the first value for the "Cache-Control" header.
// 		w.Header().Get("Cache-Control")

// 		// Retrieve a slice of all values for the "Cache-Control" header.
// 		w.Header().Values("Cache-Control")

// 		// Alternative To WriteHeader Method
// 		// 		Use the http.Error() function to send a 405 
// 		//		status code and "Method Not Allowed" string as 
// 		//		the response body.
// 		http.Error( w, "Method Not Allowed", http.StatusMethodNotAllowed)
// 		return
// 	}
	
// 	w.Write( []byte("Displaying Code !!!! \n") )
// }

// _________________________________________________________________
// _________________________________________________________________
// _________________________________________________________________
// _________________________________________________________________

func main() {
	helloWeb := HelloWeb()
	fmt.Println( helloWeb )

// GOOD APPROACH : CREATING LOCAL MULTIPLEXER
	// Creating Routing Multiplexer
	//		Routing Table

// Each time the server receives a new HTTP request it will pass the 
// request on to the servemux and — in turn — the servemux will check 
// the URL path and dispatch the request to the matching handler.

	mux := http.NewServeMux()
	// Use the http.NewServeMux() function to initialize a new servemux, 
	// Then register the ACTION function as the handler for the URL pattern.

	// Adding Routes With Path And Action
	// 		Path Is "/" 			and Action Is home
	// 		Path Is "/code/view" 	and Action Is codeView
	// 		Path Is "/code/create" 	and Action Is codeCreate
	mux.HandleFunc( "/", home )
	mux.HandleFunc( "/code/view", 	codeView )
	mux.HandleFunc( "/code/create", codeCreate )

	log.Println("Starting Server On Port: 4000")

	// Use the http.ListenAndServe() function to start a new web server.
	// We pass in two parameters: the TCP network address to listen on 
	// (in this case":4000") and the servemux we just created. 
	// If http.ListenAndServe() returns an error
	// we use the log.Fatal() function to log the error message and exit. 
	// Note that any error returned by http.ListenAndServe() is always non-nil.
	
	err := http.ListenAndServe(":4000", mux ) // Registering Multiplexer

	log.Fatal( err )

// ATLTERNATVE APPROACH : NOT RECOMMENDED
	// 		Not Recommended Practice Because Global Variables Are Bad Things
	// Behind the scenes, these functions register their routes 
	// with something called the DefaultServeMux. There’s nothing 
	// special about this — it’s just regular servemux like we’ve 
	// already been using, but which is initialized by default and 
	// stored in a net/http global variable
	// http.HandleFunc( "/", 				home )
	// http.HandleFunc( "/code/view", 		codeView )
	// http.HandleFunc( "/code/create", 	codeCreate )

	// 	Path Is "/closure" 	and Action Is Closure Syntax
	// 		Registering Action Using Closure Syntax
	// http.HandleFunc("/closure", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	// })

	// log.Println("Starting Server On Port: 4000")

	// ListenAndServe starts an HTTP server with a given address and handler. 
	// The handler is usually nil, which means to use DefaultServeMux. 
	// Handle and HandleFunc add handlers to DefaultServeMux:
	
	// err := http.ListenAndServe(":4000", nil )

	// func ListenAndServe(addr string, handler Handler) error
	// ListenAndServe listens on the TCP network address addr and then 
	// calls Serve with handler to handle requests on incoming connections. 
	// Accepted connections are configured to enable TCP keep-alives.
	// The handler is typically nil, in which case the DefaultServeMux is used.
	// ListenAndServe always returns a non-nil error.

	// log.Fatal( err )
}

// _________________________________________________________________

// Servemux Features and Quirks

// In Go’s servemux, longer URL patterns always take precedence over 
// shorter ones. So, if a servemux contains multiple patterns which 
// match a request, it will always dispatch the request to the handler 
// corresponding to the longest pattern. This has the nice side-effect that
// you can register patterns in any order and it won’t change how the 
// servemux behaves.

// Request URL paths are Automatically Sanitized. 
// If the request path contains any . or .. elements or repeated slashes, 
// the user will automatically be redirected to an equivalent clean URL. 
// For example, if a user makes a request to /foo/bar/..//baz 
//		they will automatically be sent a 301 Permanent Redirect 
//		to /foo/baz instead.

// If a subtree path has been registered and a request is received 
// for that subtree path without a trailing slash, then the user 
// will automatically be sent a 301 Permanent Redirect to the
// subtree path with the slash added. 

// For example, if you have registered the subtree path /foo/ , 
//		then any request to /foo will be redirected to /foo/ .

// _________________________________________________________________

// # Host Name Matching

// It’s possible to include host names in your URL patterns. 
// This can be useful when you want to redirect all HTTP requests 
// to a canonical URL, or if your application is acting as the back end for
// multiple sites or services. For example:

// mux := http.NewServeMux()
// mux.HandleFunc("foo.example.org/", fooHandler)
// mux.HandleFunc("bar.example.org/", barHandler)
// mux.HandleFunc("/baz", bazHandler)

// _________________________________________________________________

// What about RESTful routing?
// It’s important to acknowledge that the routing functionality 
// provided by Go’s servemux is pretty lightweight. It doesn’t support 
// routing based on the request method, it doesn’t support clean URLs
// with variables in them, and it doesn’t support regexp-based patterns. 

// _________________________________________________________________

