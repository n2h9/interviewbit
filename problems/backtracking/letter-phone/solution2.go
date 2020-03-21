package solution

// letterCombinationsNoRecursion
func letterCombinationsNoRecursion(A string) []string {
	n := len(A)
	if n <= 0 {
		return []string{}
	}

	i := n - 1
	res := phoneButtons[A[i]]
	for i := n - 2; i >= 0; i-- {
		tmp := make(
			[]string,
			len(phoneButtons[A[i]])*len(res),
		)
		k := 0
		for _, l := range phoneButtons[A[i]] {
			for _, s := range res {
				tmp[k] = l + s
				k++
			}
		}
		res = tmp
	}
	return res
}
