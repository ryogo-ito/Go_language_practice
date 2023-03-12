package section3

import "fmt"

func New() {
	s := make([]int, 0)
	fmt.Printf("%T\n", s)

	var p *int = new(int)
	fmt.Printf("%T\n", p)

	//var p2 *int
	//fmt.Println(p2)
}
