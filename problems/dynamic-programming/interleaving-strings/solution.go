package solution

func isInterleave(A string, B string, C string) int {
	if len(C) != (len(A) + len(B)) {
		return 0
	}

	c := &container{
		stringA: A,
		stringB: B,
		stringC: C,
	}

	if isInterleaveYellow(c, 0, 0, 0) {
		return 1
	}

	return 0
}

type container struct {
	stringA string
	stringB string
	stringC string
}

func isInterleaveYellow(c *container, Ai, Bi, Ci int) bool {
	if Ci >= len(c.stringC) {
		return true
	}
	couldUseA := Ai < len(c.stringA)
	couldUseB := Bi < len(c.stringB)
	if couldUseA && couldUseB && c.stringA[Ai] == c.stringB[Bi] {
		return c.stringA[Ai] == c.stringC[Ci] &&
			(isInterleaveYellow(c, Ai+1, Bi, Ci+1) ||
				isInterleaveYellow(c, Ai, Bi+1, Ci+1))
	}

	if couldUseA && c.stringC[Ci] == c.stringA[Ai] {
		return isInterleaveYellow(c, Ai+1, Bi, Ci+1)
	}

	if couldUseB && c.stringC[Ci] == c.stringB[Bi] {
		return isInterleaveYellow(c, Ai, Bi+1, Ci+1)
	}

	return false
}
