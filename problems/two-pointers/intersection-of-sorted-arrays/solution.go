package solution

func intersect(A []int, B []int) []int {
	nA := len(A)
	nB := len(B)
	if nA <= 0 || nB <= 0 {
		return []int{}
	}
	nMin := nA
	if nB < nMin {
		nMin = nB
	}

	res := make([]int, 0, nMin)
	iA, iB := 0, 0
	for iA < nA && iB < nB {
		for ; iA < nA && A[iA] < B[iB]; iA++ {
		}
		if iA == nA {
			break
		}
		for ; iB < nB && B[iB] < A[iA]; iB++ {
		}
		if iB == nB {
			break
		}
		if A[iA] == B[iB] {
			res = append(res, A[iA])
			iA++
			iB++
		}
	}

	return res
}
