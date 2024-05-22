package util

func Sum(in []int) int {
	r := 0
	for _, i := range in {
		r += i
	}
	return r
}

// Modulo implementation that always gives positive answers
func Mod(a, b int) int {
	return (a%b + b) % b
}
