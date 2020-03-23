package solution

import (
	"strconv"
)

/**
 * @input A : Integer
 * @input B : Integer
 *
 * @Output string.
 */
func fractionToDecimal(A int, B int) string {
	if B == 0 {
		return "infinity"
	}
	if B == 1 {
		return conv(A)
	}
	sign := ""
	switch {
	case A < 0 && B < 0:
		// negative / negative will be positive
		A = negative(A)
		B = negative(B)
	case A < 0 && B > 0:
		A = negative(A)
		sign = "-"
	case A > 0 && B < 0:
		B = negative(B)
		sign = "-"
	}

	intPart := A / B

	A -= intPart * B

	if A <= 0 {
		return combine(sign, conv(intPart), "")
	}

	fractionalPart := ""
	periodDigitPosition := -1
	reminders := map[int]int{}
	// divide "in the column" utill find all fractional part or determine repeat
	for i := 0; A > 0; i++ {
		if v, ok := reminders[A]; ok {
			periodDigitPosition = v
			break
		}
		reminders[A] = i
		A *= 10
		fractionalPart += conv(A / B)
		A %= B
	}
	if periodDigitPosition != -1 {
		fractionalPart = processPeriod(fractionalPart, periodDigitPosition)
	}
	return combine(sign, conv(intPart), fractionalPart)
}

func conv(x int) string {
	return strconv.FormatInt(int64(x), 10)
}

func processPeriod(fractionPart string, periodDigitPosition int) string {
	fpb := []byte(fractionPart)
	// insert '(' and ')'
	fpb = append(fpb, 0)
	copy(fpb[periodDigitPosition+1:], fpb[periodDigitPosition:])
	fpb[periodDigitPosition] = '('
	fpb = append(fpb, ')')
	return string(fpb)
}

func negative(x int) int {
	return ^x + 1
}

func combine(sign, intPart, fractPart string) string {
	res := sign + intPart
	if fractPart != "" {
		res += "." + fractPart
	}
	return res
}
