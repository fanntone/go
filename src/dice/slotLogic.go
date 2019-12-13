package main

// 引入套件
import (
	"fmt"
)

// BetOneLine is a logic function
// bet one/max()
func BetOneLine(bet int) (result []int, money int) {
	//Reel
	R1 := []string{	"LION", "MONKEY", "COCODRILE", "BANANA", "KAKADU", "BANANA", "BEGEMOT", "COCKTAIL", "COCODRILE", "MAN",
					"LION", "BONUS", "KAKADU", "MONKEY", "COCODRILE", "MAN", "COCKTAIL", "MONKEY", "BEGEMOT", "LION",
					"BANANA", "BEGEMOT", "KAKADU", "BANANA", "MAN", "COCKTAIL", "LION", "BEGEMOT", "COCODRILE", "MONKEY",
					"KAKADU", "MAN"}

	R2 := []string{ "KAKADU", "LION", "MAN", "COCKTAIL", "MONKEY", "MAN", "COCKTAIL", "LION", "BEGEMOT", "COCKTAIL",
					"COCODRILE", "BANANA", "BEGEMOT", "LION", "MONKEY", "BANANA", "COCODRILE", "KAKADU", "COCKTAIL", "BANANA",
					"MAN", "BONUS", "LION", "MONKEY", "COCODRILE", "KAKADU", "BEGEMOT", "BANANA", "COCKTAIL", "MAN",
					"LION", "MONKEY"}
	
	R3 := []string{ "MONKEY", "COCODRILE", "KAKADU", "COCKTAIL", "BEGEMOT", "LION", "MAN", "BANANA", "KAKADU", "MAN",
					"LION", "BEGEMOT", "MONKEY", "COCKTAIL", "KAKADU", "BANANA", "LION", "COCODRILE", "BANANA", "MONKEY",
					"BEGEMOT", "KAKADU", "MAN", "COCKTAIL", "MONKEY", "LION", "COCODRILE", "BONUS", "MONKEY", "COCKTAIL",
					"MAN", "LION"}

	R4 := []string{ "MAN", "MONKEY", "BONUS", "MAN", "COCODRILE", "COCKTAIL", "MONKEY", "COCKTAIL", "LION", "COCKTAIL",
					"KAKADU", "BEGEMOT", "MAN", "BANANA", "COCODRILE", "BANANA", "KAKADU", "BEGEMOT", "COCODRILE", "MONKEY",
					"BANANA", "KAKADU", "COCKTAIL", "MONKEY", "BEGEMOT", "LION", "BEGEMOT", "MAN", "COCODRILE", "LION",
					"BEGEMOT", "KAKADU"}
	
	R5 := []string{ "COCODRILE", "COCKTAIL", "BEGEMOT", "COCODRILE", "LION", "KAKADU", "MONKEY", "COCODRILE", "BEGEMOT", "KAKADU",
					"MAN", "BANANA", "LION", "COCODRILE", "MAN", "BONUS", "KAKADU", "MONKEY", "COCODRILE", "MAN", 
					"COCKTAIL", "BEGEMOT", "LION", "KAKADU", "BANANA", "BEGEMOT", "MONKEY", "BANANA", "MAN", "LION",
					"COCKTAIL", "KAKADU"}

	winvalue := 0
	x := []int{0, 0, 0, 0, 0}
	for i:=0; i< 5; i++ {
		x[i] = (GetRandom(32) + 1)
		fmt.Println("x = ", x[i]);
	}

	fmt.Println("R1 :=", R1[x[0]-1]);
	fmt.Println("R2 :=", R2[x[1]-1]);
	fmt.Println("R3 :=", R3[x[2]-1]);
	fmt.Println("R4 :=", R4[x[3]-1]);
	fmt.Println("R5 :=", R5[x[4]-1]);

	if !(bet == 1 || bet == 5) {
		return x, winvalue
	}
	
	name := []string {R1[x[0]-1], R2[x[1]-1], R3[x[2]-1], R4[x[3]-1], R5[x[4]-1]}
	winvalue = calLine(bet, name)

	return x, winvalue
}

