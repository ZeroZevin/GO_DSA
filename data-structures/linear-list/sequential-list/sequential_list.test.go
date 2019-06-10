package sequential_list

import "fmt"

func SequentialListTest() {
	list := NewList()

	list.Insert(1, 6)
	list.Insert(2, 2)
	list.Insert(3, 7)
	list.Insert(4, 0)
	list.Insert(4, 8)
	list.Insert(5, 3)
	list.Insert(6, 8)
	list.Insert(7, 1)
	list.Insert(8, 3)
	list.Insert(9, 4)

	list.Delete(10)

	locateIndex, err := list.Locate(3)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("查询3 结果:第%d位\n", locateIndex)
	}

	fmt.Println(list.Get(1))

	list.Print()

	fmt.Println(list.Length())

	fmt.Println(list)
}
