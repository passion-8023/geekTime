package main

import (
	"fmt"
	"sync"
)

const (
	mutexLocked = 1 << iota //持有锁的标记
	mutexWoken	//唤醒标记
	mutexStarving //饥饿标记
	mutexWaiterShift = iota //阻塞等待waiter数量
)

func main()  {
	fmt.Printf("持有锁的标记：%v，唤醒标记：%v，饥饿标记：%v，阻塞等待的waiter数量：%v\n", mutexLocked, mutexWoken, mutexStarving, mutexWaiterShift)
	old := 1
	demo := old&(mutexLocked|mutexStarving)
	//001 & (001|100) = 001&101 = 001
	fmt.Printf("demo is %v\n", demo)
	var mu sync.Mutex
	defer mu.Unlock()
	new := 1<<mutexWaiterShift
	fmt.Printf("new is %b\n", new)
}
