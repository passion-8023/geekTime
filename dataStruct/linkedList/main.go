package main

import (
	"fmt"
)

type Node struct {
	Value interface{}
	Next *Node
}

type List struct {
	Head *Node
	Length int
}

type Method interface {
	Insert(i int, v interface{})
	Delete(i int)
	GetLength() int
	Search(v interface{}) int
	isNull() bool
}

func NewNode(v interface{}) *Node {
	return &Node{
		Value: v,
		Next:  nil,
	}
}

func NewList() *List {
	return &List{
		Head:   NewNode(nil),
		Length: 0,
	}
}

func (l *List) Insert(i int, v interface{}) {
	node := NewNode(v)
	pre := l.Head
	for count := 0; count < i; count++ {
		if count == i-1 {
			node.Next = pre.Next
			pre.Next = node
			l.Length++
		}
		pre = pre.Next
	}
}

func (l *List) Delete(i int) {
	pre := l.Head
	for count := 0; count <= i-1; count++ {
		s := pre.Next
		if count == i-1 {
			pre.Next = s.Next
			l.Length--
		}
		pre = pre.Next
	}
}

func (l *List) GetLength() int {
	return l.Length
}

func (l *List) Search(v interface{}) int {
	pre := l.Head.Next
	for i := 0; i <= l.Length; i++ {
		if pre.Value == v {
			return i
		}
		pre = pre.Next
	}
	return 0
}

func (l *List) isNull() bool {
	pre := l.Head.Next
	if pre == nil {
		return true
	}
	return false
}

func PrintList(l *List)  {
	pre := l.Head.Next
	fmt.Println("list shows as follows: ...")
	for i := 1; i <= l.Length; i++ {
		fmt.Printf("%v\n", pre.Value)
		pre = pre.Next
	}
}

//翻转链表
func reversalList(l *List) *List {
	pre := l.Head.Next
	newlist := NewList()
	for pre != nil {
		next := pre.Next
		pre.Next = newlist.Head.Next
		newlist.Head.Next = pre
		newlist.Length++
		pre = next
	}
	return newlist
}


func reverseKGroup(head *Node, k int) *Node {
	newNode := &Node{
		Value: -1,
		Next:  head,
	}
	pre := newNode
	end := newNode
	for end.Next != nil {
		for i := 0; i < k && end != nil; i++ {
			end = end.Next
		}
		if end == nil {
			break
		}
		start := pre.Next
		next := end.Next
		end.Next = nil
		pre.Next = reversalNode(start)
		start.Next = next
		pre = start
		end = pre
	}
	return newNode.Next
}

func reversalNode(node *Node) *Node {
	cur := node
	//定义空节点或空链表
	newNode := &Node{}
	for cur != nil {
		next := cur.Next
		cur.Next = newNode
		newNode = cur
		cur = next
	}
	return newNode
}


//链表是否有环
func hasCycle(head *Node) bool {
	tmpMap := map[*Node]struct{}{}
	for head != nil {
		if _, ok := tmpMap[head]; ok {
			return true
		}
		tmpMap[head] = struct{}{}
		head = head.Next
	}
	return false
}

func hasCycle2(head *Node) bool {
	//快慢指针
	if head == nil || head.Next == nil {
		return false
	}
	low, fast := head, head.Next
	for low != fast{
		if fast == nil || fast.Next == nil {
			return false
		}
		low = low.Next
		fast = fast.Next.Next
	}
	return true
}

func detectCycle(head *Node) *Node {
	tmpMap := map[*Node]struct{}{}
	for head != nil {
		if _, ok := tmpMap[head]; ok {
			return head
		}
		tmpMap[head] = struct{}{}
		head = head.Next
	}
	return nil
}

func detectCycle2(head *Node) *Node {
	//快慢指针
	low, fast := head, head
	for fast != nil {
		low = low.Next
		if fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		if fast == low {
			p := head
			for p != low {
				p = p.Next
				low = low.Next
			}
			return p
		}
	}
	return nil
}

//合并两个有序链表
func mergeTwoLists(l1 *Node, l2 *Node) *Node {
	newList := &Node{}
	cur := newList
	for l1 != nil && l2 != nil {
		if fmt.Sprintf("%d", l1.Value) < fmt.Sprintf("%d", l2.Value) {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}
	if l1 != nil {
		cur.Next = l1
	} else {
		cur.Next = l2
	}
	return newList
}


func main() {
	list := NewList()

	var M Method
	M = list

	fmt.Println("list is null", M.isNull())
	if M.isNull() {
		M.Insert(1, 1)
		M.Insert(2, "wangxin")
		M.Insert(3, 26)
		M.Insert(4, "phper")
		M.Insert(5, "北京")
		M.Insert(6, "甘肃")
		M.Insert(7, "微梦")
	}
	fmt.Println("length of list is", M.GetLength())

	PrintList(list)

	//newlist := reversalList(list)
	//PrintList(newlist)

	fmt.Println("=========")
	node := reverseKGroup(list.Head.Next, 2)
	cur := node
	for cur != nil {
		fmt.Println(cur.Value)
		cur = cur.Next
	}
}
