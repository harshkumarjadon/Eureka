//  2 mkdir HTMLParse
//  3 cd HTMLParse/
//  4 clear
//  5 ls
//  6 ls -l
//  7 go mod init
//  8 go mod init example.com/htmlparse
//  9 ls
// 10 cat go.mod 
// 11 cat main.go
// 12 go run .
// 13 go mod tidy
// 14 cat go.mod 
// 15 go run .
 
// Write Following Code In main.go

package main

import (
	"fmt"
	"os"
	"golang.org/x/net/html"
)

// ______________________________________________________________

// visit appends to links each link found in n and returns the result.
func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

func main() {
	// FileName := os.Args[1]
	// htmlFileD := open( FileName, "r")
	// doc, err := html.Parse( htmlFileD )
	doc, err := html.Parse( os.Stdin )
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

// ______________________________________________________________
// ______________________________________________________________
// ______________________________________________________________
// ______________________________________________________________
// ______________________________________________________________
// ______________________________________________________________

