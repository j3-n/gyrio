package util

func Sum(in []int) int {
	r := 0
	for _, i := range in {
		r += i
	}
	return r
}
