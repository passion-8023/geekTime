package main

func main() {
	//arrayStack := stack.NewArrayStack(3)
	//fmt.Println(arrayStack.Push("1"))
	//fmt.Println(arrayStack.Push("2"))
	//fmt.Println(arrayStack.Push("3"))
	//fmt.Println(arrayStack.Push("4"))
	//arrayStack.Pop()
	//arrayStack.Show()

}

func fib(n int) int {
	if n < 2 {
		return n
	}
	a, b := 0, 0
	sum := 1
	for i := 2; i < n; i++ {
		a = b
		b = sum
		sum = a + b
	}
	return sum
}

