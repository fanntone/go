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

	if !(bet == 1 || bet == 6) {
		return x, winvalue
	}
	
	name := []string {R1[x[0]-1], R2[x[1]-1], R2[x[2]-1], R3[x[3]-1], R4[x[4]-1]}
	calLine(bet, name, x)
	
	return x, winvalue
}

func calLine(bet int, name []string, x []int) int {
	count := 0;

	if x[0] == x[1] {
		if x[1] == x[2] {
			if x[2] == x[3] {
				if x[3] == x[4] {
					// 5R
					count += findsymbol(bet, name[0], 5)
				} else {
					// 4R
					count += findsymbol(bet, name[0], 4)
				}
			} else {
				// 3R
				count += findsymbol(bet, name[0], 3)
			}
		} else {
			// 2R
			count += findsymbol(bet, name[0], 2)
		}
	} 

	if x[1] == x[2] {
		if x[2] == x[3] {
			if x[3] == x[4] {
				// 4R
				count += findsymbol(bet, name[1], 4)
			} else {
				// 3R
				count += findsymbol(bet, name[1], 3)
			}
		} else {
			// 2R
			count += findsymbol(bet, name[1], 3)
		}
	}

	if x[2] == x[3] {
		if x[3] == x[4] {
			// 3R
			count += findsymbol(bet, name[2], 3)
		}  else {
			// 2R
			count += findsymbol(bet, name[2], 2)
		}
	}

	if x[3] == x[4] {
		// 2R
		count += findsymbol(bet, name[3], 2)
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
		} else if name == "BEGEMO" {
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
		} else if name == "LIONmax" {
			pay = LIONmax[nr-2]
		}
	} else if bet == 1 {
		if name == "BONUS" {
			pay = BONUSone[nr-2]
		} else if name == "BANANA" {
			pay = BANANAone[nr-2]
		} else if name == "BEGEMO" {
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
		} else if name == "LIONmax" {
			pay = LIONone[nr-2]
		}
	}

	return pay;
}
