package main

import "fmt"

func main() {
	learnCondition()
}

func learnCondition() {
	currentYear := 2021

	if age := currentYear - 2012; age < 17 {
		fmt.Println("Belum Dapat Cetak KTP")
		return
	}
	fmt.Println("Dapat Mencetak KTP")
}
