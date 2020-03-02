package solution

const MAX_INT = 1<<31 - 1
const MIN_INT = ^MAX_INT
const uintSize = 32 << (^uint(0) >> 32 & 1) // 32 or 64

// The solution with bit div
func divide(A int, B int) int {
	// division by zero is not allowed :)
	if B == 0 {
		return MAX_INT
	}
	var shouldChangeResSign bool
	switch {
	case A < 0 && B > 0:
		A = oppositeNum(A)
		shouldChangeResSign = true
	case A > 0 && B < 0:
		B = oppositeNum(B)
		shouldChangeResSign = true
	case A < 0 && B < 0:
		A = oppositeNum(A)
		B = oppositeNum(B)
	}
	if A < B {
		return 0
	}

	val := 0
	switch {
	case B == 1:
		val = A
	case A == B:
		val = 1
	default:
		vu, _ := div(uint(A), uint(B))
		val = int(vu)
	}

	if shouldChangeResSign {
		val = oppositeNum(val)
	}
	// overflow case
	if val > MAX_INT || val < MIN_INT {
		return MAX_INT
	}
	return val
}

func oppositeNum(x int) int {
	return (^x) + 1
}

// divide binary x / y
// Restrictions:
// - x > y
// - y > 1
// return result and carry
func div(x, y uint) (res uint, carry uint) {
	res, carry = 0, 0
	xL := bitLen(x)
	yL := bitLen(y)

	// number of next bit to process
	// (count from low bit to high)
	// where 0 is lowest bit
	i := xL - yL
	// a part to div to y each iteration
	carry = x >> (xL - yL)
	for {
		// allocate a place for next digit of result
		res <<= 1
		// nothing to do if carry < y
		// we need to put 0 to res, and this is already done while alocating bit in result
		if carry >= y {
			// 1) write 1 to res
			res += 1
			// 2) substract y from carry so carry is again smaller than y
			carry -= y
		}
		if i == 0 {
			// stop when processed all digits
			break
		}
		// prepare data for next iteration
		i--
		// add next bit to d for next iteration
		carry <<= 1
		carry += bitVal(x, i)
	}

	return res, carry
}

// returns bit len without leading 0s
func bitLen(x uint) uint {
	var i uint = 1
	for ; i <= uintSize; i++ {
		if x != (x << i >> i) {
			return uintSize - i + 1
		}
	}

	return 1
}

// value of certain bit x with index i
// wherer i = 0 for low bit and i = 32 for high bit
func bitVal(x, i uint) uint {
	return (x & (1 << i)) >> i
}
