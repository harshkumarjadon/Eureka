package main

import (
	"fmt"
)

type Superpower interface {
	fly(interface{})
	saveWorld()
}

// class Spiderman {
type Spiderman struct {
	printValue string
}

func (r Spiderman) fly(s interface{}) {
	fmt.Println("Spiderman fly")
	fmt.Println(r.printValue)
}

func (r Spiderman) saveWorld() {
	fmt.Println("Spiderman Save world")
}

// class Superman {
type Superman struct {
	printValue string
}

func (r Superman) fly(s interface{}) {
	fmt.Println("Superman fly")
	fmt.Println(r.printValue)
}

func (r Superman) saveWorld() {
	fmt.Println("Superman Save world")
}

type HeMan struct {
	printValue string
}

func (r HeMan) fly(s interface{}) {
	fmt.Println("HeMan fly")
	fmt.Println(r.printValue)
}

func (r HeMan) saveWorld() {
	fmt.Println("HeMan Save world")
}

type WonderWoman struct {
	printValue string
}

func (r WonderWoman) fly(s interface{}) {
	fmt.Println("WonderWoman fly")
	fmt.Println(r.printValue)
}

func (r WonderWoman) saveWorld() {
	fmt.Println("WonderWoman Save world")
}

type HanumanJi struct {
	printValue string
}

func (r HanumanJi) fly(s interface{}) {
	fmt.Println("HanumanJi fly")
	fmt.Println(r.printValue)
}

func (r HanumanJi) saveWorld() {
	fmt.Println("HanumanJi Save world")
}

func HumanDemo[T Superpower](s T) {
	s.fly("")
	s.saveWorld()
}

func main() {
	sm := Spiderman{"Spider Man! Spider Man!"}
	HumanDemo[Spiderman](sm)

	super := Superman{"Super Man! Super Man!"}
	HumanDemo[Superman](super)

	heman := HeMan{"He Man! He Man!"}
	HumanDemo[HeMan](heman)

	wonderwoman := WonderWoman{"Wonder Man! Wonder Man!"}
	HumanDemo[WonderWoman](wonderwoman)

	hanumanji := HanumanJi{"Hanumanji Man! Hanumanji Man!"}
	HumanDemo[HanumanJi](hanumanji)

}
