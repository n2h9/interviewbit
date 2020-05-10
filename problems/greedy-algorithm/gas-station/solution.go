package solution

/**
 * @input A : Integer array
 * @input B : Integer array
 *
 * @Output Integer
 */
func canCompleteCircuit(A []int, B []int) int {
	for startIndex := 0; startIndex < len(A); startIndex++ {
		gas := A[startIndex] - B[startIndex]
		if gas < 0 {
			continue
		}
		ok := true
		for j := (startIndex + 1) % len(A); j != startIndex; j = (j + 1) % len(A) {
			gas += (A[j] - B[j])
			if gas < 0 {
				ok = false
				break
			}
		}
		if ok {
			return startIndex
		}
	}
	return -1
}
