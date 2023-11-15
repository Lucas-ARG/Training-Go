package main

import (
	"fmt"

	course "github.com/Lucas-ARG/Training-Go/tree/main/Encapsulasion/Curso"
)

func main() {
	Go := course.New("Go desde Cero", 0, false)
	Go.UserIDs = []uint{12, 56, 89}
	Go.Classes = map[uint]string{
		1: "Introducción",
		2: "Estructuras",
		3: "Maps",
	}
	// Go.PrintClasses()
	// Go.ChangePrice(67.12)
	fmt.Printf("%+v", Go)
}
