package library

import (
	"fmt"
	"strings"
)

type student struct {
	name  string
	grade int
}

type class struct {
	class       string
	class_grade int
}

func exc() {
	var s1 student
	s1.name = "john wick tolol"
	s1.grade = 2

	var s2 = class{"Amajing class", 5}

	fmt.Println("name  :", s1.name)
	fmt.Println("grade :", s1.grade)
	fmt.Println("kelas :", s2.class, "\n", "kelas berapa :", s2.class_grade)

	//calling method
	s1.sayHello()

	var name = s1.getNameAt(3)
	fmt.Println("nama panggilan :", name)
}

func (x student) sayHello() { //method
	fmt.Println("halo", x.name)
	fmt.Println("level", x.grade)
}

func (s student) getNameAt(i int) string {
	return strings.Split(s.name, " ")[i-1]
}

func printnama() {
	fmt.Println("hellow")
}
