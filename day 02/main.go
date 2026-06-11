package main

import "fmt"

func main() {
	// learnCondition()
	learnSwitch()
}

func learnCondition() {
	currentYear := 2021

	if age := currentYear - 2019; age < 17 {
		fmt.Println("Belum Dapat Cetak KTP")
		return
	}
	fmt.Println("Dapat Mencetak KTP")
}
func learnSwitch(){
	// score := 6

	// switch score {
	// case 8:
	// 	fmt.Println("Perfect")
	// case 7:
	// 	fmt.Println("awesome")
	// default:
	// 	fmt.Printf("not bad")
	// }

	fmt.Println("Case yang lain:")
	// hasi := 8

	// switch  {
	// case hasi == 8:
	// 	fmt.Println("Done")
	// case (hasi < 9) && (hasi > 4):
	// 	fmt .Println("Lumayan")
	// default:
	// 	{
	// 		fmt.Println("Study harder")
	// 		fmt.Println("you need to learn")
	// 	}
		
	// }

		// fallthrouhg
	hasil := 9
	switch  {
	case (hasil <= 10) && (hasil >= 8):
		fmt.Println("Great Job")
		fallthrough
	case (hasil <= 7) && (hasil >= 6):
		fmt.Println("Good")
	default:
		{
			fmt.Println("Study Hard")
			fmt.Println("You Don't Have a Good Score")
		}
		
	}
}

// Nested Conditions
