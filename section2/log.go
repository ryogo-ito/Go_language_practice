package main

//import (
//	"fmt"
//	"log"
//	"os"
//)
//
//func main() {
//	file, err := os.Open("section2/log.go")
//	defer file.Close()
//	if err != nil {
//		log.Fatal("Error!")
//	}
//
//	data := make([]byte, 100)
//	count, err := file.Read(data)
//	if err != nil {
//		log.Fatal("Error!")
//	}
//
//	fmt.Println(count, string(data))
//
//	if err = os.Chdir("test"); err != nil {
//		log.Fatal("Error!")
//	}
//
//}
