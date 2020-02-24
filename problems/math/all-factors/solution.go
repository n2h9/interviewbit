package solution

import "math"

func allFactors(A int) []int {
	if A == 1 {
		return []int{1}
	}
	if A == 2 {
		return []int{1, 2}
	}
	if A == 3 {
		return []int{1, 3}
	}

	sqrtA := int(math.Sqrt(float64(A)))

	l := make([]int, 0, 20)
	r := make([]int, 0, 20)
	l = append(l, 1)
	r = append(r, A)
	i := 1
	j := 2
	var v int
	for ; j < r[i-1] && j <= sqrtA; j++ {
		if A%j != 0 {
			continue
		}

		v = A / j
		l = append(l, j)
		r = append(r, v)
		i++
	}

	l = l[:i]
	for k := i - 1; k >= 0; k-- {
		// prevent duplicate
		if r[k] != l[k] {
			l = append(l, r[k])
		}
	}

	return l
}
