package main

import (
	"fmt"

	"Go_language_practice/section4"
)

func main() {
	if err := section4.MyFunc(); err != nil {
		fmt.Println(err)
	}
}
