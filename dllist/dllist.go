package dllist

type ListNode struct {
	Value any

	prev, next *ListNode
}

func (l *ListNode) Next() *ListNode {
	return l.next
}

type DLList struct {
	head, tail *ListNode

	Size int
}

func New() DLList {
	return DLList{
		head: nil,
		tail: nil,
		Size: 0,
	}
}

func (dl *DLList) initList(val any) *ListNode {
	newNode := &ListNode{
		Value: val,
		next:  nil,
		prev:  nil,
	}
	dl.head = newNode
	dl.tail = newNode

	dl.Size++

	return dl.head
}

func (dl *DLList) PushBack(val any) *ListNode {
	if dl.head == nil && dl.tail == nil {
		return dl.initList(val)
	}

	oldTail := dl.tail
	newTail := &ListNode{
		Value: val,
		next:  nil,
		prev:  oldTail,
	}

	dl.Size++

	dl.tail.next = newTail
	dl.tail = newTail

	return dl.tail
}

func (dl *DLList) PushFront(val any) *ListNode {
	if dl.head == nil && dl.tail == nil {
		return dl.initList(val)
	}

	oldHead := dl.head
	newHead := &ListNode{
		Value: val,
		next:  oldHead,
		prev:  nil,
	}

	dl.Size++

	dl.head.prev = newHead
	dl.head = newHead

	return dl.head
}

func (dl *DLList) Delete(val any) *ListNode {
	if dl.Size == 0 || dl.Size == 1 {
		return nil
	}

	tmp := dl.head
	for ; tmp != nil; tmp = tmp.next {
		if tmp.Value == val {
			if tmp.prev != nil && tmp.next != nil {
				prev := tmp.prev
				next := tmp.next

				prev.next = next
				next.prev = prev
			} else if tmp.prev == nil {
				dl.head = tmp.next
				tmp.next.prev = nil
			} else if tmp.next == nil {
				dl.tail = tmp.prev
				tmp.prev.next = nil
			}

			tmp.next = nil
			tmp.prev = nil
			return tmp
		}
	}

	return nil
}

func (dl *DLList) Remove(node *ListNode) *ListNode {
	tmp := node
	if tmp.prev != nil && tmp.next != nil {
		prev := tmp.prev
		next := tmp.next

		prev.next = next
		next.prev = prev
	} else if tmp.prev == nil {
		dl.head = tmp.next
		tmp.next.prev = nil
	} else if tmp.next == nil {
		dl.tail = tmp.prev
		tmp.prev.next = nil
	}

	tmp.next = nil
	tmp.prev = nil
	return tmp
}

func (dl *DLList) InsertToBack(node *ListNode) *ListNode {
	node.prev = dl.tail
	dl.tail.next = node

	dl.tail = node

	return dl.tail
}

func (dl *DLList) MoveToBack(node *ListNode) *ListNode {
	if node == dl.tail {
		return dl.tail
	}

	dl.Remove(node)
	dl.InsertToBack(node)

	return dl.tail
}

func (dl *DLList) Front() *ListNode {
	return dl.head
}

func (dl *DLList) Back() *ListNode {
	return dl.tail
}
