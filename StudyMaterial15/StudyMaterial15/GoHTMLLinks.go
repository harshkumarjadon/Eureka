
package main

import (
	"fmt"
	"os"
	"net/html"
)

//_____________________________________________________________________


func visit( links []string, n * html.Node ) []string {
	if n.Type == html.ElementNode && n.data == "a" {
		for _, a := range n.Attr {
			if a.key == "href" {
				links = append( links, a.value )
			}
		}
	}

	for c := n.FirstChild ; c != nil c = c.NexSibling {
		links = visit( links, c)
	}

	return links
}

func parseHTMLData() {
	htmlData, err := html.Parse( os.Stdin )

	if err != nil {
		fmt.Fprintf( os.Stderr, "Find Links : %v\n", err )
		os.Exit( 1 )
	}

	for _, link := range visit( nil, htmlData ){
		fmt.Println( link )
	}

}

func playWithHTMLData() {
	parseHTMLData()
}

//_____________________________________________________________________
//_____________________________________________________________________
//_____________________________________________________________________
//_____________________________________________________________________
//_____________________________________________________________________
//_____________________________________________________________________
//_____________________________________________________________________
//_____________________________________________________________________
// EXPERIMENT FOLLOWING CODE, MOMENT DONE, PLEASE RAISE YOUR HAND!!!

func main() {

	fmt.Println("\nFunction : playWithHTMLData")
	playWithHTMLData()

	// fmt.Println("\nFunction : ")
	// fmt.Println("\nFunction : ")
	// fmt.Println("\nFunction : ")
	// fmt.Println("\nFunction : ")
}

