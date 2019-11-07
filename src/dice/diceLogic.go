// 宣告程式屬於哪個 package
package main

// 引入套件
import (
    "fmt"
    "math/rand"
    "time"
)

// GetRandom is a call func
// retrun 0 ~ num-1.
func GetRandom(num int) int {
    s1 := rand.NewSource(time.Now().UnixNano())
    r1 := rand.New(s1)
    x := r1.Intn(num)
    fmt.Print("Rand result: = ", x, ", " )

    return x
}

// RollADice : 
// Bet on numbers, 1 to 6.
// Winning bet pays up to 5.88x
func RollADice( num1 int,
                num2 int,
                num3 int,
                num4 int,
                num5 int,
                num6 int,
                money float64) float64 {
    winvalue := 0.0
    x := (GetRandom(6) + 1)

    betTable := []float64{ 0, 5.88, 2.94, 1.96, 1.47, 1.18, 0} 
    total := num1 + num2 + num3 + num4 + num5 + num6

    if (total <= 0 || total >= 6) {
        return winvalue;
    }
    
    if ((1 == num1 && x == 1) || (1 == num2 && x == 2) || (1 == num3 && x == 3) ||
        (1 == num4 && x == 4) || (1 == num5 && x == 5) || (1 == num6 && x == 6)) {
        winvalue = betTable[total] * float64(money)
    }

    fmt.Println("Win value: = ", winvalue)
    return winvalue
}

// TwoDice :
// Bet on sum, 2 to 12.
// Winning bet pays up to 35.64x
func TwoDice(num [12]int, money float64) float64 {
    winvalue := 0.0
    x := (GetRandom(11) + 2)

    betTable := []float64{ 0, 35.64, 11.88, 5.94, 3.56, 2.38, 1.70, 1.37, 1.19, 1.08, 1.02, 0 }
    total := 0;
    for i := 0; i < len(betTable); i++ {
        total += num[i]
    }

    for j := 0; j < len(betTable); j++ {
        if (1 == num[j] && x == (j + 1)) {
            winvalue = betTable[total] * float64(money)
            break;
        }
    }

    fmt.Println("Win value: = ", winvalue)
    return winvalue
}

// CoinFilp :
// 50% win chance.
// Winning bet pays 1.98x
func CoinFilp(num int, money float64) float64 {
    winvalue := 0.0
    x := GetRandom(2)

    if (num < 0 || num > 1) {
        return winvalue;
    }

    if (num == x) {
        winvalue = 1.98 * money
    }

    fmt.Println("Win value: = ", winvalue)
    return winvalue
}

// Etheroll :
// Any win chance, 1% to 97%.
// Winning bet pays up to 98x
func Etheroll(num int, money float64) (int,float64) {
    winvalue := 0.0
    x := GetRandom(100) + 1

    if (num <= 0 || num >= 98) {
        return x,winvalue;
    }

    if (x <= num) {
        winvalue = (98 /(float64)(num)) * money
    }

    fmt.Println("Win value: = ", winvalue)
    return x,winvalue
}
