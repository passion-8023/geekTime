package main

import (
	"fmt"
	"geekTime/dataStruct/sort/countingSort"
)

func main() {
	//rand.Seed(time.Now().UnixNano())
	//fmt.Println(rand.Intn(5))

	nums := []int{4, 5, 6, 3, 2, 1}
	//nums := []int{3,2,1,4,5,6}
	//fmt.Println(bubbleSort.BubbleSort(nums))
	//fmt.Println(insertSort.InsertSort(nums))
	//fmt.Println(selectSort.SelectSort(nums))
	//fmt.Println(mergeSort.MergeSort(nums))
	//fmt.Println(quickSort.QuickSort(nums))
	//fmt.Println(bucketSort.BucketSort(nums))
	fmt.Println(countingSort.CountingSort(nums))


	//arr := make([]int, 9)
	//fmt.Println(len(arr))

	//fmt.Println(bubble(nums))
	//fmt.Println(insert(nums))
	//fmt.Println(selection(nums))
	//fmt.Println(merge(nums))
}

func bubble(nums []int) []int {
	n := len(nums)
	if n < 2 {return nums}
	for i := 0; i < n; i++ {
		flag := false
		for j := 0; j < n-1-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
	return nums
}

func insert(nums []int) []int {
	n := len(nums)
	if n < 2 {
		return nums
	}
	//for i := 1; i < n; i++ {
	//	val := nums[i]
	//	j := i - 1
	//	for j >= 0 && nums[j] > val {
	//		nums[j+1] = nums[j]
	//		j--
	//	}
	//	nums[j+1] = val
	//	fmt.Println(nums)
	//}

	for i := 0; i < n-1; i++ {
		val := nums[i+1]
		j := i
		for j >= 0 && nums[j] > val {
			nums[j+1] = nums[j]
			j--
		}
		nums[j+1] = val
	}

	return nums
}

func selection(nums []int) []int {
	n := len(nums)
	if n < 2 {
		return nums
	}
	for i := 0; i < n; i++ {
		index := i
		for j := i+1; j < n; j++ {
			if nums[j] < nums[index] {
				index = j
			}
		}
		nums[i], nums[index] = nums[index], nums[i]
	}
	return nums
}

func merge(nums []int) []int {
	n := len(nums)
	if n < 2 {
		return nums
	}
	var mid int = n / 2
	left := nums[0:mid]
	right := nums[mid:n]
	return mergeHandle(merge(left), merge(right))
}

func mergeHandle(left, right []int) []int {
	n, m := len(left), len(right)
	result := make([]int, n+m)
	for i, j, index := 0, 0, 0; index < len(result); index++ {
		if i >= n {
			result[index] = right[j]
			j++
		} else if j >= m {
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






