package search

// search --
// hay - a sorted ASC slice of int
// needle - an element to find
// l - low index (>=0) to border part of array where to search
// h - high index  (< en(hay)to border part of array where to search
//
// return an index of an element || -1 if not found
func search(hay []int, needle int) int {
	l := 0
	h := len(hay) - 1
	for l <= h {
		m := l + (h-l)/2
		if hay[m] == needle {
			return m
		}
		if hay[m] < needle {
			l = m + 1
		} else {
			h = m - 1
		}
	}
	return -1
}
