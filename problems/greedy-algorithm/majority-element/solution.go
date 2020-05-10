package solution

/**
 * @input A : Integer array
 *
 * @Output Integer
 */
func majorityElement(A []int) int {
	majCondition := len(A) / 2
	occurenceNum := make(map[int]int, len(A))
	for _, v := range A {
		occurenceNum[v]++
		if occurenceNum[v] > majCondition {
			return v
		}
	}
	return -1
}
