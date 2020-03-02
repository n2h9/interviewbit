package solution

const maxBitsNum uint = 32
const modulo = 1000000000 + 7

func cntBits(A []int) int {
	n := len(A)
	if n <= 0 {
		return 0
	}
	val := 0
	var zeros int
	var i uint
	for i = 0; i < maxBitsNum; i++ {
		zeros = 0
		for j := 0; j < n; j++ {
			if (A[j] & (1 << i)) == 0 {
				zeros++
			}
		}
		val += zeros * (n - zeros) * 2
		// do not put zero in else condition because val is already full of zeros :)
	}

	return val % modulo
}
