package spiral

/**
 * @input A : 2D integer array
 *
 * @Output Integer array.
 */
func spiralOrder(A [][]int) []int {
	n := len(A)
	m := len(A[0])
	resL := n * m
	res := make([]int, resL)
	k := 0
	bl := 0
	bt := 0
	br := m - 1
	bb := n - 1

	for k < resL {
		i := bt
		j := bl
		for ; k < resL && j <= br; j++ {
			res[k] = A[i][j]
			k++
		}
		i = bt + 1
		j = br
		for ; k < resL && i <= bb; i++ {
			res[k] = A[i][j]
			k++
		}
		i = bb
		j = br - 1
		for ; k < resL && j >= bl; j-- {
			res[k] = A[i][j]
			k++
		}
		i = bb - 1
		j = bl
		for ; k < resL && i >= bt+1; i-- {
			res[k] = A[i][j]
			k++
		}
		bt++
		bl++
		bb--
		br--
	}

	return res
}
