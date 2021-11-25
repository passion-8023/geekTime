package main

import (
	"fmt"
	"geekTime/mutex/demo5/relock"
)

func main() {
	r := &relock.RecursiveMutex{}
	StartLayer(r)
}

func StartLayer(r *relock.RecursiveMutex) {
	r.Lock()
	fmt.Println("开始")
	TwoLayer(r)
	r.Unlock()
}

func TwoLayer(r *relock.RecursiveMutex) {
	r.Lock()
	fmt.Println("进入第二层")
	ThreeLayer(r)
	r.Unlock()
}

func ThreeLayer(r *relock.RecursiveMutex) {
	r.Lock()
	fmt.Println("最后一层")
	r.Unlock()
}
