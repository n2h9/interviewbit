package algorithm

func sort(a []int) {
	n := len(a)
	if n < 2 {
		return
	}

	for i := n/2 - 1; i >= 0; i-- {
		heapify(a, i, n)
	}

	heapSort(a, 0, n)
}

func heapSort(a []int, rootIndex, length int) {
	for i := len(a) - 1; i > 0; i-- {
		swap(a, 0, i)
		heapify(a, 0, i)
	}
}

func swap(a []int, i, j int) {
	a[i], a[j] = a[j], a[i]
}

func heapify(a []int, rootIndex, length int) {
	largestIndex := rootIndex
	leftIndex := 2*rootIndex + 1
	rightIndex := 2*rootIndex + 2

	if leftIndex < length && a[leftIndex] > a[largestIndex] {
		largestIndex = leftIndex
	}
	if rightIndex < length && a[rightIndex] > a[largestIndex] {
		largestIndex = rightIndex
	}

	if largestIndex != rootIndex {
		swap(a, largestIndex, rootIndex)
		heapify(a, largestIndex, length)
	}
}
