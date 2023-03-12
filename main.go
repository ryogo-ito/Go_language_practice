package main

import (
	"fmt"

	"Go_language_practice/section4"
)

func main() {
	v := section4.Vertex{3, 2}

	v.Scale(10)
	fmt.Println(v.Area())
}
