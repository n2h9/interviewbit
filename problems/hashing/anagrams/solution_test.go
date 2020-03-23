package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol(t *testing.T) {
	A := []string{"cat", "dog", "god", "tca"}
	expected := [][]int{[]int{1, 4}, []int{2, 3}}

	assert.Equal(t, expected, anagrams(A))
}

func TestSol2(t *testing.T) {
	A := []string{"cat", "dog"}
	expected := [][]int{[]int{1}, []int{2}}

	assert.Equal(t, expected, anagrams(A))
}

func TestSol3(t *testing.T) {
	A := []string{
		"abbbaabbbabbbbabababbbbbbbaabaaabbaaababbabbabbaababbbaaabbabaabbaabbabbbbbababbbababbbbaabababba",
		"abaaabbbabaaabbbbabaabbabaaaababbbbabbbaaaabaababbbbaaaabbbaaaabaabbaaabbaabaaabbabbaaaababbabbaa",
		"babbabbaaabbbbabaaaabaabaabbbabaabaaabbbbbbabbabababbbabaabaabbaabaabaabbaabbbabaabbbabaaaabbbbab",
		"bbbabaaabaaaaabaabaaaaaaabbabaaaabbababbabbabbaabbabaaabaabbbabbaabaabaabaaaabbabbabaaababbaababb",
		"abbbbbbbbbbbbabaabbbbabababaabaabbbababbabbabaaaabaabbabbaaabbaaaabbaabbbbbaaaabaaaaababababaabab",
		"aabbbbaaabbaabbbbabbbbbaabbababbbbababbbabaabbbbbbababaaaabbbabaabbbbabbbababbbaaabbabaaaabaaaaba",
		"abbaaababbbabbbbabababbbababbbaaaaabbbbbbaaaabbaaabbbbbbabbabbabbaabbbbaabaabbababbbaabbbaababbaa",
		"aabaaabaaaaaabbbbaabbabaaaabbaababaaabbabbaaaaababaaabaabbbabbababaabababbaabaababbaabbabbbaaabbb",
	}
	expected := [][]int{
		[]int{1},
		[]int{2},
		[]int{3, 5},
		[]int{4},
		[]int{6},
		[]int{7},
		[]int{8},
	}

	assert.Equal(t, expected, anagrams(A))
}
