package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	var devis Person
	devis.name = "wnaus"
	devis.age = 22
	fmt.Println(devis.name, devis.age)

	var person1 struct {
		person Person
		devisi string
		hobby  string
		class  int
	}
	person1.person.name = "Bibbaa"
	person1.person.age = 20
	person1.devisi = "Player"
	person1.hobby = "Football"
	person1.class = 30

	fmt.Println(person1.person.name, person1.person.age, person1.devisi, person1.hobby, person1.class)
}
