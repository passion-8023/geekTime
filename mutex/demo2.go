package main

import (
	"fmt"
	"sync"
	"time"
)

//线程安全的计数器类型
type Counter2 struct {
	CounterType int
	Name string

	mu sync.Mutex
	count uint64
}

//加1的方法，内部使用互斥锁保护
func (c *Counter2) Incr()  {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

//得到计数器的值，也需要锁保护
func (c *Counter2) Count() uint64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

func main()  {
	now := time.Now()
	// 封装好的计数器
	var counter2 Counter2

	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 100000; j++ {
				counter2.Incr()
			}
		}()
	}
	wg.Wait()
	fmt.Println(counter2.Count())
	duration := time.Since(now)
	fmt.Println(duration)
}
