package section4

import "fmt"

type Person2 struct {
	Name string
	Age  int
}

// Stringメソッドを記述すればOK
// 出力が下記のformatに変わる
func (s Person2) String() string {
	return fmt.Sprintf("My Print %v.", s.Name)
}
