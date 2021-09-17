package selectSort

//选择排序
/**
1、选择排序是原地排序吗？
	是，空间复杂度O(1)
2、选择排序是稳定排序吗？
	不是，选择排序每次都要从未排序元素中找到最小值，然后和前面的元素交换位置，这样就破坏了原有的元素顺序
3、选择排序的时间复杂度是多少？
	最好、最坏、平均都是O(n^2)
 */

func SelectSort(nums []int) []int {
	n := len(nums)
	if n <= 1 {return nums}
	for i := 0; i < n-1; i++ {
		min := nums[i]
		minIndex := i
		for j := i+1; j < n; j++ {
			if min > nums[j] {
				min = nums[j]
				minIndex = j
			}
		}
		if minIndex != i {
			nums[i], nums[minIndex] = nums[minIndex], nums[i]
		}
	}
	return nums
}

