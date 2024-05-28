package util

// Sum calculates the sum of all elements in the input integer slice.
func Sum(in []int) int {
	r := 0
	for _, i := range in {
		r += i
	}
	return r
}

// Modulo implementation that always gives positive answers.
func Mod(a, b int) int {
	return (a%b + b) % b
}

// WrapString inserts newlines into the given string to keep it at the given maximum width.
// Negative or zero widths are not supported.
func WrapString(in string, width int) string {
	if width < 1 {
		return ""
	}

	res := ""
	for i, c := range in {
		if i%width == 0 && i > 0 {
			res += "\n"
		}
		res += string(c)
	}

	return res
}
