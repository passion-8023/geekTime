package countingSort


func CountingSort(nums []int) []int {
	n := len(nums)
	if n < 2 {
		return nums
	}
	//获取数组的最大值
	max := nums[0]
	for _, num := range nums {
		if max < num {
			max = num
		}
	}
	//定义一个长度为max+1的切片，用来计数，相当于划分了max+1个桶，每个桶里的数据值都相同
	bucketCount := max + 1
	bucket := make([]int, bucketCount)
	//计算每个相同元素的个数，然后以元素值为key，出现的个数作为value放入到bucket中
	for _, num := range nums {
		bucket[num]++
	}

	//依次累加
	for i := 1; i < bucketCount; i++ {
		bucket[i] = bucket[i] + bucket[i-1]
	}

	//最终排序的结果切片
	res := make([]int, n)
	//原生数组倒序遍历
	for i := n - 1; i >= 0 ; i-- {
		var index int = bucket[nums[i]]-1
		res[index] = nums[i]
		bucket[nums[i]]--
	}
	return res
}
