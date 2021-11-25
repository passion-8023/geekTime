package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	//互斥锁保护计数器
	//var mu sync.Mutex
	count := 0

	//辅助变量，用来确定所有的goroutine都完成
	//var wg sync.WaitGroup
	//wg.Add(10)

	for i := 0; i < 10; i++ {
		//go func() {
		//	defer wg.Done()
			//对变量count执行10次加1
			for j := 0; j < 100000; j++ {
				//mu.Lock()
				count++
				//mu.Unlock()
			}
		//}()
	}

	//等待10个goroutine完成
	//wg.Wait()
	fmt.Println(count)
	duration := time.Since(now)
	fmt.Println(duration)
}
