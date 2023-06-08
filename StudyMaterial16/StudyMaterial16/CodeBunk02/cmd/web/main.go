package main

import (
	"fmt"
	"net/http"
	"log"

	// "golang.org/x/net/html"
)

// _________________________________________________________________

func HelloWeb() string {
	return "Starting Web Programming...."
}

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

