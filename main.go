package main

import (
	"fmt"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	// variable()
	// dataTypes()
	// constantsOperators()
	// array()
	// slice()
	// conditions()
	loopings()
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
	const appName string = "Aplikasi Jual Beli Murah"
	fmt.Println(appName)

	number := (2 + 2) * 3
	nilai := 2 + 2*3
	fmt.Println("Nilai :", nilai)
	fmt.Println("Number :", number)
	var firstCondition bool = 3 < 4
	var secondCondition bool = 3 <= 4
	var thirdCondition bool = "thoby" == "tHoby"
	var fourthCondition bool = "thoby" != "tHoby"
	fmt.Println("First Condition :", firstCondition)
	fmt.Println("Second Condition :", secondCondition)
	fmt.Println("third Condition :", thirdCondition)
	fmt.Println("Fourth Condition :", fourthCondition)

	wrong := false
	wright := true

	wrongAndWright := wrong && wright
	wrongOrWright := wrong || wright
	reverse := wrong != wright

	spew.Dump("wrong and wrighr :", wrongAndWright)
	spew.Dump("wrong atau wright :", wrongOrWright)
	spew.Dump("Reverse :", reverse)
}
func array() {
	names := [3]string{}
	names[0] = "thoby"
	names[1] = "parviddey"
	names[2] = " junior"
	for _, v := range names {
		fmt.Println("Nama :", v)
	}

	number := [4]int{20, 30, 40, 50}
	fmt.Println(number)
	var angka [3]int
	angka = [3]int{12, 13, 16}

	fmt.Println(angka)

	fruits := [3]string{}
	fruits[0] = "Manggo"
	fruits[1] = "Banana"
	fruits[2] = "Strawberry"

	for index, v := range fruits {
		fmt.Println(index, v)
	}

	fmt.Println("\n", strings.Repeat("=", 45), "\n")
	for i := 0; i < len(fruits); i++ {
		fmt.Printf("index : %d ,  value : %s \n", i, fruits)
	}

	numbers := [2][3]int{{1, 2, 3}, {22, 33, 44}}
	for _, v := range numbers {
		for _, arry := range v {
			fmt.Printf("index : %d \n", arry)
		}
	}
	// Array (modify element through index) 36
}
func slice() {
	fruits := []string{"Apple", "Papaya", "Grapes", "Banana", "Durian"}
	_ = fruits
	fmt.Println(fruits)
	for _, v := range fruits {
		fmt.Println(v)
	}
	// Slice (make function)

	fruit := make([]string, 3)
	fruit[0] = "Sagu"
	fruit[1] = "Avocado"
	fruit[2] = "Drink"
	fruit = append(fruit, "Sayur - Mayur")
	fmt.Printf("%#v\n", fruit)
	student := make([]string, 2)
	_ = student
	student[0] = "Biba"
	student[1] = "thobbiii"
	student = append(student, "thobby", "junior")
	fmt.Println(student)
	for _, v := range student {
		fmt.Println(v)
	}
	// Slice (append function with ellipsis)
	fruit1 := []string{"Jambu", "Salak", "Nangka"}
	frui2 := []string{"Dragon Fruit", "Star Fruit"}
	fruit1 = append(fruit1, frui2...)
	fmt.Println(fruit1)

	// Slice (copy function)

	buah := copy(fruit1, frui2)
	fmt.Println(buah)

	// Slice (slicing)

	students := []string{"thoby", "parviddey", "junior", "luiz", "arthur", "kandmy"}
	fmt.Println(students[:2])
	fmt.Println(students[2:])
	fmt.Println(students[1:3])
	fmt.Println(students[:])

	// Slice (combining slicing and append)

	universities := []string{"Unsia", "Binus", "UGM", "UNY", "UKDW"}
	universities = append(universities[1:], "Janabadra")
	fmt.Println(universities)

	// Slice (backing array)

	players := []string{"Thoby", "Parviddey", "Junior", "Luiz", "Arthur"}
	players1 := players[1:]
	players1[0] = "Biiiibbbaaa"

	fmt.Println("Players :", players)
	fmt.Println("Players1 :", players1)

	// Slice (cap function) 52

	names := []string{"wanus", "marvel", "marcel", "imal"}
	fmt.Println(len(names))
	fmt.Println(cap(names))

	fmt.Println("\n", strings.Repeat("=", 35), "\n")

	names2 := names[0:3]
	fmt.Println(len(names2))
	fmt.Println(cap(names2))
	fmt.Println("\n", strings.Repeat("=", 35), "\n")

	names3 := names[1:]
	fmt.Println(len(names3))
	fmt.Println(cap(names3))

	// Slice (creating a new backing array)

	cars := []string{"Nizan", "Honda", "Yamaha", "Ertiga"}
	newCras := []string{}

	newCras = append(newCras, cars[1:]...)

	fmt.Println("Cras :", cars)
	fmt.Println("New Cars :", newCras)
}

// Challenge

func conditions() {
	currentYear := 2026

	if age := currentYear - 2015; age < 16 {
		fmt.Println("Kamu Belum Dapat Cetak KTP")
	} else {
		fmt.Println("Kamu Dapat Cetak KTP")
	}

	number := 5

	switch {
	case number == 8:
		fmt.Println("Perfect")
	case number == 6:
		fmt.Println("Not Bad")
	case number == 5:
		fmt.Println("Come Back Stronger")
		fallthrough
	// Switch (fallthrough keyword)
	default:
		{
			fmt.Println("Semangat Masih Ada tahun Depan")
			fmt.Println("Just Believe in second Chance")
		}
	}
	// Nested Conditions
	score := 0

	if score > 8 {
		switch score {
		case 9:
			fmt.Println("Perfect")
		default:
			fmt.Println("Good Job")
		}
	} else {
		if score == 6 {
			fmt.Println("Good")
		} else if score == 4 {
			fmt.Println("Kok Bisa 4")
		} else {
			fmt.Println("Mantap")
			if score == 0 {
				fmt.Println("goblok Loh")
			}
		}
	}

}
func loopings() {
	for i := 0; i < 5; i++ {
		fmt.Println("Angka :", i)
	}

	// Loopings (second way of looping)

	a := 0

	for {
		fmt.Println("Hasil :", a)
		a++
		if a == 5 {
			break
		}
	}
}

// Loopings (break and continue keywords) 17
