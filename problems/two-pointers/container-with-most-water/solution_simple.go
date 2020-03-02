package solution

func maxAreaSimple(A []int) int {
	n := len(A)
	if n < 2 {
		return 0
	}

	s := 0
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			a := area(A, j, i)
			if a > s {
				s = a
			}
		}
	}

	return s
}
