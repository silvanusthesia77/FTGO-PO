package main

import "fmt"

func main() {
	name := "Wanus"
	age := 23

	fmt.Printf("Nama %s , Usia %d \n", name, age)

	// Multiple variable declarations
	name1, name2, name3 := "wanus", "Lukas", "Marvel"
	var first, second, third int = 1, 44, 55
	devision, rank, hobby := "wanus", 22, "Sport"
	fmt.Println(name1, name2, name3)
	fmt.Println(first, second, third)
	fmt.Println(devision, rank, hobby)

	// Underscore variable
	var firstVariable string

	fruit, drink, seafood := "manggo", "juice", "fish"

	_, _, _, _ = firstVariable, fruit, drink, seafood
	fmt.Println(seafood, fruit)
}

// Data type
