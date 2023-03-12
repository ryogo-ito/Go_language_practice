package main

import (
	"fmt"

	"Go_language_practice/section4"
)

func main() {
	v := section4.New(3, 4, 5)

	fmt.Printf("%p\n", v)

	v.Scale3D(10)
	fmt.Println(v.Area3D())
}
