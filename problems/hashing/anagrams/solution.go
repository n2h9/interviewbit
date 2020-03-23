package solution

import "sort"

/**
 * @input A : array of strings
 *
 * @Output 2D int array.
 */
func anagrams(A []string) [][]int {
	n := len(A)
	if n <= 0 {
		return [][]int{}
	}

	if n == 1 {
		return [][]int{[]int{1}}
	}

	anagrams := make(map[int][]int, n)
	order := make([]int, 0, n)
	for i, s := range A {
		h := hash(s)
		if anagrams[h] == nil {
			anagrams[h] = []int{i + 1}
			order = append(order, h)
			continue
		}
		anagrams[h] = append(anagrams[h], i+1)
	}
	ol := len(order)
	res := make([][]int, ol)
	for i := 0; i < ol; i++ {
		res[i] = anagrams[order[i]]
	}
	return res
}

const MAX_INT = 1<<31 - 1

type ByteSlice []byte

func (p ByteSlice) Len() int           { return len(p) }
func (p ByteSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p ByteSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p ByteSlice) Sort()              { sort.Sort(p) }

func hash(a string) int {
	b := ByteSlice(a)
	b.Sort()
	hash := 0
	for _, v := range b {
		hash = (hash*31 + int(v)) % MAX_INT
	}
	return hash
}
