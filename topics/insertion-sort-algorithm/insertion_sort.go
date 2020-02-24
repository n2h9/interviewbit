package algorithm

func sort(a []int) {
	n := len(a)
	for i := 1; i < n; i++ {
		v := a[i]
		j := i - 1
		for ; j >= 0 && v < a[j]; j-- {
			a[j+1] = a[j]
		}
		a[j+1] = v
	}
}
