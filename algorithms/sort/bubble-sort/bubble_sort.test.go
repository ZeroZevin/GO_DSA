package bubble_sort

import "fmt"

func SortTest() {
	nums := []int{6, 2, 7, 8, 3, 8, 1, 3, 4}

	Sort(nums)

	fmt.Println(nums)
}
