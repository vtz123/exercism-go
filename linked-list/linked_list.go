package linkedlist

import (
	"errors"
)

type Node struct {
	pre, next *Node
	Val interface{}
}

type List struct {
	head,tail *Node
}

var ErrEmptyList = errors.New("2333")

func (e *Node) Next() *Node {
	if e == nil {
		return nil
	}

	return e.next
}

func (e *Node) Prev() *Node {
	if e == nil {
		return nil
	}

	return e.pre
}

func NewList(args ...interface{}) *List {

	list := &List{
		head:	nil,
	}

	h := new(Node)
	hcopy := h

	for i := range args{
		h.next = &Node{
			pre:	h,
			next:	nil,
			Val:	args[i],
		}
		h = h.Next()
	}

	list.head = hcopy.Next()
	if list.head != nil {
		list.head.pre = nil
	}

	list.tail = h
	return list
}

func (l *List) PushFront(v interface{}) {
	if l.head == nil {
		l.head = &Node{
			pre:   nil,
			next:  nil,
			Val: v,
		}
		l.tail = l.head

		return
	}

	newnode := &Node{
		pre:   nil,
		next:  l.head,
		Val: v,
	}
	l.head.pre = newnode
	l.head = newnode

	return
}

func (l *List) PushBack(v interface{}) {
	if l.head == nil {
		l.head = &Node{
			pre:   nil,
			next:  nil,
			Val: v,
		}
		l.tail = l.head
		return
	}

	newnode := &Node{
		pre:   l.tail,
		next:  nil,
		Val: v,
	}
	l.tail.next = newnode
	l.tail = newnode

	return
}

func (l *List) PopFront() (interface{}, error) {
	if l.head == nil {
		return nil, ErrEmptyList
	}

	node := l.head

	l.head = l.head.next
	if l.head == nil {
		l.tail = nil
		return node.Val, nil
	}else if l.head.next == nil {
		l.tail = l.head
	}
	l.head.pre = nil

	return node.Val, nil
}

func (l *List) PopBack() (interface{}, error) {
	if l.tail == nil {
		return nil, ErrEmptyList
	}

	node := l.tail
	if l.tail.pre == nil {
		l.head = nil
		l.tail = nil
	}else {
		l.tail = l.tail.pre
		l.tail.next = nil

	}

	return node.Val, nil
}

func (l *List) Reverse() *List {
	if l.head == nil {
		return l
	}else if l.head.next == nil {
		return l
	}

	left := &Node{}
	left1 := left
	headcopy := l.head
	var right *Node
	for headcopy.next != nil {
		//fmt.Println(headcopy.Val)

		right = headcopy.next

		headcopy.next = left
		left.pre = headcopy

		left = headcopy
		headcopy = right

	}

	headcopy.next = left
	left.pre = headcopy

	l.tail = left1.pre
	l.tail.next = nil
	left1.pre = nil

	l.head = headcopy
	headcopy.pre = nil


	return l
}

func (l *List) First() *Node {

	//fmt.Println(l.head.Val)

	return l.head
}

func (l *List) Last() *Node {
	return l.tail
}