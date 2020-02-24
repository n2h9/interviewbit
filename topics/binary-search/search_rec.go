package search

// search --
// hay - a sorted ASC slice of int
// needle - an element to find
// l - low index (>=0) to border part of array where to search
// h - high index  (< en(hay)to border part of array where to search
//
// return an index of an element || -1 if not found
func searchRec(hay []int, needle, l, h int) int {
	if l > h || l < 0 || h >= len(hay) {
		// nothing found || or invalid index values
		return -1
	}
	m := l + (h-l)/2
	if hay[m] == needle {
		return m
	}
	if hay[m] < needle {
		return searchRec(hay, needle, m+1, h)
	}
	return searchRec(hay, needle, l, m-1)
}
