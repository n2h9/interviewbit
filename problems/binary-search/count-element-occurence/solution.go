package solution

func findCount(A []int, B int) int {
	n := len(A)
	if n <= 0 {
		return 0
	}
	if n == 1 {
		if A[0] == B {
			return 1
		}
		return 0
	}
	first := searchFirstOrLast(A, B, true)
	if first == -1 {
		return 0
	}
	return searchFirstOrLast(A, B, false) - first + 1
}

// searchFirstOrLast --
// hay - a sorted ASC slice of int
// needle - an element to find
// l - low index (>=0) to border part of array where to search
// h - high index  (< en(hay)to border part of array where to search
// first - to find first index if true, last index if false
//
// return an index of an element || -1 if not found
// search the first
func searchFirstOrLast(hay []int, needle int, first bool) int {
	l := 0
	h := len(hay) - 1
	index := -1
	for l <= h {
		m := l + (h-l)/2
		if hay[m] == needle {
			index = m
			if first {
				h = m - 1
			} else {
				l = m + 1
			}
			continue
		}
		if hay[m] < needle {
			l = m + 1
		} else {
			h = m - 1
		}
	}
	return index
}
