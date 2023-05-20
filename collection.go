package main

import (
	"fmt"
	"strings"
)

//_____________________________________________________________________

func playWithArrays() {

	// Array Of string Type Data
	planets := [...]string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}

	planets1 := [8]string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}

	fmt.Println(planets)
	fmt.Println(planets1)

	livablePlanets := planets[2:4]
	icePlanets := planets[6:8]

	fmt.Println(livablePlanets, icePlanets)

	terrestrial := planets[:4]
	gasGiants := planets[4:6]
	iceGiants := planets[6:]
	fmt.Println(terrestrial, gasGiants, iceGiants)
}

//_____________________________________________________________________

func playWithMaps() {
	// Map Is Collection Of Tuples
	//		Each Tuple Is (Key, Value) Pair
	// Key Type Is string and Value Type int

	temperature := map[string]int{
		"Earth": 15,
		"Mars":  -65,
	}

	fmt.Println(temperature)
	temp := temperature["Earth"]
	fmt.Printf("On average the Earth is %vº C.\n", temp)

	temperature["Earth"] = 16
	temperature["Venus"] = 464
	fmt.Println(temperature)
}

//_____________________________________________________________________
//_____________________________________________________________________
//_____________________________________________________________________
//_____________________________________________________________________
//_____________________________________________________________________
//_____________________________________________________________________

func basenameAgain(s string) string {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:len(s)]
	}
	return s
}

func basename(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:] // Taking String Slice
			break
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:len(s)] // Taking String Slice
			break
		}
	}
	return s
}

func playWithBaseName() {
	fmt.Println(basename("/media/WorkArea/Trainings.Running/Coforge/Progress"))
	fmt.Println(basename("/media/WorkArea/Trainings.Running/Coforge/Progress/GoTypes.go"))
	fmt.Println(basenameAgain("/media/WorkArea/Trainings.Running/Coforge/Progress"))
	fmt.Println(basenameAgain("/media/WorkArea/Trainings.Running/Coforge/Progress/GoTypes.go"))
}

func main() {
	playWithBaseName()
	// fmt.Println("\nFunction :playWithArrays")
	// playWithArrays()

	// fmt.Println("\nFunction : playWithMaps")
	// playWithMaps()

	// fmt.Println(baseName("/media/workArea/Trainings.Running/Coforge/Progress"))
	// fmt.Println(baseName("/media/workArea/Trainings.Running/Coforge/Progress/GoType.go"))

	// fmt.Println(baseNameAgain("/media/workArea/Trainings.Running/Coforge/Progress"))
	// fmt.Println(baseNameAgain("/media/workArea/Trainings.Running/Coforge/Progress"))

	// fmt.Println("\nFunction : ")
	// fmt.Println("\nFunction : ")
	// fmt.Println("\nFunction : ")
	// fmt.Println("\nFunction : ")
	// fmt.Println("\nFunction : ")
	// fmt.Println("\nFunction : ")
	// fmt.Println("\nFunction : ")
	// fmt.Println("\nFunction : ")
	// fmt.Println("\nFunction : ")
	// fmt.Println("\nFunction : ")
	// fmt.Println("\nFunction : ")
	// fmt.Println("\nFunction : ")
	// fmt.Println("\nFunction : ")
}
