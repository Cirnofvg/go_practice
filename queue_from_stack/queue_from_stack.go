package queuefromstack

import "practice/stack"

type QueueFromStack struct {
	stack1 stack.Stack // insert here
	stack2 stack.Stack // pop from here
}

func (q *QueueFromStack) Peek() any {
	if q.stack2.IsEmpty() {
		for !q.stack1.IsEmpty() {
			q.stack2.Push(q.stack1.Pop())
		}
	}

	return q.stack2.Peek()
}

func (q *QueueFromStack) Push(value any) {
	q.stack1.Push(value)
}

func (q *QueueFromStack) Pop() any {
	if q.stack2.IsEmpty() {
		for !q.stack1.IsEmpty() {
			q.stack2.Push(q.stack1.Pop())
		}
	}

	return q.stack2.Pop()
}
