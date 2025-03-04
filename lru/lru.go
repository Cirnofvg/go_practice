package lru

import (
	list "practice/dllist"
)

type LRU struct {
	data list.DLList

	cache map[int]*list.ListNode

	len int
	cap int
}

type cacheElement struct {
	key   int
	value int
}

func New(cap int) LRU {
	return LRU{
		data:  list.New(),
		cache: make(map[int]*list.ListNode),
		len:   0,
		cap:   cap,
	}
}

func (l *LRU) Put(key int, val int) {
	if elem, ok := l.cache[key]; !ok {
		node := l.data.PushBack(cacheElement{key: key, value: val})
		l.cache[key] = node
		l.len++
		if l.len > l.cap {
			delete(l.cache, l.data.Front().Value.(cacheElement).key)
			l.data.Remove(l.data.Front())
		}
	} else {
		elem.Value = val
		l.data.MoveToBack(elem)
	}
}

func (l *LRU) Peek() int {
	return l.data.Back().Value.(cacheElement).value
}

func (l *LRU) Get(key int) int {
	if elem, ok := l.cache[key]; ok {
		l.data.MoveToBack(elem)
		return elem.Value.(cacheElement).value
	}

	return -1
}
