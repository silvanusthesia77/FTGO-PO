package main

import "fmt"

func main() {
	// learnCondition()
	// learnSwitch()
	// learnNested()
	learnLooping()
	// latihan()
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

func learnNested(){
	nilai := 0

	if nilai > 7 {
		switch nilai{
		case 8:
			fmt.Println("Perfect")
		default:
			fmt.Println("Nice")
		}
	}else{
		if nilai == 5 {
		fmt.Println("Not Bad")	
		}else if nilai == 3  {
			fmt.Println("Keep Trying")
		}else{
			fmt.Println("You Can Do It")
			if nilai == 0 {
				fmt.Println("try Harder")
			}
		}

	}
}
 func latihan(){
	angka := 22

	if like:= angka - 21; like <5 {
		fmt.Println("Sealmat Datang")
	}else{
		fmt.Println("Lain waktu")
	}
 }

 func learnLooping(){
	for i := 0; i < 4; i++ {
		fmt.Println("Angka", i)
	}

	j := 0

	for j < 5{
		fmt.Println("Hasil AKhir :",j)
		j++
	}

	k := 0

	for{
		fmt.Println("Halo :",k)
		k++

		if k == 5{
		break
		}
	}
 }
// 15Loopings (third way of looping)
