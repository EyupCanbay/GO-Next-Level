package main

import "fmt"

// 1 -> 2 -> 3 ->

type MyItem struct {
	val int
}

type Node struct {
	Next *Node
	Data MyItem
}

type LinkedList struct {
	Head *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (ll *LinkedList) Insert(data MyItem) {
	n := &Node{
		Next: nil,
		Data: data,
	}

	if ll.Head == nil {
		ll.Head = n
		return
	}

	current := ll.Head
	for current.Next != nil {
		current = current.Next
	}

	current.Next = n

}

func (ll *LinkedList) Print() {
	current := ll.Head
	for current != nil {
		fmt.Printf("%d -> ", current.Data.val)
		current = current.Next
	}
	fmt.Print("nil")
}

func (ll *LinkedList) Delete(data MyItem) {
	if ll.Head == nil {
		return
	}

	if ll.Head.Data == data {
		ll.Head = ll.Head.Next
		return
	}

	current := ll.Head
	for current.Next != nil {
		if current.Next.Data == data {
			current.Next = current.Next.Next
			return
		}
		current = current.Next
	}
}
func main() {
	ll := NewLinkedList()
	ll.Insert(MyItem{3})
	ll.Insert(MyItem{4})
	ll.Insert(MyItem{5})
	ll.Insert(MyItem{6})
	ll.Insert(MyItem{7})
	ll.Insert(MyItem{8})

	ll.Delete(MyItem{8})
	ll.Print()
}
