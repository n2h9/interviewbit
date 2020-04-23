package solution

const modulo = 1000000000 + 7

func nchoc(A int, B []int) int {
	n := len(B)

	sum := 0
	for i := n/2 - 1; i >= 0; i-- {
		heapify(B, i, n)
	}

	for ; A > 0; A-- {
		sum = (sum + B[0]) % modulo
		B[0] >>= 1
		heapify(B, 0, n)
	}

	return sum
}

func heapify(a []int, rootIndex, lenght int) {
	largestIndex := rootIndex
	leftIndex := 2*rootIndex + 1
	rightIndex := 2*rootIndex + 2

	if leftIndex < lenght && a[leftIndex] > a[largestIndex] {
		largestIndex = leftIndex
	}

	if rightIndex < lenght && a[rightIndex] > a[largestIndex] {
		largestIndex = rightIndex
	}

	if largestIndex != rootIndex {
		swap(a, largestIndex, rootIndex)
		heapify(a, largestIndex, lenght)
	}
}

func swap(a []int, i, j int) {
	a[i], a[j] = a[j], a[i]
}
