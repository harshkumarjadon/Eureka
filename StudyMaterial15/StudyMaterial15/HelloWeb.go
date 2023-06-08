
package main

import (
    "fmt"
    "net/http"
)

func welcomeHandler( w http.ResponseWriter, r * http.Request ) {
    fmt.Fprintf( w, "<h1>Welcome To My Home Page!</h1>")
    fmt.Fprintf( w, "Request URL Path: %s", r.URL.Path )
}

func thanksHandler( w http.ResponseWriter, r * http.Request ) {
    fmt.Fprintf( w, "<h1>Thanks! Visit Again</h1>")
    fmt.Fprintf( w, "You Visted URL Path: %s", r.URL.Path )
}

func main() {
    var port = 9999
    http.HandleFunc("/", welcomeHandler )
    http.HandleFunc("/thanks", thanksHandler )
    
    fmt.Println("\nStarting Server At Port :", port )
    http.ListenAndServe( ":9999", nil )
}

