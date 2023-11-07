package main

import "fmt"

func main() {

	var a uint8 = 100
	var b uint16 = 23000

	c := uint16(a) + b
	_ = 234

	var _ string = "test"
	fmt.Printf("Tipo:%T , Valor: %v", c, c)

}
