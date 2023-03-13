package main

import (
	"fmt"

	"Go_language_practice/section4"
)

func main() {
	var mike section4.Human = &section4.Person{"Mike"}
	section4.DriveCar(mike)
	fmt.Println(mike)
}
