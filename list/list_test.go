package list

import (
	"testing"
)

func TestListInsertion(t *testing.T) {
	{
		const op = "test.list.PushBack1"
		l := New()

		l.PushBack(1)

		if l.Front().Value != 1 && l.Back().Value != 1 {
			t.Errorf("Test failed in %s", op)
		}
	}
	{
		const op = "test.list.PushBack2"
		l := New()

		l.PushBack(1)
		l.PushBack(10)

		if l.Back().Value != 10 && l.Front().Value != 1 {
			t.Errorf("Test failed in %s", op)
		}
	}
}

func TestListPop(t *testing.T) {
	l := New()

	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)

	{
		const op = "test.list.Pop1"

		n := l.PopBack()

		if n.Value != 3 && l.Front().Value != 1 && l.Back().Value != 2 {
			t.Errorf("Test failed in %s", op)
		}
	}
	{
		const op = "test.list.Pop2"

		n := l.PopBack()

		if n.Value != 2 && l.Front().Value != 1 && l.Back().Value != 1 {
			t.Errorf("Test failed in %s", op)
		}
	}
	{
		const op = "test.list.Pop3"

		n := l.PopBack()

		if n.Value != 1 && l.Front().Value != nil && l.Back().Value != nil {
			t.Errorf("Test failed in %s", op)
		}
	}
}
