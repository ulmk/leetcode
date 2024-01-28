package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) Add(val int) {
	newNode := &ListNode{Val: val, Next: nil}
	curr := l
	for curr.Next != nil {
		curr = curr.Next
	}
	curr.Next = newNode
	fmt.Println(curr)

}

func (l *ListNode) Print() {
	curr := l
	for curr != nil {
		fmt.Printf("%d -> ", curr.Val)
		curr = curr.Next
	}
	fmt.Println("nill")
}

func main() {
	fmt.Println("sLL . .. .")
	l1 := &ListNode{}

	l1.Add(5) // 5
	l1.Add(2) // 2
	l1.Add(5) // 3

	l1.Print()

	l2 := &ListNode{}

	l2.Add(8) // 7
	l2.Add(5) // 5
	l2.Add(7) // 8

	l2.Print()

	result := add(l1, l2)
	result.Print()

}

func add(l1 *ListNode, l2 *ListNode) *ListNode {
	sum := 0
	carry := 0
	lres := &ListNode{}
	curr := lres

	for l1 != nil || l2 != nil {
		sum = 0

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		sum += carry
		carry = sum / 10
		sum %= 10

		lres.Next = &ListNode{Val: sum}
		lres = lres.Next

	}

	if carry != 0 {
		lres.Next = &ListNode{Val: carry}
	}
	return curr.Next
}
