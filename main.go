package main

import (
	"fmt"
	"strings"
)

func main() {
	// learnVariable()
	// learnDataType()
	// learnConstantDanOperator()
	learnArray()
}
func learnVariable() {
	name := "thby"
	age := 20
	fmt.Println("ini adalah nana saya ==>", name)
	fmt.Println("ini adalah usia saya ==>", age)
	fmt.Printf("%T , %T \n", name, age)

	fmt.Println("\n", strings.Repeat("=", 40), "\n")

	student, lecture, grade := "luiz", "thby", 20
	fmt.Println(student, lecture, grade)

	var first, second, third int
	first = 10
	second = 20
	third = 30

	fmt.Println(first, second, third)

	// Underscore variable

	var firstVariable string
	msg, value, resault := "welcome", 17, 25

	_, _, _, _ = firstVariable, msg, value, resault

}
func learnDataType() {
	first := 89
	second := -17
	fmt.Printf("tipe data first : %T \n", first)
	fmt.Printf("tipe data second : %T \n", second)

	var decimal float32 = 1.64
	fmt.Printf("Float32 : %T \n", decimal)
	fmt.Printf("decimal : %f \n", decimal)
	fmt.Printf("decimal : %.3f \n", decimal)

	condtion := true

	fmt.Printf("Hasil : %t\n", condtion)

}
func learnConstantDanOperator() {
	const fulname string = "Applikasi Papua UMKM"

	fmt.Printf("Hasil Aplikasi : %s \n", fulname)

	jumlah := 23 + 5*2
	fmt.Println("Haslnya :", jumlah)

	perbandinagn := 4 < 2
	bools := "wanus" == "thby"
	cari := 3 != 4
	fmt.Printf("%v , %v , %v \n", perbandinagn, bools, cari)

	wrong := false
	right := true

	fmt.Println(wrong && right)
	fmt.Println(wrong || right)
	fmt.Println(wrong != right)
}
func learnArray() {
	arr := [4]string{"wanus", "thby", "luiz", "arthur"}
	for _, v := range arr {
		fmt.Println(v)
	}

	fruits := [3]string{}
	fruits[0] = "manggo"
	fruits[1] = "banana"
	fruits[2] = "papaya"
	fmt.Println(fruits)

}

// pages 37
