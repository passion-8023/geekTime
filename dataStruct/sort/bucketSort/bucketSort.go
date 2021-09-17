package bucketSort

func BucketSort(nums []int) []int {
	n := len(nums)
	if n < 2 {
		return nums
	}
	//获取数组中的最大最小值
	max, min := nums[0], nums[0]
	for i := 0; i < n; i++ {
		if nums[i] > max {
			max = nums[i]
		}
		if nums[i] < min {
			min = nums[i]
		}
	}

	array := make([]int, max+1)
	for _, val := range nums {
		array[val]++
	}
	res := make([]int, 0, n)
	for index, val := range array {
		for val > 0 {
			res = append(res, index)
			val--
		}
	}
	return res
}
