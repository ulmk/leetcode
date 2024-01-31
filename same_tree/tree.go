package main

import "fmt"

type Node[T comparable] struct {
	//parent   *Node[T]
	data     T
	children []*Node[T]
}

// func (n *Node[T]) GetParent() *Node[T] {
// 	return n.parent
// }

func (n *Node[T]) GetChildren() []*Node[T] {
	return n.children
}

func (n *Node[T]) AddChildren(children ...*Node[T]) {
	if n.children == nil {
		n.children = []*Node[T]{}
	}
	n.children = append(n.children, children[:]...)
}

func (n *Node[T]) ReplaceChildren(children ...*Node[T]) {
	n.children = []*Node[T]{}
	n.AddChildren(children...)
}

// func (n *Node[T]) SetParent(parent *Node[T]) {
// 	if parent == nil {
// 		return
// 	}
// 	n.parent = parent
// }

func (n *Node[T]) GetData() T {
	return n.data
}

func (n *Node[T]) SetData(data T) {
	n.data = data
}

func (n *Node[T]) FindNode(value T) *Node[T] {
	if n.data == value {
		return n
	}
	for _, child := range n.children {
		if found := child.FindNode(value); found != nil {
			return found
		}
	}
	return nil
}

func (n *Node[T]) RemoveNode(value T) bool {
	for i, child := range n.children {
		if child.data == value {
			n.children = append(n.children[:i], n.children[i+1:]...)
			return true
		}
		if child.RemoveNode(value) {
			return true
		}
	}
	return false
}

func (n *Node[T]) DeepCopy() *Node[T] {
	copyNode := &Node[T]{data: n.data}
	for _, child := range n.children {
		copyNode.AddChildren(child.DeepCopy())
	}
	return copyNode
}

func (n *Node[T]) BFS(target T) *Node[T] {

	queue := []*Node[T]{n}

	for len(queue) > 0 {
		curr := queue[0]
		fmt.Printf("curr BFS %v \n", curr.data)
		queue = queue[1:]

		if curr.data == target {
			//fmt.Printf("%v \n", curr.data)
			return curr
		}

		queue = append(queue, curr.children...)
		// if len(curr.children) > 0 {
		// 	for _, child := range curr.children {
		// 		queue = append(queue, child)
		// 	}
		// }
	}
	return nil
}

func (n *Node[T]) DFS(target T) *Node[T] {
	if n.data == target {
		return n
	}

	for _, child := range n.children {
		if found := child.DFS(target); found != nil {
			return found
		}
	}

	return nil
}

func (n *Node[T]) Equals(target *Node[T]) bool {
	if target == nil && n == nil {
		return true
	}

	if target == nil || n == nil {
		return false
	}

	if target.data != n.data || len(n.children) != len(target.children) {
		return false
	}

	for i := range n.children {
		if !n.children[i].Equals(target.children[i]) {
			return false
		}
	}
	return true
}

func (n *Node[T]) PrintTree() {
	fmt.Printf("{data:%+v children:[", n.data)
	for i, child := range n.children {
		if i != 0 {
			fmt.Printf(" ")
		}
		//fmt.Printf("data:%+v", child.GetData()
		child.PrintTree()
	}
	fmt.Printf("]}\n")
}

func (n *Node[T]) PrintRecTree() {
	n.printTreeRecursive(0)
}

func (n *Node[T]) printTreeRecursive(indentLevel int) {
	fmt.Printf("%s- %v\n", getIndentString(indentLevel), n.data)

	for _, child := range n.children {
		child.printTreeRecursive(indentLevel + 1)
	}
}

func getIndentString(indentLevel int) string {
	indent := ""
	for i := 0; i < indentLevel; i++ {
		indent += "  "
	}
	return indent
}

// type Tree[T any] struct {
// 	root Node[T]
// }
