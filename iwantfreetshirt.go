package main

type ListNode struct {
	next *ListNode
	prev *ListNode
	val  interface{}
}

func (l *ListNode) Value() interface{} {
	return l.val
}

func (l *ListNode) Next() *ListNode {
	return l.next
}

func (l *ListNode) Prev() *ListNode {
	return l.prev
}

// DoublyLinked List
type LinkedList struct {
}

func main() {}
