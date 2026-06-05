package main

import "fmt"

func main() {
	// dataType()
	array()
}

func dataType() {
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

	jumlah := 361
	jumlah++
	fmt.Println(jumlah)

	isCorect := 8 != 7
	fmt.Println(isCorect)

	wrong := false
	right := true

	wrongRight := wrong && right
	fmt.Printf("Wrong Dan Right (%t) \n", wrongRight)
}

func array() {
	number := [5]int{1, 2, 3, 4, 5}
	name := [4]string{"thoby", "enchap", "Bie", "Marvel"}
	for _, v := range name {
		fmt.Println(v)
	}
	fmt.Println(name)
	fmt.Println(number)

	fruits := [4]string{}
	fruits[0] = "Manggo"
	fruits[1] = "start Fruit"
	fruits[2] = "Banana"
	fruits[3] = "Mellon"

	for _, fruit := range fruits {
		fmt.Println(fruit)
	}

	for i := 0; i < len(fruits); i++ {
		fmt.Printf("Index : %v =  Value : %v \n", i, fruits[i])
	}

	// arry multi dimensi

	hari := [3][3]string{{"Wanus", "lukas", "enchap"}, {"drama", "lukas", "yohanis"}}
	for _, v := range hari {
		for _, jumlah := range v {
			fmt.Println(jumlah)
		}
	}
}

// Slice
