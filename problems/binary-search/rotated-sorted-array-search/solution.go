package solution

func search(A []int, B int) int {
	n := len(A)
	if n <= 0 {
		return -1
	}
	if n == 1 && A[0] == B {
		return 0
	}

	minLeft := searchMinOrNeedle(A, 0, (n-1)/2, n/2, B)
	if A[minLeft] == B {
		return minLeft
	}
	minRight := searchMinOrNeedle(A, ((n-1)/2)+1, n-1, n/2, B)
	if A[minRight] == B {
		return minRight
	}
	min := minLeft
	if minRight < min {
		min = minRight
	}

	return bsearch(A, min, B)
}

// search min in rotated sorted array
// simultaneously try to find needle
// if needle is found return its index instead
func searchMinOrNeedle(hay []int, start, end, mprev, needle int) int {
	index := start
	l := start
	h := end
	if l == mprev {
		l++
	} else {
		if h == mprev {
			h--
		}
	}
	for l <= h {
		m := l + (h-l)/2
		if hay[m] == needle {
			index = m
			break
		}
		if hay[m] < hay[index] {
			index = m
		}
		if hay[m] < hay[mprev] {
			h = m - 1
		} else {
			l = m + 1
		}
		mprev = m
	}
	return index
}

func bsearch(hay []int, correctionIndex, needle int) int {
	n := len(hay)
	l := 0
	h := n - 1
	for l <= h {
		m := l + (h-l)/2
		ri := realIndex(n-1, correctionIndex, m)
		if hay[ri] == needle {
			return ri
		}
		if hay[ri] < needle {
			l = m + 1
		} else {
			h = m - 1
		}
	}
	return -1
}

func realIndex(last, correction, current int) int {
	des := correction + current
	if des > last {
		des -= last
	}
	return des
}
