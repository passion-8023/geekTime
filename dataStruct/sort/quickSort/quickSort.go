package quickSort
//快速排序
/**
1、快速排序是原地排序吗？
	是
2、快速排序是稳定排序吗？
	不稳定
3、快速排序的时间复杂度是多少？
	快速排序的时间复杂度和分区是否均衡有很大的关系，最好情况是O(n)，最坏情况是O(n^2)，平均情况是O(nlogn)
 */
func QuickSort(nums []int) []int {
	n := len(nums)
	if n < 2 {
		return nums
	}
	quickSortHandle(nums, 0, n-1)
	return nums
}

func quickSortHandle(nums []int, start, end int) {
	if start > end {
		return;
	}
	var index int = partition(nums, start, end)
	quickSortHandle(nums, start, index - 1)
	quickSortHandle(nums, index + 1, end)
}

func partition(nums []int, start, end int) int {
	var pivot int = end
	i := start
	for j := start; j < end; j++ {
		if nums[j] < nums[pivot] {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	nums[i], nums[pivot] = nums[pivot], nums[i]
	return i
}
