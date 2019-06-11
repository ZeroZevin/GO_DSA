package singly_linked_list

type Node struct {
	data interface{}
	next *Node
}

type SinglyLinkedList struct {
	first *Node
}
