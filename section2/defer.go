package main

import "fmt"

func main() {
	defer fmt.Println("main関数が終わったら実行")
	fmt.Println("run")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("end")
}
