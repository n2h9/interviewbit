package solution

const SHAPE_BYTE = 'X'

/**
 * @input A : array of strings
 *
 * @Output Integer
 */
func black(A []string) int {
	n := len(A)
	if n <= 0 {
		return 0
	}
	m := len(A[0])
	res := 0

	visitedChars := make(map[int]bool, n*m)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if A[i][j] != SHAPE_BYTE {
				continue
			}
			if visitedChars[i*m+j] {
				continue
			}
			markAdjacentAsVisited(A, i, j, n, m, visitedChars)
			res++
		}
	}
	return res
}

func markAdjacentAsVisited(A []string, i, j, n, m int, visitedChars map[int]bool) {
	if i < 0 || i >= n {
		return
	}
	if j < 0 || j >= m {
		return
	}
	if A[i][j] != SHAPE_BYTE {
		return
	}
	if visitedChars[i*m+j] {
		return
	}
	visitedChars[i*m+j] = true
	markAdjacentAsVisited(A, i-1, j, n, m, visitedChars)
	markAdjacentAsVisited(A, i+1, j, n, m, visitedChars)
	markAdjacentAsVisited(A, i, j-1, n, m, visitedChars)
	markAdjacentAsVisited(A, i, j+1, n, m, visitedChars)
}
