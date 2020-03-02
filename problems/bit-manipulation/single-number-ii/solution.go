package solution

const maxBitsNum uint = 32

func singleNumber(A []int) int {
	n := len(A)
	val := 0
	var zeros, unit int
	var i uint
	for i = 0; i < maxBitsNum; i++ {
		zeros = 0
		// use unit to extract bit and to set bit as well (see below)
		unit = (1 << i)
		for j := 0; j < n; j++ {
			if (A[j] & unit) == 0 {
				zeros++
			}
		}
		if zeros%3 == 0 {
			// number of zeros is divided by 3 without cary =>
			// need to put a unit in i bit
			val |= unit
		}
		// do not put zero in else condition because val is already full of zeros :)
	}

	return val
}
