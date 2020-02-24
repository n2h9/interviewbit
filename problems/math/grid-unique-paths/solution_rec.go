package solution

func uniquePathsRec(A int, B int) int {
	if A <= 0 && B <= 0 {
		return 0
	}
	if A == 1 && B == 1 {
		return 1
	}
	return pathNum(0, 0, A-1, B-1)
}

func pathNum(x, y, bX, bY int) int {
	if x == bX && y == bY {
		// we are one our finish cell
		return 1
	}
	if x > bX || y > bY {
		// we are out of border
		return 0
	}

	// move futher
	return pathNum(x+1, y, bX, bY) + pathNum(x, y+1, bX, bY)
}
