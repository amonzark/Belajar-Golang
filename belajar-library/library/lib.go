package library

import "fmt"

func SayHello(name string, angka int) {
	fmt.Println("hello")
	introduce(name)
	fmt.Println("Score saya : ", getInteger(angka))
	var s1 = student{"john", 3}
	fmt.Println("\nTemannya beranama : ", s1.Name, "\ngrade sekarang adalah :", s1.grade)
	fmt.Println(testing(5))
}

func introduce(name string) {
	fmt.Println("nama saya", name)
}

func getInteger(angka int) int {
	var a int = 5 + angka
	return a
}
