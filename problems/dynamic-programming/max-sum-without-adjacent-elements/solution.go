package solution

func adjacent(A [][]int) int {
	if len(A) != 2 {
		return 0
	}
	accumulator := make([]int, len(A[0]))
	return sum(A, accumulator, 0)
}

func sum(A [][]int, acc []int, i int) int {
	if i >= len(A[0]) {
		return 0
	}
	if acc[i] > 0 {
		return acc[i]
	}
	acc[i] = max(
		max(A[0][i], A[1][i])+sum(A, acc, i+2),
		sum(A, acc, i+1),
	)

	return acc[i]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
