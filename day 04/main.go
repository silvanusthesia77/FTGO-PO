package main

import (
	"fmt"
	"strings"
)

func main() {
	var names =[]string{"Thoby"}
	var str = geet("Hiiiiiieee",names)
	fmt.Println(str)
}
func geet(msg string, names[]string)string{
	var funcStr = strings.Join(names, "")
	var result string = fmt.Sprintf("%s , %s ", msg, funcStr)
	return result
}

// Function (Returning multiple values)