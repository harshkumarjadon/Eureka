package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"regexp"
)

// Create a Customer type
type Customer struct {
	Name string
	Age  int
}

// Implement a WriteJSON method that takes an io.Writer as the parameter.
// It marshals the customer struct to JSON, and if the marshal worked
// successfully, then calls the relevant io.Writer's Write() method.
func (c *Customer) WriteJSON(w io.Writer) error {
	js, err := json.Marshal(c)
	if err != nil {
		return err
	}

	_, err = w.Write(js)
	return err
}

var digitRegexp = regexp.MustCompile("[0-9]+")

func CopyDigits(filename string) []byte {
	b, _ := ioutil.ReadFile("./customer.json")
	b = digitRegexp.Find(b)
	c := make([]byte, len(b))
	copy(c, b)
	return c
}

func main() {
	// Initialize a customer struct.
	c := &Customer{Name: "Alice", Age: 21}

	// We can then call the WriteJSON method using a buffer...
	var buf bytes.Buffer
	err := c.WriteJSON(&buf)
	if err != nil {
		log.Fatal(err)
	}

	// Or using a file.
	f, err := os.Create("./customer.json")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	err = c.WriteJSON(f)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(CopyDigits(""))
}
