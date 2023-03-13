package section4

type MyInt int

func (i MyInt) Double() int {
	return int(i * 2)
}
