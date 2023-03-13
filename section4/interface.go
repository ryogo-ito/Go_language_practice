package section4

import "fmt"

type Human interface {
	Say() string
}

type Person struct {
	Name string
}

func DriveCar(human Human) {
	if human.Say() == "Mr.Mike" {
		fmt.Println("run")
	} else {
		fmt.Println("Get out")
	}
}

func (p *Person) Say() string {
	p.Name = "Mr." + p.Name
	return p.Name
}
