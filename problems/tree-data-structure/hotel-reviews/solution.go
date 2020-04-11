package solution

import (
	"sort"
	"strings"
)

func solve(A string, B []string) []int {
	goodWords := strings.Split(A, "_")
	goodWordsMap := make(map[string]bool, len(goodWords))
	for _, s := range goodWords {
		goodWordsMap[s] = true
	}

	reviews := make([]*review, len(B))
	for i, v := range B {
		reviews[i] = &review{
			index:  i,
			weight: coundWords(strings.Split(v, "_"), goodWordsMap),
		}
	}

	sort.Sort(reviewSlice(reviews))

	res := make([]int, len(reviews))
	for i, r := range reviews {
		res[i] = r.index
	}
	return res
}

type review struct {
	index  int
	weight int
}

type reviewSlice []*review

// Len is the number of elements in the collection.
func (s reviewSlice) Len() int {
	return len(s)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (s reviewSlice) Less(i, j int) bool {
	if s[i].weight == s[j].weight {
		return s[i].index < s[j].index
	}
	return s[i].weight > s[j].weight
}

// Swap swaps the elements with indexes i and j.
func (s reviewSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func coundWords(text []string, words map[string]bool) int {
	counter := 0
	for _, w := range text {
		if words[w] {
			counter++
		}
	}
	return counter
}
