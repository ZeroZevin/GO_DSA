package singly_linked_list

import (
	"errors"
	"fmt"
)

type Node struct {
	data interface{}
	next *Node
}

type SinglyLinkedList struct {
	first *Node
}

func NewList() *SinglyLinkedList {
	return new(SinglyLinkedList)
}

func (list *SinglyLinkedList) Length() int {
	size := 0
	p := list.first

	for p != nil {
		p = p.next
		size++
	}

	return size
}

func (list *SinglyLinkedList) Get(i int) (interface{}, error) {
	j := 0
	p := list.first

	if i < 0 || i > list.Length() {
		return nil, errors.New("下标超出范围")
	}

	for p != nil && j < i-1 {
		j++
		p = p.next
	}

	return p.data, nil
}

func (list *SinglyLinkedList) Locate(v interface{}) (int, error) {
	i := 0
	p := list.first

	for p != nil {
		if p.data == v {
			return i + 1, nil
		}
		i++
		p = p.next
	}

	return -1, errors.New("元素不存在")
}

func (list *SinglyLinkedList) Insert(i int, x interface{}) error {

	if i < 0 || i > list.Length()+1 {
		return errors.New("下标超出范围")
	}

	if i == 1 {
		list.first = &Node{x, list.first}
	} else {
		j := 0
		p := list.first

		for p != nil && j < i-2 {
			p = p.next
			j++
		}

		node := &Node{x, nil}

		if p.next != nil {
			node.next = p.next
		}

		p.next = node
	}

	return nil
}

func (list *SinglyLinkedList) Delete(i int) error {

	if i < 0 || i > list.Length() {
		return errors.New("下标超出范围")
	}

	if i == 1 {
		list.first = list.first.next
	} else {
		j := 0
		p := list.first

		for p != nil && j < i-2 {
			p = p.next
			j++
		}

		p.next = p.next.next

	}

	return nil
}

func (list *SinglyLinkedList) Print() {
	p := list.first
	i := 1

	for p != nil {
		fmt.Printf("第%d位是：%v\n", i, p.data)
		p = p.next
		i++
	}
}
