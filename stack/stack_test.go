package stack_test

import (
	"practice/stack"
	"testing"
)

func TestStack(t *testing.T) {
	{
		const op = "test.stack.IsEmpty1"

		s := stack.New([]any{1, 2, 3, 4})

		if s.IsEmpty() == true {
			t.Errorf("Test failed in %s", op)
		}
	}
	{
		const op = "test.stack.IsEmpty2"

		var s stack.Stack

		if s.IsEmpty() != true {
			t.Errorf("Test failed in %s", op)
		}
	}
	{
		const op = "test.stack.Peek1"

		s := stack.New([]any{1, 2, 3, 4})

		if s.Peek() != 4 {
			t.Errorf("Test failed in %s", op)
		}
	}
	{
		const op = "test.stack.Peek2"

		var s stack.Stack

		if s.Peek() != nil {
			t.Errorf("Test failed in %s", op)
		}
	}
	{
		const op = "test.stack.PopPeek1"

		s := stack.New([]any{1, 2, 3, 4})

		if p := s.Peek(); s.Pop() != p {
			t.Errorf("Test failed in %s", op)
		}

		if s.Peek() != 3 {
			t.Errorf("Test failed in %s", op)
		}
	}
	{
		const op = "test.stack.PopPeek2"

		var s stack.Stack

		if s.Pop() != nil {
			t.Errorf("Test failed in %s", op)
		}

		if s.Peek() != nil {
			t.Errorf("Test failed in %s", op)
		}
	}
	{
		const op = "test.stack.Push1"

		var s stack.Stack

		s.Push(5)

		if s.Peek() != 5 {
			t.Errorf("Test failed in %s", op)
		}

		s.Push(10)

		if s.Peek() != 10 {
			t.Errorf("Test failed in %s", op)
		}
	}
	{
		const op = "test.stack.Init1"

		var s stack.Stack

		if s.Peek() != nil {
			t.Errorf("Test failed in %s", op)
		}
	}
	{
		const op = "test.stack.Init2"

		s := stack.New([]int{1, 2, 3})

		if s.Peek() != 3 {
			t.Errorf("Test failed in %s", op)
		}
	}
}
