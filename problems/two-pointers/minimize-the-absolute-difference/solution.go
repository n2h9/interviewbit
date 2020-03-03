package solution

/**
 * @input A : Integer array
 * @input B : Integer array
 * @input C : Integer array
 *
 * @Output Integer
 */
func solve(a []int, b []int, c []int) int {
	an := len(a)
	bn := len(b)
	cn := len(c)

	if an <= 0 || bn <= 0 || cn <= 0 {
		return 0
	}

	x, y, z := an-1, bn-1, cn-1
	dif := absDiff(maxMin(a, b, c, x, y, z))
	for x > 0 || y > 0 || z > 0 {
		d := absDiff(maxMin(a, b, c, x, y, z))
		if d < dif {
			dif = d
		}
		for x > 0 {
			d := absDiff(maxMin(a, b, c, x-1, y, z))
			if d >= dif {
				break
			}
			dif = d
			if x > 0 {
				x--
			}
		}
		for y > 0 {
			d := absDiff(maxMin(a, b, c, x, y-1, z))
			if d >= dif {
				break
			}
			dif = d
			if y > 0 {
				y--
			}
		}
		for z > 0 {
			d := absDiff(maxMin(a, b, c, x, y, z-1))
			if d >= dif {
				break
			}
			dif = d
			if z > 0 {
				z--
			}
		}
		switch {
		case y <= 0 && z <= 0 && x > 0:
			x--
		case x <= 0 && z <= 0 && y > 0:
			y--
		case x <= 0 && y <= 0 && z > 0:
			z--
		case y <= 0 && a[x] >= c[z]:
			x--
		case y <= 0 && c[z] >= a[x]:
			z--
		case x <= 0 && b[y] >= c[z]:
			y--
		case x <= 0 && c[z] >= b[y]:
			z--
		case z <= 0 && a[x] >= b[y]:
			x--
		case z <= 0 && b[y] >= a[x]:
			y--
		case a[x] >= b[y] && a[x] >= c[z]:
			x--
		case b[y] >= a[x] && b[y] >= c[z]:
			y--
		case c[z] >= a[x] && c[z] >= b[y]:
			z--
		}
	}

	return dif
}

func maxMin(a, b, c []int, x, y, z int) (max int, min int) {
	min = a[x]
	max = a[x]
	if b[y] < min {
		min = b[y]
	} else {
		if b[y] > max {
			max = b[y]
		}
	}
	if c[z] < min {
		min = c[z]
	} else {
		if c[z] > max {
			max = c[z]
		}
	}

	return max, min
}

func absDiff(k, l int) int {
	v := k - l
	if v < 0 {
		v *= -1
	}
	return v
}
