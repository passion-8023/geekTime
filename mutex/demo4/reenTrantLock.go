package main

import (
	"fmt"
	"sync"
)

func main() {
	l := &sync.Mutex{}
	foo(l)
}

func foo(l sync.Locker)  {
	fmt.Println("in foo")
	l.Lock()
	car(l)
	l.Unlock()
}

func car(l sync.Locker)  {
	l.Lock()
	fmt.Println("in car")
	l.Unlock()
}
