package main

import (
	"fmt"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	// variable()
	dataTypes()
}
func variable() {
	name := "wanus"
	age := 22

	fmt.Println(name, age)
	fmt.Printf("%T %T \n", name, age)
	data1, data2, data3, data4 := "thoby", "parviddey", "junior", "solossa"
	fmt.Println(data1, data2, data3, data4)

	// Underscore variable

	var firstVariable string

	day1, day2, day3 := "day1", "day2", "day3"
	_, _, _, _ = firstVariable, day1, day2, day3
	fmt.Println(day1)
	fmt.Println("\n", strings.Repeat("=", 45), "\n")

	first, second := 1, "2"
	fmt.Printf("Tipe data first dan Second adalah : %T  %T\n", first, second)

	nama := "Ariel"
	usia := 22
	address := " Jalan Sudirman"
	fmt.Printf("hi , saya %s usia saya %d alamat saya %s", nama, usia, address)
}
func dataTypes() {
	first := 89
	second := -12
	fmt.Printf("tipe data first %T \n:", first)
	fmt.Printf("tipe data Second %T \n:", second)

	fmt.Println("\n", strings.Repeat("=", 40), "\n")
	var decimalNumbers float32 = 3.63
	fmt.Printf("decimal number : %f\n", decimalNumbers)
	fmt.Printf("Decimal : %.3f \n", decimalNumbers)
	spew.Dump(decimalNumbers)
}
func constantsOperators() {
	// 24
}
