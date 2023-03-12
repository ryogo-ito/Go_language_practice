package main

import (
	"fmt"

	"Go_language_practice/section3"
)

func main() {
	v := section3.Vertex{1, 2, "test"}

	fmt.Println("関数実行前", v)
	section3.ChangeVertex(v)
	fmt.Println("関数実行後", v)

	v2 := &section3.Vertex{1, 2, "test"}
	section3.ChangeVertex2(v2)
	fmt.Println(v2)
}
