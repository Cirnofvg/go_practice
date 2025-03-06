package list

type List struct {
	head, tail *ListNode

	len int
}

type ListNode struct {
	Value any
	next  *ListNode
}

func New() List {
	return List{
		head: nil,
		tail: nil,
		len:  0,
	}
}

func (l *List) Front() *ListNode {
	return l.head
}

func (l *List) Back() *ListNode {
	return l.tail
}

func (l *List) initList(val any) *ListNode {
	newNode := &ListNode{
		Value: val,
		next:  nil,
	}
	l.head = newNode
	l.tail = newNode

	l.len++

	return l.head
}

func (l *List) PushBack(val any) *ListNode {
	if l.len == 0 {

	}
	newNode := &ListNode{
		Value: val,
		next:  nil,
	}

	l.tail.next = newNode
	l.tail = newNode

	l.len++

	return l.tail
}

func (l *List) PushFront(val any) *ListNode {
	newNode := &ListNode{
		Value: val,
		next:  l.head,
	}

	l.head = newNode

	return l.head
}

func (l *List) PopBack() *ListNode {
	if l.len == 0 {
		return nil
	}
	if l.len == 1 {
		l.len--
		l.head, l.tail = nil, nil
		return nil
	}

	tmp := l.head
	for ; tmp.next.next != nil; tmp = tmp.next {

	}

	toDelete := tmp.next
	tmp.next = nil
	l.len--

	return toDelete
}

func (l *List) PopFront() *ListNode {
	if l.len == 0 {
		return nil
	}
	if l.len == 1 {
		l.len--
		l.head, l.tail = nil, nil
		return nil
	}

	toDelete := l.head

	l.head = l.head.next

	toDelete.next = nil

	return toDelete
}
