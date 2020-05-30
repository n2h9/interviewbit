package solution

/**
 * @input A : String
 * @input B : String
 * @input C : array of strings
 *
 * @Output Integer
 */
func solve(A string, B string, C []string) int {
	n := len(C)
	if n <= 0 || len(A) <= 0 || len(B) <= 0 || A == B {
		return 0
	}
	if isMaxOneCharDiff(A, B) {
		return 2
	}

	const verticeA = -1
	const verticeB = -2
	adjacentMap := make(map[int][]int, n+2)

	for i := 0; i < n; i++ {
		if isMaxOneCharDiff(A, C[i]) {
			adjacentMap[verticeA] = append(adjacentMap[verticeA], i)
			adjacentMap[i] = append(adjacentMap[i], verticeA)
		}
		if isMaxOneCharDiff(B, C[i]) {
			adjacentMap[verticeB] = append(adjacentMap[verticeB], i)
			adjacentMap[i] = append(adjacentMap[i], verticeB)
		}
		for j := i + 1; j < n; j++ {
			if isMaxOneCharDiff(C[i], C[j]) {
				adjacentMap[i] = append(adjacentMap[i], j)
				adjacentMap[j] = append(adjacentMap[j], i)
			}
		}
	}

	q := newQueue()
	q = queuePush(
		q,
		newNode(
			newValue(verticeA, 1),
		),
	)
	visited := make(map[int]bool, n+2)

	for !queueIsEmpty(q) {
		val := queueTop(q).value
		q = queuePop(q)
		if visited[val.vertice] {
			continue
		}
		visited[val.vertice] = true
		length := val.length + 1
		for _, vertice := range adjacentMap[val.vertice] {
			if vertice == verticeB {
				return length
			}
			if visited[vertice] {
				continue
			}
			q = queuePush(
				q,
				newNode(
					newValue(vertice, length),
				),
			)
		}
	}

	return 0
}

type value struct {
	vertice int
	length  int
}

type node struct {
	value *value
	next  *node
	prev  *node
}

type queue struct {
	head *node
	tail *node
}

func newValue(vertice, length int) *value {
	return &value{
		vertice: vertice,
		length:  length,
	}
}

func newNode(value *value) *node {
	return &node{
		value: value,
	}
}

func newQueue() *queue {
	return &queue{}
}

func queueIsEmpty(q *queue) bool {
	return q.head == nil
}

func queuePush(q *queue, n *node) *queue {
	if queueIsEmpty(q) {
		q.head = n
		q.tail = n
		return q
	}
	n.next = q.tail
	q.tail.prev = n
	q.tail = n
	return q
}

func queuePop(q *queue) *queue {
	prev := q.head.prev
	if prev != nil {
		prev.next = nil
	}
	q.head.prev = nil
	q.head = prev
	if q.head == nil {
		q.tail = nil
	}
	return q
}

func queueTop(q *queue) *node {
	return q.head
}

// a,b same len
func isMaxOneCharDiff(a, b string) bool {
	n := len(a)
	diff := 0
	for i := 0; i < n; i++ {
		if a[i] != b[i] {
			if diff > 0 {
				return false
			}
			diff++
		}
	}

	return true
}
