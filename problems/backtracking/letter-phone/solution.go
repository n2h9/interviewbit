package solution

var phoneButtons map[byte][]string = map[byte][]string{
	'0': []string{"0"},
	'1': []string{"1"},
	'2': []string{"a", "b", "c"},
	'3': []string{"d", "e", "f"},
	'4': []string{"g", "h", "i"},
	'5': []string{"j", "k", "l"},
	'6': []string{"m", "n", "o"},
	'7': []string{"p", "q", "r", "s"},
	'8': []string{"t", "u", "v"},
	'9': []string{"w", "x", "y", "z"},
}

/**
 * @input A : String
 *
 * @Output string array.
 */
func letterCombinations(A string) []string {
	n := len(A)
	if n <= 0 {
		return []string{}
	}
	return findAll(A, 0)
}

func findAll(A string, index int) []string {
	n := len(A)
	if (index + 1) == n {
		return phoneButtons[A[index]]
	}

	combinations := findAll(A, index+1)
	result := make(
		[]string,
		len(phoneButtons[A[index]])*len(combinations),
	)
	i := 0
	for _, v := range phoneButtons[A[index]] {
		for _, c := range combinations {
			result[i] = v + c
			i++
		}
	}

	return result
}
