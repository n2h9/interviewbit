package solution

import (
	"fmt"
)

func maximalRectangle(A [][]int) int {
	// max := 0
	n := len(A)
	m := len(A[0])
	// if n <= 0 || m <= 0 {
	// 	return 0
	// }
	acc := newAcc(n * m)
	// for i := 0; i < n; i++ {
	// 	for j := 0; j < m; j++ {
	// 		l1, l2 := maxS(A, i, j, n-1, m-1, acc)
	// 		s := l1 * l2
	// 		if s > max {
	// 			max = s
	// 		}
	// 	}
	// }

	l1, l2 := maxS(A, 0, 0, n-1, m-1, acc)

	return l1 * l2
}

// return (length to bottom, length to right)
// for the max possible square from (x, y) which is top left
// For example:
// 1 1 1 1
// 1 1 1 0
// 0 1 1 1
// result is (2, 3)
func maxS(A [][]int, x, y, bb, rb int, a *acc) (int, int) {
	// out of border
	if x > bb || y > rb {
		return 0, 0
	}

	if A[x][y] == 0 {
		return 0, 0
	}
	if v1, v2, ok := a.get(x, y, bb, rb); ok {
		return v1, v2
	}

	bz := x+1 > bb || A[x+1][y] == 0
	rz := y+1 > rb || A[x][y+1] == 0
	// both right and bottom are zero, return 1 as suare of one cell
	if rz && bz {
		a.set(x, y, bb, rb, 1, 1)
		return 1, 1
	}

	// one way to the right, with bottom border eq to x
	// 1 1
	// 0 nm
	if bz && !rz {
		_, rl := maxS(A, x, y+1, x, rb, a)
		a.set(x, y, bb, rb, 1, 1+rl)
		return 1, 1 + rl
	}

	// one way to the bottom, with right border eq to y
	// 1 0
	// 1 nm
	if !bz && rz {
		bl, _ := maxS(A, x+1, y, bb, y, a)
		a.set(x, y, bb, rb, 1+bl, 1)
		return 1 + bl, 1
	}

	// both next bottom and next right cells are equal to 1
	// calculate both rectable right and bottom

	// right rectanle length
	_, rrl := maxS(A, x, y+1, x, rb, a)
	// bottom rectanle length
	brl, _ := maxS(A, x+1, y, bb, y, a)

	// if next diagonal also equal to one, we could create a biggger square
	// 1 1 ?
	// 1 1 ?
	// ? ? ?
	if A[x][y] != 0 {
		// get bottom length and right length of diagonal rectangle
		bl, rl := maxS(A, x+1, y+1, bb, rb, a)
		// "cut" rectangle  lengths,  aacordint to left and top rectangles of dimension 1xn and nx1
		// see above
		if brl < bl {
			bl = brl
		}
		if rrl < rl {
			rl = rrl
		}

		// increase lengths to cover current cell A[x][y]
		bl++
		rl++
		s := bl * rl
		// return only if square is bigger than square of small left and top rectangles of dimension 1xn and nx1
		if s > (1+rrl) && s > (1+brl) {
			a.set(x, y, bb, rb, bl, rl)
			return bl, rl
		}
	}

	if rrl > brl {
		a.set(x, y, bb, rb, 1, 1+rrl)
		return 1, 1 + rrl
	}

	a.set(x, y, bb, rb, 1+brl, 1)
	return 1 + brl, 1
}

type acc struct {
	vals map[string][]int
}

func newAcc(l int) *acc {
	return &acc{
		vals: make(map[string][]int, l),
	}
}

func (a *acc) get(x, y, bb, rb int) (int, int, bool) {
	fmt.Println("get", x, y, bb, rb)
	if v, ok := a.vals[a.key(x, y, bb, rb)]; ok {
		return v[0], v[1], ok
	}
	return 0, 0, false
}

func (a *acc) set(x, y, bb, rb, value1, value2 int) {
	fmt.Println("set", x, y, bb, rb, " vals ", value1, value2)
	a.vals[a.key(x, y, bb, rb)] = []int{value1, value2}
}

func (a *acc) key(x, y, bb, rb int) string {
	return fmt.Sprintf(
		"%d:%d:%d:%d",
		x, y, bb, rb,
	)
}
