package solution

/**
 * @input A : Integer array
 *
 * @Output 2D int array.
 */
func permute(A []int) [][]int {
	// fmt.Println(A)
	n := len(A)
	if n <= 0 {
		return [][]int{}
	}

	if n == 1 {
		return [][]int{[]int{A[0]}}
	}

	res := [][]int{}
	for i := 0; i < n; i++ {
		a := append([]int{}, A[0:i]...)
		a = append(a, A[i+1:]...)
		// all other permutration without A[i] element
		p := permute(a)
		for _, v := range p {
			res = append(res, append([]int{A[i]}, v...))
		}
	}

	return res
}
