package solution

func sqrt(A int) int {
	if A <= 0 {
		return 0
	}
	if A <= 3 {
		return 1
	}
	if A < 9 {
		return 2
	}
	v := 0
	l := 0
	h := A / 2
	for l <= h {
		m := l + (h-l)/2
		tmp := m * m
		if tmp == A {
			v = m
			break
		}
		if tmp > A {
			h = m - 1
			continue
		}
		v = m
		l = m + 1
	}
	return v
}
