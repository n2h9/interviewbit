package solution

import (
	"sort"
)

/**
 * @input A : Integer array
 * @input B : Integer array
 *
 * @Output Integer
 */
func mice(A []int, B []int) int {
	if len(B) < len(A) {
		return -1
	}

	sort.Ints(A)
	sort.Ints(B)
	if len(B) == len(A) {
		max := 0
		for i := 0; i < len(A); i++ {
			d := A[i] - B[i]
			if d < 0 {
				d = oppositeNum(d)
			}
			if d > max {
				max = d
			}
		}
		return max
	}

	max := 0
	miceNumber := len(A)
	holeNumber := len(B)
	miceIndex := 0
	holeIndex := 0
	holeLeftLastIndex := -1
	for ; miceIndex < miceNumber; miceIndex++ {
		// means that holes more than mices
		for ; (holeNumber-holeIndex-2) > (miceNumber-miceIndex-1) && A[miceIndex] > B[holeIndex]; holeIndex++ {
		}
		if A[miceIndex] < B[holeIndex] {
			holeIndex--
		}
		// the left hole is already used by a mice
		if holeIndex == holeLeftLastIndex {
			// we could only use right hole
			d := A[miceIndex] - B[holeIndex+1]
			if d < 0 {
				d = oppositeNum(d)
			}
			if d > max {
				max = d
			}
			holeLeftLastIndex = holeIndex + 1
			holeIndex++
			continue
		}

		// we use both left and right holes
		// count distance to left hole
		d := A[miceIndex] - B[holeIndex]
		if d < 0 {
			d = oppositeNum(d)
		}
		holeLeftLastIndex = holeIndex
		// count distance to right hole
		d1 := A[miceIndex] - B[holeIndex+1]
		if d1 < 0 {
			d1 = oppositeNum(d1)
		}

		// take right hole only if distance is less than to left hole
		if d1 < d {
			d = d1
			holeLeftLastIndex = holeIndex + 1
		}
		if d > max {
			max = d
		}
		holeIndex++
	}

	return max
}

func oppositeNum(x int) int {
	return (^x) + 1
}
