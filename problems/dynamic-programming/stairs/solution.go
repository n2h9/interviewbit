package solution

func climbStairs(A int) int {
	if A <= 2 {
		return A
	}
	x, y := 2, 1
	for i := 3; i <= A; i++ {
		x, y = x+y, x
	}

	return x
}
