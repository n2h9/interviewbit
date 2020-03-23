package solution

import (
	"strconv"
)

func getPermutation(A int, B int) string {
	if A <= 0 || B <= 0 {
		return ""
	}

	if A > 12 {
		return strconv.FormatInt(1<<31-1, 10)
	}

	sequence := generateNumbers(A)
	nf := fact(A - 1)
	return find(sequence, B, nf)
}

// @input slice of int equals to sequence number to create permutation
// @input kth element of permutation to return
// @input (len(A)-1)! to avoid calculate factorial every time, just path it in function
func find(A []int, k, nf int) string {
	if len(A) <= 0 {
		return ""
	}
	if len(A) == 1 {
		strconv.FormatInt(int64(A[0]), 10)
	}
	n := len(A)
	// index of element in A on which starts kth permutation
	i := (search(n, nf, k) - 1)
	// --
	// a equals A without element A[i]
	a := append([]int{}, A[0:i]...)
	a = append(a, A[i+1:]...)
	// --
	k -= i * nf // decrease k to search futher
	if n > 1 {
		nf /= (n - 1) // (n-1)!
	} else {
		nf = 1
	}

	return strconv.FormatInt(int64(A[i]), 10) + find(a, k, nf)
}

func search(num, nf, k int) int {
	// find the number which starts the kth permutation using binary search
	l := 1
	r := num
	m := 1
	for l <= r {
		m = r - (r-l)/2
		a := 1
		b := m * nf
		if m > 1 {
			a = (b - nf + 1) // (m-1)*nf + 1
		}
		if k >= a && k <= b {
			break
		}
		if k < a {
			if r == m {
				r--
				continue
			}
			r = m
			continue
		}
		if l == m {
			l++
			continue
		}
		l = m
		continue
	}

	return m
}

func fact(n int) int {
	v := 1
	for i := 2; i <= n; i++ {
		v *= i
	}

	return v
}

func generateNumbers(n int) []int {
	v := make([]int, n)
	for i := 0; i < n; i++ {
		v[i] = i + 1
	}
	return v
}
