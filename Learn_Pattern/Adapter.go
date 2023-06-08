package main

import "fmt"

//Compute.GO
type computer interface {
	insertInSquarePort()
}

//Compute.GO

//Mac.go
type mac struct {
}

func (m *mac) insertInSquarePort() {
	fmt.Println("Insert square port into mac machine")
}

//Mac.go

//WindowAdapter.go
type windowsAdapter struct {
	windowMachine *windows
}

func (w *windowsAdapter) insertInSquarePort() {
	w.windowMachine.insertInCirclePort()
}

//WindowAdapter.go

//Window.go
type windows struct{}

func (w *windows) insertInCirclePort() {
	fmt.Println("Insert circle port into windows machine")
}

//Window.go

//Client.go
type client struct {
}

func (c *client) insertSquareUsbInComputer(com computer) {
	com.insertInSquarePort()
}

//Client.go

//Main function
func main() {
	client := &client{}
	mac := &mac{}
	client.insertSquareUsbInComputer(mac)
	windowsMachine := &windows{}
	windowsMachineAdapter := &windowsAdapter{
		windowMachine: windowsMachine,
	}
	client.insertSquareUsbInComputer(windowsMachineAdapter)
}