func calLine(bet int, name []string) int {
	count := 0;

	if name[0] == name[1] {
		if name[1] == name[2] {
			if name[2] == name[3] {
				if name[3] == name[4] {
					// 5R
					count += findsymbol(bet, name[0], 5)
					fmt.Println("5R: R1==R2==R3==R4==R5")
				} else {
					// 4R
					count += findsymbol(bet, name[0], 4)
					fmt.Println("4R: R1==R2==R3==R4")
				}
			} else {
				// 3R
				count += findsymbol(bet, name[0], 3)
				fmt.Println("3R: R1==R2==R3")
			}
		} else {
			// 2R
			count += findsymbol(bet, name[0], 2)
			fmt.Println("2R: R1==R2")
		}
	} else if name[1] == name[2] {
		if name[2] == name[3] {
			if name[3] == name[4] {
				// 4R
				count += findsymbol(bet, name[1], 4)
				fmt.Println("4R: R2==R3==R4==R5")
			} else {
				// 3R
				count += findsymbol(bet, name[1], 3)
				fmt.Println("3R: R2==R3==R4")
			}
		} else {
			// 2R
			count += findsymbol(bet, name[1], 3)
			fmt.Println("2R: R2==R3")
		}
	} else if name[2] == name[3] {
		if name[3] == name[4] {
			// 3R
			count += findsymbol(bet, name[2], 3)
			fmt.Println("3R: R3==R4==R5")
		}  else {
			// 2R
			count += findsymbol(bet, name[2], 2)
			fmt.Println("2R: R3==R4")
		}
	} else if name[3] == name[4] {
		// 2R
		count += findsymbol(bet, name[3], 2)
		fmt.Println("2R: R4==R5")
	}

	return count;
}

func findsymbol(bet int, name string, nr int) int {
	// symbol            	 2R   3R   4R   5R  
	BONUSmax 		:= []int{800, 1000, 1600, 2000}
	BANANAmax		:= []int{100, 200, 260, 300}
	BEGEMOTmax		:= []int{50, 100, 160, 200}
	COCODRILEmax	:= []int{50, 100, 160, 200}
	COCKTAILmax		:= []int{5, 100, 160, 200}
	KAKADUmax		:= []int{5, 75, 90, 100}
	MANmax			:= []int{5, 75, 90, 100}
	MONKEYmax		:= []int{2, 75, 90, 100}
	LIONmax			:= []int{2, 50, 40, 50}

	// symbol            	 2R   3R   4R   5R  
	BONUSone 		:= []int{50, 100, 100, 170}
	BANANAone		:= []int{10, 20, 20, 80}
	BEGEMOTone		:= []int{5, 10, 10, 20}
	COCODRILEone	:= []int{5, 10, 10, 20}
	COCKTAILone		:= []int{2, 10, 10, 15}
	KAKADUone		:= []int{2, 5, 5, 8}
	MANone			:= []int{2, 5, 5, 8}
	MONKEYone		:= []int{1, 5, 5, 8}
	LIONone			:= []int{1, 2, 2, 3}

	pay := 0;
	if bet == 5 {
		if name == "BONUS" {
			pay = BONUSmax[nr-2]
		} else if name == "BANANA" {
			pay = BANANAmax[nr-2]
		} else if name == "BEGEMOT" {
			pay = BEGEMOTmax[nr-2]
		} else if name == "COCODRILE" {
			pay = COCODRILEmax[nr-2]
		} else if name == "COCKTAIL" {
			pay = COCKTAILmax[nr-2]
		} else if name == "KAKADU" {
			pay = KAKADUmax[nr-2]
		} else if name == "MAN" {
			pay = MANmax[nr-2]
		} else if name == "MONKEY" {
			pay = MONKEYmax[nr-2]
		} else if name == "LION" {
			pay = LIONmax[nr-2]
		}
	} else if bet == 1 {
		if name == "BONUS" {
			pay = BONUSone[nr-2]
		} else if name == "BANANA" {
			pay = BANANAone[nr-2]
		} else if name == "BEGEMOT" {
			pay = BEGEMOTone[nr-2]
		} else if name == "COCODRILE" {
			pay = COCODRILEone[nr-2]
		} else if name == "COCKTAIL" {
			pay = COCKTAILone[nr-2]
		} else if name == "KAKADU" {
			pay = KAKADUone[nr-2]
		} else if name == "MAN" {
			pay = MANone[nr-2]
		} else if name == "MONKEY" {
			pay = MONKEYone[nr-2]
		} else if name == "LION" {
			pay = LIONone[nr-2]
		}
	}
	fmt.Println("nr = ", nr)
	fmt.Println("Pay = ", pay)
	return pay;
}
