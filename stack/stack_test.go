package stack

import (
	// "practice/stack"
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	{
		const op = "test.stack.IsEmpty1"

		s := New([]any{1, 2, 3, 4})

		if s.IsEmpty() == true {
			t.Errorf("Test failed in %s", op)
		}
	}
	{
		const op = "test.stack.IsEmpty2"

		var s Stack

		if s.IsEmpty() != true {
			t.Errorf("Test failed in %s", op)
		}
	}
	{
		const op = "test.stack.Peek1"

		s := New([]any{1, 2, 3, 4})

		if s.Peek() != 4 {
			t.Errorf("Test failed in %s", op)
		}
	}
	{
		const op = "test.stack.Peek2"

		var s Stack

		if s.Peek() != nil {
			t.Errorf("Test failed in %s", op)
		}
	}
	{
		const op = "test.stack.PopPeek1"

		s := New([]any{1, 2, 3, 4})

		if p := s.Peek(); s.Pop() != p {
			t.Errorf("Test failed in %s", op)
		}

		if s.Peek() != 3 {
			t.Errorf("Test failed in %s", op)
		}
	}
	{
		const op = "test.stack.PopPeek2"

		var s Stack

		if s.Pop() != nil {
			t.Errorf("Test failed in %s", op)
		}

		if s.Peek() != nil {
			t.Errorf("Test failed in %s", op)
		}
	}
	{
		const op = "test.stack.Push1"

		var s Stack

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

		var s Stack

		if s.Peek() != nil {
			t.Errorf("Test failed in %s", op)
		}
	}
	{
		const op = "test.stack.Init2"

		s := New([]any{1, 2, 3})
		fmt.Println(s.Peek())
		if s.Peek() != 3 {
			t.Errorf("Test failed in %s", op)
		}
	}
}
