package main

import "fmt"

func main() {

	//Operatores Aritmeticos (), *, /, %, +, -
	var a = 4 + 2*3
	fmt.Println(a)

	//Operadores de asignacion: =, +=, -=, *=, /=, %=
	var b = 10
	b += 2
	fmt.Println(b)
	/*
		declaracion post-incremento y post-decremento: ++, --
		no son una expresion sino una declaracion
	*/
	var c = 3
	c--
	fmt.Println(c)

	//Operadores Comparacion: >, <, ==, !=, >=, <=
	fmt.Println(4 != 4)

	//Operadores Logicos &&, ||
	var edad = 30
	fmt.Println("edad:", edad >= 18 && edad <= 60)

	fmt.Println("Niño o Viejo:", edad < 18 || edad > 60)

	//Unario: !
	fmt.Println(!(4 == 4))
}
