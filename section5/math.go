package section5

// Average returns the Average int
func Average(s []int) int {
	total := 0
	for _, i := range s {
		total += i
	}
	return total / len(s)
}
