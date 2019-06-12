package singly_linked_list

import "fmt"

func SinglyLinkedListTest() {
	list := NewList()

	list.Insert(1, 6)
	list.Insert(2, 2)
	list.Insert(3, 7)
	list.Insert(4, 8)
	list.Insert(5, 3)
	list.Insert(6, 8)
	list.Insert(7, 1)
	list.Insert(8, 3)
	list.Insert(9, 4)
	list.Insert(1, 0)
	list.Insert(6, 0)
	list.Insert(12, 0)
	list.Insert(-1, 0)
	list.Insert(14, 0)

	list.Delete(1)
	list.Delete(5)
	list.Delete(10)
	list.Delete(-1)
	list.Delete(10)

	fmt.Println(list.Get(1))
	fmt.Println(list.Get(5))
	fmt.Println(list.Get(9))
	fmt.Println(list.Get(-1))
	fmt.Println(list.Get(10))

	fmt.Println(list.Locate(6))
	fmt.Println(list.Locate(3))
	fmt.Println(list.Locate(4))
	fmt.Println(list.Locate(9))

	fmt.Println(list.Length())

	list.Print()
}
