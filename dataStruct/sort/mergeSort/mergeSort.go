package mergeSort

// 归并排序
/**
1、归并排序是稳定排序吗？
	是一个稳定排序
2、归并排序的时间复杂度是多少？
	O(nlogn)
3、归并排序的空间复杂度是多少？
	O(n)
 */

func MergeSort(nums []int) []int {
	n := len(nums)
	if n <= 1 {
		return nums
	}
	var mid int = n / 2
	left := nums[0:mid]
	right := nums[mid:n]
	return merge(MergeSort(left), MergeSort(right));
}


func merge(left, right []int) []int {
	result := make([]int, len(left)+len(right))
	for index, i, j := 0, 0, 0; index < len(result); index++ {
		if i >= len(left) {
			result[index] = right[j]
			j++
		} else if j >= len(right){
			result[index] = left[i]
			i++
		} else if left[i] > right[j] {
			result[index] = right[j]
			j++
		} else {
			result[index] = left[i]
			i++
		}
	}
	return result
}


