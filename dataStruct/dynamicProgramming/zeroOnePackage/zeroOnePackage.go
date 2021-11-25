package zeroOnePackage

import (
	"fmt"
	"math"
)

var (
	maxW int = math.MinInt
	maxV int = math.MinInt
	weight []int = []int{2, 2, 4, 6, 3}
	value []int = []int{3, 4, 8, 9, 6}
	n int = 5
	w int = 9
	//优化-1
	mem [5][10]bool
)

func HandlerFun(i, cw int) int {
	if cw == w || i == n {
		if cw > maxW {maxW = cw}
		return maxW
	}
	if mem[i][cw] {
		return maxW
	}
	mem[i][cw] = true
	HandlerFun(i+1, cw)
	if (cw + weight[i]) <= w {
		HandlerFun(i+1, cw + weight[i])
	}
	return maxW
}

func HandlerValueFun(i, cw, cv int) int {
	if cw == w || i == n {
		if cv > maxV {
			maxV = cv
		}
		return maxV
	}
	HandlerValueFun(i+1, cw, cv)
	if cw + weight[i] <= w {
		HandlerValueFun(i+1, cw+weight[i], cv+value[i])
	}
	return maxV
}

func Knapsack(weight []int, n, w int) int {
	var states = make([][]bool, n)
	for i := 0; i < n; i++ {
		column := make([]bool, w+1)
		states[i] = column
	}
	states[0][0] = true
	if weight[0] <= w {
		states[0][weight[0]] = true
	}
	for i := 1; i < n; i++ {
		for j := 0; j < w; j++ {
			if states[i-1][j] {
				states[i][j] = states[i-1][j]
			}
		}
		for j := 0; j <= w-weight[i]; j++ {
			if states[i-1][j] {
				states[i][j+weight[i]] = true
			}
		}
	}
	for i := w; i >= 0 ; i-- {
		if states[n-1][i] {
			return i
		}
	}
	return 0
}

func Knapsack2(weight []int, n, w int) int  {
	var states = make([]bool, w+1)
	states[0] = true
	if weight[0] <= w {
		states[weight[0]] =  true
	}
	for i := 1; i < n; i++ {
		for j := w-weight[i]; j >= 0; j-- {
			if states[j] == true {
				states[j+weight[i]] = true
			}
		}
	}
	for i := w; i >= 0 ; i-- {
		if states[i] {
			return i
		}
	}
	return 0
}

func Knapsack3(weight, value []int, n, w int) int {
	var states = make([][]int, n)
	for i := 0; i < n; i++ {
		column := make([]int, w+1)
		for j := 0; j < w+1; j++ {
			column[j] = -1
		}
		states[i] = column
	}
	states[0][0] = 0
	if weight[0] <= w {
		states[0][weight[0]] = value[0]
	}
	for i := 1; i < n; i++ {
		for j := 0; j <= w; j++ {
			if states[i-1][j] >= 0 {
				states[i][j] = states[i-1][j]
			}
		}
		for j := 0; j <= w - weight[i]; j++ {
			if states[i-1][j] >= 0 {
				v := states[i-1][j] + value[i]
				if v > states[i][j+weight[i]] {
					states[i][j+weight[i]] = v
				}
			}
		}
	}
	fmt.Println(states)
	//找出最大值
	maxValue := -1
	for i := 0; i <= w; i++ {
		if states[n-1][i] > maxValue {
			maxValue = states[n-1][i]
		}
	}
	return maxValue
}

func Knapsack4(weight, value []int, n, w int) int {
	var states = make([]int, w+1)
	states[0] = 0
	if weight[0] <= w {
		states[weight[0]] = value[0]
	}
	for i := 1; i < n; i++ {
		for j := w-weight[i]; j >= 0 ; j-- {
			if states[j] >= 0 {
				v := states[j] + value[i]
				if v > states[j+weight[i]] {
					states[j+weight[i]] = v
				}
			}
		}
	}
	fmt.Println(states)
	for i := w; i >= 0; i-- {
		if states[i] >= 0 {
			return states[i]
		}
	}
	return 0
}

