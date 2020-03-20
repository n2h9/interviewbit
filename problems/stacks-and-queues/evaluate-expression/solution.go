package solution

import (
	"log"
	"strconv"
)

type listNode struct {
	value int
	next  *listNode
}

func listNode_new(val int) *listNode {
	var node *listNode = new(listNode)
	node.value = val
	node.next = nil
	return node
}

func evalRPN(A []string) int {
	l := len(A)
	if l <= 0 {
		return 0
	}

	stack := newStack()
	for _, s := range A {
		if !isOperator(s) {
			i, err := strconv.Atoi(s)
			if err != nil {
				log.Fatal(err)
			}
			stack = push(stack, i)
			continue
		}
		v2, ok := top(stack)
		if !ok {
			log.Fatal("stack is empty")
		}
		stack = pop(stack)
		v1, ok := top(stack)
		if !ok {
			log.Fatal("stack is empty")
		}
		stack = pop(stack)
		stack = push(stack, calculateData(v1, v2, s))
	}

	v, ok := top(stack)
	if !ok {
		log.Fatal("stack is empty")
	}

	return v
}

func newStack() *listNode {
	return nil
}

func isEmpty(stack *listNode) bool {
	return stack == nil
}

func push(stack *listNode, v int) *listNode {
	node := listNode_new(v)
	node.next = stack
	return node
}

func pop(stack *listNode) *listNode {
	if isEmpty(stack) {
		return stack
	}
	node := stack.next
	stack.next = nil
	return node
}

func top(stack *listNode) (int, bool) {
	if isEmpty(stack) {
		return 0, false
	}
	return stack.value, true
}

func isOperator(s string) bool {
	if len(s) > 1 {
		return false
	}
	return s[0] == '+' ||
		s[0] == '-' ||
		s[0] == '*' ||
		s[0] == '/'
}

func calculateData(x, y int, op string) int {
	switch {
	case op == "+":
		x += y
	case op == "-":
		x -= y
	case op == "/":
		x /= y
	case op == "*":
		x *= y
	}

	return x
}
