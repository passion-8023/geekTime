package main

import (
	"fmt"
)

func main() {
	//arr := [5]int{1, 2, 3, 4, 5}
	//for i, v := range arr {
	//	if i == len(arr) - 1 {
	//		arr[0] = v + arr[0]
	//	} else {
	//		arr[i+1] += v
	//	}
	//}
	//fmt.Println(arr)  //[6, 3, 5, 7, 9]
	//
	//slice := []int{1, 2, 3, 4, 5}
	//for i, v := range slice {
	//	if i == len(slice) - 1 {
	//		slice[0] += v
	//	} else {
	//		slice[i+1] += v
	//	}
	//}
	//fmt.Println(slice) //[3, 5, 7, 9, 6]

	//var d demo
	//fmt.Println(unsafe.Sizeof(d))
	//fmt.Println(unsafe.Offsetof(d.b))
	//
	//var d2 demo2
	//fmt.Println(unsafe.Sizeof(d2))
	//
	//var d3 struct{}
	//fmt.Println(unsafe.Sizeof(d3))
	//
	var sliceDemo []int
	sliceDemo = []int{1, 2, 3, 4, 5, 6}
	fmt.Println(sliceDemo)
	var a = (*[5]int)(sliceDemo)

	fmt.Printf("a type %T\n", a)
	fmt.Println(&sliceDemo[1])
	fmt.Println(&a[1])

}

type demo struct { //1+7 + 8 + 2+6
	 b byte // 1字节
	 i int64 // 8字节
 	 u uint16 // 2自己
}

type demo2 struct { //1+1 + 2 + 4 + 8
	b byte //1
	u uint16 //2
	i int64 //8
}