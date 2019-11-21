// 宣告程式屬於哪個 package
package main

// 引入套件
import (
	"fmt"
	"math/rand"

	"log"
	"strconv"

	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"github.com/tyler-smith/go-bip39"
)

// GetSefeRandomSeed is prodiver safe seed function
func GetSefeRandomSeed() int64 {
	entropy, err := bip39.NewEntropy(128)
	if err != nil {
		log.Fatal(err)
	}

	mnemonic, _ := bip39.NewMnemonic(entropy)
	seed := bip39.NewSeed(mnemonic, "")

	fmt.Println("memonic: ", mnemonic)
	fmt.Println("seed: ", seed)

	wallet, err := hdwallet.NewFromSeed(seed)
	if err != nil {
		log.Fatal(err)
	}

	path := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/0")
	account, err := wallet.Derive(path, false)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(account.Address.Hex())
	str := account.Address.Hex()
	subs := str[2:9]
	res, err := strconv.ParseInt(subs, 16, 64)

	if err != nil {
		log.Fatal(err)
	}

	return res
}

// GetRandom is a call func
// retrun 0 ~ num-1.
func GetRandom(num int) int {
	s1 := rand.NewSource(GetSefeRandomSeed())
	r1 := rand.New(s1)
	x := r1.Intn(num)
	fmt.Print("Rand result: = ", x, ", ")

	return x
}

// RollADice :
// Bet on numbers, 1 to 6.
// Winning bet pays up to 5.88x
func RollADice(num [6]int, money float64) (int, float64) {
	winvalue := 0.0
	x := (GetRandom(6) + 1)

	betTable := []float64{0, 5.88, 2.94, 1.96, 1.47, 1.18, 0}
	total := num[0] + num[1] + num[2] + num[3] + num[4] + num[5]

	if total <= 0 || total >= 6 {
		return x, winvalue
	}

	if (1 == num[0] && x == 1) || (1 == num[1] && x == 2) || (1 == num[2] && x == 3) ||
		(1 == num[3] && x == 4) || (1 == num[4] && x == 5) || (1 == num[5] && x == 6) {
		winvalue = betTable[total] * float64(money)
	}

	fmt.Println("Win value: = ", winvalue)
	return x, winvalue
}

// TwoDice :
// Bet on sum, 2 to 12.
// Winning bet pays up to 35.64x
func TwoDice(num [12]int, money float64) (int, float64) {
	winvalue := 0.0
	x := (GetRandom(6) + GetRandom(6) + 2)

	chanceTable := []float64{2.78, 5.56, 8.33, 11.11, 13.89, 16.67, 13.89, 11.11, 8.33, 5.56, 2.78}
	total := 0
	for i := 0; i < 11; i++ {
		total += num[i]
	}

	if total <= 0 || total >= 11 {
		return x, winvalue
	}
	chance := float64(0)
	for i := 0; i < 11; i++ {
		if num[i] > 0 {
			chance += chanceTable[i]
		}
	}

	if num[x-2] == 1 {
		winvalue = 98 * float64(money) / chance
	}

	fmt.Println("Win value: = ", winvalue)
	return x, winvalue
}

// CoinFilp :
// 50% win chance.
// Winning bet pays 1.98x
func CoinFilp(num int, money float64) (int, float64) {
	winvalue := 0.0
	x := GetRandom(2)

	if num < 0 || num > 1 {
		return x, winvalue
	}

	if num == x {
		winvalue = 1.98 * money
	}

	fmt.Println("Win value: = ", winvalue)
	return x, winvalue
}

// Etheroll :
// Any win chance, 1% to 97%.
// Winning bet pays up to 98x
// Return (random, winvalue)
func Etheroll(num int, money float64) (int, float64) {
	winvalue := 0.0
	x := GetRandom(100) + 1

	if num <= 0 || num >= 98 {
		return x, winvalue
	}

	if x <= num {
		winvalue = (98 / (float64)(num)) * money
	}

	fmt.Println("Win value: = ", winvalue)
	return x, winvalue
}
