package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

var randomizer = rand.New(rand.NewSource(time.Now().Unix()))

func main() {
	//single return
	var randomValue int

	randomValue = randomWithRange(2, 10)
	fmt.Println("random number:", randomValue)
	randomValue = randomWithRange(2, 10)
	fmt.Println("random number:", randomValue)
	randomValue = randomWithRange(2, 10)
	fmt.Println("random number:", randomValue)

	//multiple return dengan func closure
	var getMulti = func(diamet float64) (float64, float64) { //multiple return
		// atau bisa "(area float64, circumference float64)""
		// hitung luas
		var area = math.Pi * math.Pow(diamet/2, 2)
		// hitung keliling
		var circumference = math.Pi * diamet

		// kembalikan 2 nilai
		return area, circumference
	}

	var diameter float64 = 15
	var area, circumference = getMulti(diameter)

	fmt.Printf("luas lingkaran\t\t: %.2f \n", area)
	fmt.Printf("keliling lingkaran\t: %.2f \n", circumference)
}

func randomWithRange(min, max int) int {
	var value = randomizer.Int()%(max-min+1) + min
	return value
}
