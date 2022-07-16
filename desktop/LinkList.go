package main

type Node[T any] struct {
	prev *Node[T]
	next *Node[T]
	data *T
}

type LinkedList[T any] struct {
	root   *Node[T]
	length int
}

func NewLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{}
}

func (l *LinkedList[T]) PushBack(t *T) {
	n := &Node[T]{
		prev: nil,
		next: nil,
		data: t,
	}
	if l.root == nil {
		l.root = n
		n.next = n
		n.prev = n
	} else {
		last := l.root.prev
		last.next = n
		l.root.prev = n
		n.next = l.root
		n.prev = last
	}
	l.length++
}

func (l *LinkedList[T]) Len() int {
	return l.length
}

func (l *LinkedList[T]) PopBack() *T {
	if l.root == nil {
		return nil
	}
	l.length--
	if l.root.next == l.root.prev {
		tmp := l.root.next
		if tmp == l.root {
			l.root = nil
		} else {
			l.root.next = l.root
			l.root.prev = l.root
		}
		return tmp.data
	} else {
		tmp := l.root.prev
		tmp.prev = l.root
		return tmp.data
	}
}
