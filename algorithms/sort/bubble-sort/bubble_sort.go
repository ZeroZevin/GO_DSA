package bubble_sort

func Sort(nums []int) {
	length := len(nums)
	flag := true

	for i := 0; i < length-1; i++ {
		flag = true

		for j := 0; j < length-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				flag = false
			}
		}

		if flag {
			break
		}
	}
}
