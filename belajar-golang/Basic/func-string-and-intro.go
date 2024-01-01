package main

import (
	"fmt"
	"strings"
)

var p string = "gayung"

var a, b int = 1, 5 //declare variable berbentuk seperti array

var angka, yes, double, txt = 1, true, 5.3, "arnothz"

const ( //konstanta
	square         = "kotak"
	isToday  bool  = true
	numeric  uint8 = 1
	floatNum       = 2.2
)

const ( // x dan y memiliki value yang sama
	x = "konstanta"
	y
)

var ( //declare variable berbentuk seperti json
	g int = 1
	r     = 2
)

var fruits = [4]string{"apple", "grape", "banana", "melon"} //array

func main() {
	fmt.Println(angka, ",", yes, ",", double, ",", txt)
	fmt.Print("\n")

	/*var point int = 4
	if point == 10 { //if else
		fmt.Println("lulus dengan nilai sempurna")
	} else if point > 5 {
		fmt.Println("lulus")
	} else if point == 4 {
		fmt.Println("hampir lulus")
	} else {
		fmt.Printf("tidak lulus. nilai anda %d\n", point)
	}*/

	/*for i := 0; i < 5; i++ { //for perulangan
		fmt.Println("Angka", i)
	}*/

	/*for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print(j, " ")
		}

		fmt.Println()
	}*/

	var names = []string{"John", "Wick"} //func string
	printMessage("halo", names)
}

func printMessage(message string, arr []string) { //func string
	var nameString = strings.Join(arr, " ")
	fmt.Println(message, nameString)
}
