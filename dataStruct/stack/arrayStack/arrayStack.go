package arrayStack

import "fmt"

type ArrayStack struct {
	items []string
	count int
	n int
}

type Method interface {
	Show()
	Push(str string) bool
	Pop() string
}

func NewArrayStack(n int) *ArrayStack {
	return &ArrayStack{
		items: make([]string, n),
		count: 0,
		n: n,
	}
}

func (a *ArrayStack) Show()  {
	fmt.Println(a.items)
}

func (a *ArrayStack) Push(str string) bool {
	if a.count == a.n {
		return false
	}
	a.items[a.count] = str
	a.count++
	return true
}

func (a *ArrayStack) Pop() string {
	if a.count == 0 {
		return ""
	}
	tmp := a.items[a.count-1]
	a.count--
	a.items = a.items[:a.count]
	return tmp
}







