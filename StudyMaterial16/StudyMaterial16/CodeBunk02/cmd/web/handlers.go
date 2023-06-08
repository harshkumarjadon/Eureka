package main

import (
	"fmt"
	"net/http"
	"strconv"
	"html/template"
	"log"
	// "golang.org/x/net/html"
)

// _________________________________________________________________
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

//________________________________________________________________
// Attempt 01
// Loading A HTML File
// func home( w http.ResponseWriter, r * http.Request ) {
// 	// In Go By Default 
// 	//		"/" Pattern Act Like A Catch-All Pattern
// 	// To Restrict "/" Pattern To Domain Name Only
// 	//		Use Following Condition
// 	if r.URL.Path != "/" {
// 		http.NotFound(w, r)
// 		return
// 	}

//     // Use the template.ParseFiles() function to read the template 
//     // file into a template set. If there's an error, we log the 
//     // detailed error message and use the http.Error() function to 
//     // send a generic 500 Internal Server Error response to the user.
// 	template, err := template.ParseFiles("./ui/html/pages/home.tmpl")
// 	if err != nil {
// 		log.Println( err.Error() )
// 		http.Error( w, "Internal Server Error", 500)
// 		return
// 	}

//     // We then use the Execute() method on the template set to 
//     // write the template content as the response body. The last 
//     // parameter to Execute() represents any dynamic data that we 
//     // want to pass in, which for now we'll leave as nil.
// 	err = template.Execute( w, nil )
// 	if err != nil {
// 		log.Println( err.Error() )
// 		http.Error( w, "Internal Server Error", 500)
// 	}

// 	// w.Write( []byte("Hello World!!!! Welcome To Web!!!") )
// }

//________________________________________________________________
// Attempt 02
// Loading A HTML File
func home( w http.ResponseWriter, r * http.Request ) {
	// In Go By Default 
	//		"/" Pattern Act Like A Catch-All Pattern
	// To Restrict "/" Pattern To Domain Name Only
	//		Use Following Condition
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	files := []string {
		"./ui/html/base.tmpl",
		"./ui/html/partials/nav.tmpl",
		"./ui/html/pages/home.tmpl",
	}

    // Use the template.ParseFiles() function to read the template 
    // file into a template set. If there's an error, we log the 
    // detailed error message and use the http.Error() function to 
    // send a generic 500 Internal Server Error response to the user.
	// template, err := template.ParseFiles("./ui/html/pages/home.tmpl")
	template, err := template.ParseFiles( files... )
	if err != nil {
		log.Println( err.Error() )
		http.Error( w, "Internal Server Error", 500)
		return
	}

    // We then use the Execute() method on the template set to 
    // write the template content as the response body. The last 
    // parameter to Execute() represents any dynamic data that we 
    // want to pass in, which for now we'll leave as nil.
	err = template.ExecuteTemplate( w, "base", nil )
	if err != nil {
		log.Println( err.Error() )
		http.Error( w, "Internal Server Error", 500)
	}

	// w.Write( []byte("Hello World!!!! Welcome To Web!!!") )
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
