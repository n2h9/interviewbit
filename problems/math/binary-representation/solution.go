package solution

func findDigitsInBinary(A int) string {
	if A == 0 {
		return "0"
	}
	if A == 1 {
		return "1"
	}
	b := ""
	z := "0"
	o := "1"

	for A >= 2 {
		r := A % 2
		A = A / 2
		if r == 0 {
			b = z + b
		} else {
			b = o + b
		}
	}
	b = o + b

	return b
}
