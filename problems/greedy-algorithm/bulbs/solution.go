package solution

/**
 * @input A : Integer array
 *
 * @Output Integer
 */
func bulbs(A []int) int {
	output := 0
	// instead of going through the array and change all stage we consider 0 means turn on and 1 means turn off
	// starting with 1 is turn on
	turnOnVal := 1
	for i := 0; i < len(A); i++ {
		if A[i] == turnOnVal {
			continue
		}
		// switch on
		output++
		// change turnOnVal
		if turnOnVal == 1 {
			turnOnVal = 0
		} else {
			turnOnVal = 1
		}
	}

	return output
}
