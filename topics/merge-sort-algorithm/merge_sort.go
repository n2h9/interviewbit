package algorithm

func sort(a []int, l, r int) {
	if l >= r {
		return
	}
	q := (l + r) / 2
	sort(a, l, q)
	sort(a, q+1, r)
	merge(a, l, q, r)
}

func merge(a []int, l, q, r int) {
	b := make([]int, q+1-l)
	copy(b, a[l:q+1])
	i := 0      // because i is an index of b slice which is starts from 0 always
	iq := q - l // this is a border for left side instead of q because of use of b slice
	j := q + 1
	k := l
	for k <= r {
		// -- left side
		if j > r {
			for ; i <= iq; i++ {
				a[k] = b[i]
				k++
			}
		} else {
			for ; i <= iq && b[i] <= a[j]; i++ {
				a[k] = b[i]
				k++
			}
		}
		// --

		// -- right side
		if i > iq {
			for ; j <= r; j++ {
				a[k] = a[j]
				k++
			}
		} else {
			for ; j <= r && a[j] < b[i]; j++ {
				a[k] = a[j]
				k++
			}
		}
		// --
	}
}
