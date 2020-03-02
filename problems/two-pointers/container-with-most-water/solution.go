package solution

func maxArea(A []int) int {
	n := len(A)
	if n < 2 {
		return 0
	}

	i := n - 1
	j := 0
	s := 0
	for j < i {
		a := area(A, j, i)
		for j < i {
			x := area(A, j+1, i)
			if x >= a {
				a = x
				j++
				continue
			}
			break
		}
		for j < i {
			x := area(A, j, i-1)
			if x >= a {
				a = x
				i--
				continue
			}
			break
		}
		if a > s {
			s = a
		}
		if A[j] < A[i] {
			j++
		} else {
			i--
		}

	}

	return s
}

// restrictions:
// x,y valid indexes
// y >= x
func area(A []int, x, y int) int {
	if y <= x {
		return 0
	}
	m := A[y]
	if A[x] < m {
		m = A[x]
	}
	return (y - x) * m
}
