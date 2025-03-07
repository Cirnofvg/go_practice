package main

import (
	"fmt"
	"practice/list"
	qu "practice/queue_from_stack"
	"practice/stack"

	"practice/dllist"

	"practice/lru"
)

func doSmthWithDllist() {
	l := dllist.New()
	l.PushBack(10)
	l.PushFront(11)
	n := l.PushFront(12)
	l.PushBack(23)
	l.PushFront(25)
	l.PushBack(30)

	// fmt.Println(l.Front().Value)
	l.Remove(l.Front())

	l.MoveToBack(n)
	fmt.Println(l.Back())
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	l.Delete(11)
	l.Delete(30)
	l.Delete(25)

	// for e := l.Front(); e != nil; e = e.Next() {
	// 	fmt.Println(e.Value)
	// }

	fmt.Println(l.ListToSlice())
}

func doSmthWithLRU() {
	lRUCache := lru.New(1)

	// LRUCache lRUCache = new LRUCache(2);
	lRUCache.Put(2, 1) // cache is {1=1}
	lRUCache.Get(2)    // return 1
	lRUCache.Put(3, 2)
	// lRUCache.Get(2)
	// lRUCache.Get(3)

}

func doSmthWithList() {
	l := list.New()
	l.PushBack(10)
	l.PopBack()
	fmt.Println(l.Front(), l.Back()) // nil.Back()
	// l.Front()

	// fmt.Println(l.PopBack())

	// fmt.Println(l.Front().Value, l.Back().Value)

	l.PushFront(11)
	fmt.Println(l.Front().Value, l.Back().Value)

	l.PushFront(13)
	fmt.Println(l.Front().Value, l.Back().Value)

	l.PushBack(15)
	fmt.Println(l.Front().Value, l.Back().Value)

	// l.PrintList()

	// fmt.Println(l.Front().Value, l.Back().Value)

	fmt.Println(l.ListToSlice())

	l.PopFront()
	fmt.Println(l.Front().Value, l.Back().Value)

	fmt.Println(l.ListToSlice())

	l.PopFront()
	l.PopFront()
	l.PopFront()
	l.PopBack()
	fmt.Println(l.ListToSlice())
	l.PopFront()

	l.PushBack(7)
	l.PushFront(8)
	fmt.Println(l.ListToSlice())

}

func doSmthWithStack() {
	var s stack.Stack

	values := []any{"aboba", "danil", "artyom"}

	s = stack.New(values)
	fmt.Println(s.Peek())

	s.Push(1)
	s.Push(true)
	s.Push("aboba")
	s.Push(2)

	fmt.Println(s.Peek())
	s.Pop()
	fmt.Println(s.Peek())
	s.Pop()
	fmt.Println(s.Peek())
	s.Pop()
	fmt.Println(s.Peek())
	s.Pop()
	fmt.Println(s.Peek())
	s.Pop()
	fmt.Println(s.Peek())
	s.Pop()
	fmt.Println(s.Peek())
}

func main() {
	doSmthWithStack()

	// doSmthWithList()

	// wp.DoPrint(5)

	// doSmthWithLRU()

	// doSmthWithDllist()

	// fibonacci.Fibonacci(10)

	var queue qu.QueueFromStack
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	queue.Push(4)

	// fmt.Println(queue.Peek())
	// fmt.Println(queue.Pop())
	// fmt.Println(queue.Peek())
}
