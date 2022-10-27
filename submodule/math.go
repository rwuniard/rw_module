package submodule

func Multiply(a ...int) int {
	total := 1

	for _, v := range a {
		total *= v
	}
	return total
}
