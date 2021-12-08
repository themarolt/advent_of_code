package libs

import (
	"fmt"
)

type ListElement struct {
	list  List
	next  *ListElement
	prev  *ListElement
	Value interface{}
}

type List interface {
	Size() int
	First() *ListElement
	Last() *ListElement
	Get(i int) *ListElement
	Add(at *ListElement, data interface{}) *ListElement
	Push(data interface{}) *ListElement
	RemoveFrom(toRemove *ListElement)
	PrintLn()
}

type LinkedList struct {
	first *ListElement
	last  *ListElement
	size  int
}

func (e *ListElement) Next() *ListElement {
	return e.next
}

func (e *ListElement) Prev() *ListElement {
	return e.prev
}

func (l *LinkedList) Size() int {
	return l.size
}

func (l *LinkedList) First() *ListElement {
	return l.first
}

func (l *LinkedList) Last() *ListElement {
	return l.last
}

func (l *LinkedList) Get(i int) *ListElement {
	if i < l.size {
		el := l.first
		index := 0

		for index < i {
			el = el.next
			index++
		}

		return el
	}

	return nil
}

func (l *LinkedList) Add(at *ListElement, data interface{}) *ListElement {
	// create a new element
	newElement := ListElement{
		list:  l,
		next:  at,
		prev:  at.prev,
		Value: data,
	}

	if l.first == at {
		l.first = &newElement
	}

	if at.prev != nil {
		at.prev.next = &newElement
	}

	at.prev = &newElement

	l.size++

	return &newElement
}

func (l *LinkedList) Push(data interface{}) *ListElement {
	newElement := ListElement{
		list:  l,
		next:  nil,
		prev:  l.last,
		Value: data,
	}

	if l.size == 0 {
		l.first = &newElement
		l.last = &newElement
	} else {
		l.last.next = &newElement
		l.last = &newElement
	}

	l.size++

	return &newElement
}

func (l *LinkedList) RemoveFrom(toRemove *ListElement) {
	if l.size == 1 {
		l.first = nil
		l.last = nil
	} else if l.first == toRemove {
		l.first = l.first.next
		l.first.prev = nil
	} else if l.last == toRemove {
		l.last = l.last.prev
		l.last.next = nil
	} else {
		prev := toRemove.prev
		prev.next = toRemove.next
		toRemove.next.prev = prev
	}

	l.size--
}

func (l *LinkedList) PrintLn() {
	fmt.Println("List Size: ", l.size)
	fmt.Print("[")
	for i, e := 0, l.first; e != nil; e = e.next {
		fmt.Print(e.Value)
		if i != l.size-1 {
			fmt.Print(", ")
		}
		i++
	}
	fmt.Println("]")
}

func NewLinkedList() List {
	newList := new(LinkedList)

	return newList
}
