package solution

import (
	"sort"
)

/**
* @input A : Integer array
*
* @Output 2D int array.
 */
func subsets(A []int) [][]int {
	n := len(A)
	if n <= 0 {
		return [][]int{[]int{}}
	}
	if n == 1 {
		return [][]int{[]int{}, A}
	}

	// sort first O(nlogn)
	sort.Ints(A)

	cache := make([][][]int, n)
	res := [][]int{[]int{}}
	for i := 0; i < n; i++ {
		res = append(res, findAll(A, i, cache)...)
	}

	return res
}

// all subsets is equal to A[i] combine with all subsets for A[i+1]
func findAll(A []int, index int, cache [][][]int) [][]int {
	n := len(A)
	if cache[index] != nil {
		return cache[index]
	}
	if (index + 1) == n {
		cache[index] = [][]int{[]int{A[index]}}
		return cache[index]
	}

	res := [][]int{[]int{A[index]}}
	for i := index + 1; i < n; i++ {
		subs := findAll(A, i, cache)
		for i := 0; i < len(subs); i++ {
			res = append(res, append([]int{A[index]}, subs[i]...))
		}
	}
	cache[index] = res
	return cache[index]
}
