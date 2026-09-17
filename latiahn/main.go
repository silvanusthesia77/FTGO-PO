package main

import (
	"fmt"
)

type universityStudent struct {
	name  string
	age   int
	hobby string
}

func (u universityStudent) intoduce(msg string) string {
	return fmt.Sprintf("%s my name is %s , I am %d years old", msg, u.name, u.age)
}
func main() {
	var universtudent = universityStudent{
		name:  "Bibba",
		age:   20,
		hobby: "Football",
	}
	fmt.Println(universtudent.intoduce("Hiii,"))
}

// Method (Pointer method) pages 25
