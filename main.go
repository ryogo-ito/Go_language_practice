package main

import "fmt"

func main() {
	var mike UserError = New("Mike", 200)

	fmt.Println(mike)
}
