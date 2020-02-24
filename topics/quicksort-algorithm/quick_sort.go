package algorithm

func sort(a []int, l, h int) {
	if l >= h {
		return
	}
	p := partition(a, l, h)
	sort(a, l, p-1)
	sort(a, p+1, h)
}

func partition(a []int, l, h int) int {
	// take h as a pivot
	fg := -1
	// i < h BUT NOT <= h because h is a pivot, avoid cmp pivot with a pivot
	for i := l; i < h; i++ {
		if fg == -1 {
			// a part when we haven't found an item gt a[h]
			if a[i] <= a[h] {
				continue // while items are lower just continue
			}
			fg = i // we find first item gt a[h], save its index
			continue
		}
		if a[i] > a[h] {
			continue // while items are higher, just continue
		}
		swap(a, i, fg) // found first item lower than a[h] after item(s) gt a[h]
		fg++           // go to next greter item
	}

	if fg == -1 {
		return h // all items are lower than a[h] so a[h] is on correct position
	}

	swap(a, h, fg) // move a[h] to fg position
	return fg      // and return it as pivot value
}

func swap(a []int, i1, i2 int) {
	tmp := a[i1]
	a[i1] = a[i2]
	a[i2] = tmp
}
