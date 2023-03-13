package section4

import "fmt"

type UserNotFound struct {
	UserName string
}

func (e *UserNotFound) Error() string {
	return fmt.Sprintf("User not found: %v", e.UserName)
}

func MyFunc() error {
	ok := false

	// ...何かしらの処理が走ってるいる途中
	if ok {
		return nil
	}
	return &UserNotFound{"Mike"}
}
