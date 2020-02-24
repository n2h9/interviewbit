package solution

import "sort"

func wave(A []int) []int {
	l := len(A)
	if l <= 1 {
		return A
	}
	sort.IntSlice(A).Sort()
	for i := 1; i < l; i += 2 {
		swap(A, i, i-1)
	}
	return A
}

// func sortOwn(a []int) {
// 	n := len(a)
// 	for i := 1; i < n; i++ {
// 		v := a[i]
// 		j := i - 1
// 		for ; j >= 0 && v < a[j]; j-- {
// 			a[j+1] = a[j]
// 		}
// 		a[j+1] = v
// 	}
// }

func swap(a []int, i1, i2 int) {
	tmp := a[i1]
	a[i1] = a[i2]
	a[i2] = tmp
}
