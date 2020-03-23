package solution

/**
 * @input A : Integer
 *
 * @Output Integer
 */
func colorful(A int) int {
	// result -A is the same as for A
	if A < 0 {
		A = ^A + 1
	}
	if A < 10 {
		return 1
	}
	products := make(map[int]bool)
	digits := partition(A)
	for i := len(digits) - 1; i > -1; i-- {
		p := digits[i]
		if products[p] {
			return 0
		}
		products[p] = true
		for j := i - 1; j > -1; j-- {
			p *= digits[j]
			if products[p] {
				return 0
			}
			products[p] = true
		}
	}

	return 1
}

func partition(x int) []int {
	a := []int{}
	for x > 0 {
		r := x % 10
		a = append(a, r)
		x = (x - r) / 10
	}
	return a
}
