package main

import (
	"fmt"
	"geekTime/mutex/demo6/getwaiter"
	"geekTime/mutex/demo6/trylock"
	"math/rand"
	"time"
)

func main() {
	//try()
	count()
}

func try() {
	var mu trylock.Mutex
	go func() {
		mu.Lock()
		time.Sleep(time.Duration(rand.Intn(2))*time.Second)
		mu.Unlock()
	}()

	time.Sleep(time.Second)

	ok := mu.TryLock()
	if ok {
		fmt.Println("got the lock")
		mu.Unlock()
		return
	}

	fmt.Println("can`t get the lock")
}


func count() {
	var mu getwaiter.Mutex
	var count int = 0
	for i := 0; i < 1000; i++ { // 启动1000个goroutine
		go func() {
			mu.Lock()
			time.Sleep(time.Second)
			count++
			mu.Unlock()
		}()
	}

	time.Sleep(time.Second)
	// 输出锁的信息
	fmt.Println(count)
	fmt.Printf("waitings: %d, isLocked: %t, woken: %t,  starving: %t\n", mu.Count(), mu.IsLocked(), mu.IsWoken(), mu.IsStarving())
}
