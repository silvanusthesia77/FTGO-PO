package main

import (
	"fmt"
	"strings"
)

func main() {
	// learnVariable()
	// learnDataType()
	// learnConstantDanOperator()
	// learnArray()
	// learnSlice()
	learnCondition()
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
	for key, v := range fruits {
		fmt.Printf("Index : %v ==> Value : %v\n", key, v)
	}
	fmt.Println("\n", strings.Repeat("=", 43), "\n")

	student := [2][3]string{{"Kuda", "Dog", "Cow"}, {"snake", "fish", "crocodile"}}

	for _, v := range student {
		for _, students := range v {
			fmt.Println(students)
		}
	}

}
func learnSlice() {
	players := []string{"thby", "luiz", "rez", "Golx"}
	for _, v := range players {
		fmt.Println(v)
	}

	student := make([]string, 3)
	student[0] = "Anissa"
	student[1] = "Noval"
	student[2] = "firdy"
	student = append(student, "wanus", "thby", "reza")
	fmt.Println(student)

	fmt.Println("\n", strings.Repeat("=", 40), "\n")

	msg := []string{"manajement", "sistem informasi", "informatika", "akuntansi"}
	msg1 := []string{"Unsia", "ui", "itb", "its"}

	msg = append(msg, msg1...)
	fmt.Println(msg)
	copied := copy(msg, msg1)
	fmt.Println(copied)

	fmt.Println("\n", strings.Repeat("=", 40), "\n")

	// Slice (slicing)

	buah := []string{"Jambu", "maggo", "pisang", "salak", "rambuatan", "melon", "smangka"}

	buah1 := buah[2:]
	buah2 := buah[1:3]
	buah3 := buah[:3]

	fmt.Println(buah1)
	fmt.Println(buah2)
	fmt.Println(buah3)

	// Slice (combining slicing and append)

	buah = append(buah[2:], "Sagu")
	fmt.Println(buah)

	// Slice (backing array)

	lectures := []string{"fes", "halim", "luk", "nathan"}
	fmt.Println(lectures)
	lectures[1] = "thby"
	lecture1 := lectures[:3]
	fmt.Println(lecture1)

	fmt.Println("\n", strings.Repeat("=", 45), "\n")

	ojol := []string{"Rkl", "DRT", "MTH", "JDK", "LBH"}
	fmt.Println("hasil cap :", cap(ojol))
	fmt.Println("hasil leng :", len(ojol))
	ojol1 := ojol[2:]
	fmt.Println("Hasil Di slice :", cap(ojol1))
	fmt.Println("hasil slice leng :", len(ojol1))

	cars := []string{"Zuzuki", "Honda", "Yamaha", "Beat", "Mio", "Pajero"}
	newCars := []string{}
	fmt.Println("Jumlah Cars :", cars)
	newCars = append(newCars, cars[2:]...)
	fmt.Println("hasil akhir :", newCars)

}
func learnCondition() {
	currentYear := 2026
	age := currentYear - 2005

	if age < 17 {
		fmt.Println("Kamu Belum Boleh Masuk")
	} else {
		fmt.Println("Silahkan Masuk")
	}

	nilai := 2

	switch nilai {
	case 8:
		fmt.Println("Perfect")
	case 7:
		fmt.Println("Not Bad")
	default:
		{
			fmt.Println("Maaf belum Lulus")
			fmt.Println("Come back stronger")
		}
	}

	// Switch with relational operators

	hasil := 8

	switch {
	case hasil == 6:
		fmt.Println("Congrats")
	case (hasil < 10) && (hasil > 7):
		fmt.Println("Selamat Datang")
	default:
		{
			fmt.Println("Not Bad")
			fmt.Println("But Not This time")
		}
	}

}

// pages 9
