package solution

// firstMissingPositiveB
// No extra space, use only input slice
// every element of an array will be processed max 4 times
// which is O(n)
func firstMissingPositiveB(A []int) int {
	n := len(A)
	if n <= 0 {
		return 1
	}

	// value to mark the number is present in array
	// any value <= 0 or > n
	marker := 0
	insteadOfMarker := -1
	// remove all marker value from array
	// n times
	for i := 0; i < n; i++ {
		if A[i] == marker {
			A[i] = insteadOfMarker
		}
	}

	var tmp, j int
	// 2n times max
	for i := 0; i < n; i++ {
		// a value is to small or to big => skip
		if A[i] <= 0 || A[i] > n {
			continue
		}
		j = A[i]
		for 0 < j && j <= n {
			// save the value to tmp variable
			tmp = A[j-1]
			// mark number j is present in array
			A[j-1] = marker
			// set j to value stored in array with index j-1
			j = tmp
		}
	}

	// return first non marked value
	// n times max
	for i := 0; i < n; i++ {
		if A[i] != marker {
			return i + 1
		}
	}

	// means that array contain all values from 1 to n
	// so return n + 1
	return n + 1
}
