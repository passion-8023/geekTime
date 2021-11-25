package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	sync.Mutex
	Count int
}

func main() {
	var c Counter
	c.Lock()
	defer c.Unlock()
	//函数传参 复制了已经使用的锁
	foo(c)
}

func foo(c Counter)  {
	c.Lock()
	defer c.Unlock()
	fmt.Println("in foo")
}
