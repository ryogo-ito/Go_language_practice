package section3

func one(x int) {
	x = 1
}

func Pointer() int {
	var n int = 100

	one(n)
	result := n
	//result2 := &n

	return result
}
