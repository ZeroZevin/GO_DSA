package selection_sort

func Sort(nums []int) {
	length := len(nums)

	for i := 0; i < length; i++ {
		max := 0

		for j := 1; j < length-i; j++ {
			if nums[max] < nums[j] {
				max = j
			}
		}

		nums[max], nums[length-i-1] = nums[length-i-1], nums[max]
	}
}
