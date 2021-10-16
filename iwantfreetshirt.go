package main

import "fmt"

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

func main() {
	fmt.Println("GIVE ME SHIRT PLSSSSSSSSSSSSSSSSS i love you github and digitalocean. i use you daily. but i also love a free shirt. thx sorry. the waiting is the hardest part.")
}
