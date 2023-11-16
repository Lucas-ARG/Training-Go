package main

import (
	"fmt"

	course "github.com/Lucas-ARG/Training-Go/tree/main/Encapsulasion/Curso"
)

func main() {
	Go := course.New("Go desde Cero", 0, false)
	Go.SetUserIDs([]uint{12, 56, 89})
	Go.SetClasses(map[uint]string{
		1: "Introducción",
		2: "Estructuras",
		3: "Maps",
	})

	Go.SetName("POO con Go")
	fmt.Println(Go.Name())
	Go.PrintClasses()
	fmt.Printf("%+v", Go)
}
