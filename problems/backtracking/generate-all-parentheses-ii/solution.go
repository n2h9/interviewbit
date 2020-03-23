package solution

/**
 * @input A : Integer
 *
 * @Output string array.
 */
func generateParenthesis(A int) []string {
	if A <= 0 {
		return []string{}
	}
	return solve("", 0, 2*A)
}

func solve(path string, openCount, maxLength int) []string {
	if openCount > (maxLength - len(path)) {
		return []string{}
	}
	if len(path) >= maxLength {
		if openCount != 0 {
			return []string{}
		}
		return []string{path}
	}

	// add open parentheses
	res := solve(path+"(", openCount+1, maxLength)

	// try to add close parentheses
	if openCount > 0 {
		res = append(res, solve(path+")", openCount-1, maxLength)...)
	}
	return res
}
