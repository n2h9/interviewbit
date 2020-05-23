package solution

/**
 * @input A : array of strings
 *
 * @Output Integer
 */
func blackNoRec(A []string) int {
	n := len(A)
	if n <= 0 {
		return 0
	}
	m := len(A[0])
	res := 0

	visitedChars := make(map[int]bool, n*m)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if A[i][j] != SHAPE_BYTE {
				continue
			}
			if visitedChars[i*m+j] {
				continue
			}
			// push current node to stack
			stack := stackNode_new(&stackNodeValue{i, j})
			for !stack_isEmpty(stack) {
				ii := stack.value.i
				jj := stack.value.j
				// mark current node as visited
				visitedChars[ii*m+jj] = true
				// release node from stack
				stack = stack_pop(stack)
				// add every adjacent not visited SHAPE_BYTE node to stack
				coords := [][]int{
					[]int{ii - 1, jj}, // upper
					[]int{ii + 1, jj}, // bottom
					[]int{ii, jj - 1}, // left
					[]int{ii, jj + 1}, // right
				}
				for _, coord := range coords {
					if coord[0] < 0 || coord[0] >= n {
						continue
					}
					if coord[1] < 0 || coord[1] >= m {
						continue
					}
					if A[coord[0]][coord[1]] != SHAPE_BYTE {
						continue
					}
					if visitedChars[coord[0]*m+coord[1]] {
						continue
					}
					stack = stack_push(
						stack,
						stackNode_new(&stackNodeValue{coord[0], coord[1]}),
					)
				}
			}
			res++
		}
	}
	return res
}

type stackNodeValue struct {
	i int
	j int
}

type stackNode struct {
	value *stackNodeValue
	next  *stackNode
}

func stackNode_new(value *stackNodeValue) *stackNode {
	return &stackNode{
		value: value,
		next:  nil,
	}
}

func stack_isEmpty(stack *stackNode) bool {
	return stack == nil
}

func stack_pop(stack *stackNode) *stackNode {
	next := stack.next
	stack.next = nil
	return next
}

func stack_push(stack *stackNode, node *stackNode) *stackNode {
	node.next = stack
	return node
}
