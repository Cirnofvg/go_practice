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

	fmt.Println()
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

	fmt.Println(l.Front().Value, l.Back().Value)

	l.PushFront(11)
	fmt.Println(l.Front().Value, l.Back().Value)

	l.PushFront(13)
	fmt.Println(l.Front().Value, l.Back().Value)

	l.PushBack(15)
	fmt.Println(l.Front().Value, l.Back().Value)

	l.PrintList()

	l.PopBack()
	fmt.Println(l.Front().Value, l.Back().Value)

	l.PrintList()

}

func main() {
	doSmthWithList()

	// wp.DoPrint(5)

	// doSmthWithLRU()

	// doSmthWithDllist()

	// fibonacci.Fibonacci(10)

	var s stack.Stack

	// values := []any{"loh", "danil", "artyom"}

	s = stack.New("values")

	s.Push(1)
	s.Push(true)
	s.Push("loh")
	s.Push(2)

	// fmt.Println(s.Peek())
	// s.Pop()
	// fmt.Println(s.Peek())
	// s.Pop()
	// fmt.Println(s.Peek())
	// s.Pop()
	// fmt.Println(s.Peek())
	// s.Pop()
	// fmt.Println(s.Peek())
	// s.Pop()
	// fmt.Println(s.Peek())
	// s.Pop()
	// fmt.Println(s.Peek())

	var queue qu.QueueFromStack
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	queue.Push(4)

	// fmt.Println(queue.Peek())
	// fmt.Println(queue.Pop())
	// fmt.Println(queue.Peek())
}
