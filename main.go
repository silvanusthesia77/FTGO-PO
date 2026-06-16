package main

import "fmt"

func main() {
	// dataType()
	// array()
	// learnSlice()
	latih()
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

func learnSlice() {
	fruits := make([]string, 4)
	fruits[0] = "Thobiaz"
	fruits[1] = "Reza"
	fruits[2] = "Marvel"
	fruits[3] = "David"

	for _, v := range fruits {
		fmt.Println(v)
	}

	// append

	schools1 := []string{"Smansa", "SMA YPK", "SMA Agustinus"}
	schools2 := []string{"SDN 15", "SDN 14", "SD YPPK"}
	schools2 = append(schools2, schools1...)
	fmt.Println(schools2)

	fmt.Println("Masuk pada Copy Func :")

	// slice copy function
	member1 := []string{"Golix", "William", "Syukri"}
	member2 := []string{"Imal", "Clif", "Marcel"}
	mm := copy(member1, member2)

	fmt.Println("Member 1 -->", member1)
	fmt.Println("Member 2 -->", member2)
	fmt.Println("Hasil Coppy -->", mm)

	// Slice (slicing)

	var car = []string{"Toyota", "Honda", "Suzuki", "Jave", "python", "vscode"}
	var car1 = car[2:]
	fmt.Println(car1)

	cars := []string{"abu", "ikan", "nasi", "sayur"}
	cars = append(cars[:1], "Thobbyy")
	carss := cars[1:]
	fmt.Println(carss)
	fmt.Println()

	// Slice (backing array)

	fruits1 := []string{"Jambu", "Mangga", "Strewberry", "pineaple"}

	fruits2 := fruits1[:2]
	fruits1[0] = "Encap"
	fmt.Println("Hasil Dari slicing :", fruits2)
	fmt.Println("hasil dari Perubahan :", fruits1)

	// Slice (cap function)
	students := []string{"Wanus", "Thobiaz", "reza", "Iman", "Marvel"}

	fmt.Println("Hasil Cap :", cap(students))
	fmt.Println("Hasil len :", len(students))

	// Slice (creating a new backing array)

	trucks := []string{"Suzuki", "Honda", "Matic", "Green Car"}
	trucks[2] = "Yamaha"

	fmt.Println(trucks)

	for _, v := range trucks {
		fmt.Println(v)
	}

	newTrucks := []string{}
	newTrucks = append(newTrucks, trucks[2:]...)
	fmt.Println(newTrucks)

	// latihan
	i := 21
	fmt.Printf("%v\n", i)
	fmt.Printf("%T \n", i)
}

func latih(){
	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print(j," ")
		}

		fmt.Println()
	}
}
