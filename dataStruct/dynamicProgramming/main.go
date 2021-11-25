package main

import (
	"fmt"
	"geekTime/dataStruct/dynamicProgramming/zeroOnePackage"
)

func main() {
	//fmt.Println(zeroOnePackage.HandlerFun(0, 0))
	//fmt.Println(zeroOnePackage.HandlerValueFun(0, 0, 0))
	var (
		weight []int = []int{2, 2, 4, 6, 3}
		value []int = []int{3, 4, 8, 9, 6}
		n int = 5
		w int = 9
	)

	//fmt.Println(zeroOnePackage.Knapsack2(weight, n, w))
	//fmt.Println(zeroOnePackage.Knapsack3(weight, value, n, w))
	fmt.Println(zeroOnePackage.Knapsack4(weight, value, n, w))

}

