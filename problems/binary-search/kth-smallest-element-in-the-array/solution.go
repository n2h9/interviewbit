package solution

func kthsmallest(A []int, B int) int {
	if B <= 0 {
		return 0
	}
	l := len(A)
	if l <= 0 {
		return 0
	}

	min, max := 0, ((1 << 31) - 1)
	nMin := max
	var n, nExact, nCount int

	for min <= max {
		n = max - (max-min)/2
		nExact = 0
		nCount = 0
		for i := 0; i < l; i++ {
			if A[i] <= n {
				// count number of int which is <= current n value
				nCount++
				if A[i] > nExact {
					// save exact value which most close to current n
					nExact = A[i]
				}
			}
		}
		if nCount >= B {
			// so we have at least B elements less than or equal to n
			if nExact < nMin {
				nMin = nExact
			}
			// decrease n range
			max = n - 1
		} else {
			// not enought elements, increase n range
			min = n + 1
		}
	}

	return nMin
}
