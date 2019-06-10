package sequential_list

import (
	"errors"
	"fmt"
)

const MaxSize = 100

type SequentialList struct {
	data   [MaxSize]interface{}
	length int
}

func NewList() *SequentialList {
	list := new(SequentialList)

	list.length = 0

	return list
}

func (list *SequentialList) Length() int {
	return list.length
}

func (list *SequentialList) Get(i int) (interface{}, error) {
	if i < 1 || i > list.Length() {
		return nil, errors.New("下标超出范围")
	}

	return list.data[i-1], nil
}

func (list *SequentialList) Locate(v interface{}) (int, error) {
	index := 0
	found := false

	for i := 0; i < list.length; i++ {
		if list.data[i] == v {
			index = i + 1
			found = true
			break
		}
	}

	if found {
		return index, nil
	} else {
		return -1, errors.New("元素不存在")
	}
}

func (list *SequentialList) Insert(i int, v interface{}) error {
	if list.Length() >= MaxSize {
		return errors.New("List is full")
	}

	if i < 1 || i > list.Length()+1 {
		return errors.New("下标超出范围")
	}

	for j := list.Length(); j >= i; j-- {
		list.data[j] = list.data[j-1]
	}

	list.data[i-1] = v
	list.length += 1

	return nil
}

func (list *SequentialList) Delete(i int) (interface{}, error) {
	if i < 1 || i > list.Length() {
		return nil, errors.New("下标超出范围")
	}

	value := list.data[i-1]

	for j := i - 1; j < list.length; j++ {
		list.data[j] = list.data[j+1]
	}

	list.data[list.Length()-1] = nil

	list.length -= 1

	return value, nil
}

func (list *SequentialList) Print() {
	for i := 0; i < list.length; i++ {
		fmt.Printf("第%d位是：%v\n", i+1, list.data[i])
	}
}
