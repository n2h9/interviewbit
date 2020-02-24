package solution

func isPalindrome(A int) int {
	if A < 0 {
		return 0
	}
	if A < 10 {
		return 1
	}
	if A%10 == 0 {
		return 0
	}

	i := 100000000000

	// deterine rank
	for ; i > 0; i /= 10 {
		if A/i == 0 {
			continue
		}
		break
	}

	// divide by 100 because we cut left and right digit every time
	for ; i > 0; i /= 100 {
		if A/i != A%10 {
			// not a palindrome
			return 0
		}
		// drop highest and lowest digits
		A = (A - A/i*i - A%10) / 10
	}

	// this is  palindrop
	return 1
}
