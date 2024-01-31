package main

import "fmt"

func main() {
	fmt.Println("Same Tree ... ")

	root := &Node[int]{data: 8}
	node1 := &Node[int]{data: 1}
	node2 := &Node[int]{data: 2}
	node3 := &Node[int]{data: 3}
	node4 := &Node[int]{data: 4}
	node5 := &Node[int]{data: 5}
	node6 := &Node[int]{data: 6}
	node7 := &Node[int]{data: 7}
	root.AddChildren(node1, node2, node3)
	node3.AddChildren(node4)
	node2.AddChildren(node5)
	node5.AddChildren(node6, node7)

	root.PrintTree()
	//rootnode.PrintRecTree()
	targetBFS := root.BFS(7)
	fmt.Printf("BFS res: %+v\n", targetBFS)

	targetDFS := root.DFS(3)
	fmt.Printf("DFS res: %+v\n", targetDFS)

	tree1 := &Node[int]{data: 1}
	child1 := &Node[int]{data: 2}
	child2 := &Node[int]{data: 3}
	child3 := &Node[int]{data: 4}
	grandchild := &Node[int]{data: 5}

	tree1.children = []*Node[int]{child1, child2, child3}
	child2.children = []*Node[int]{grandchild}

	tree2 := &Node[int]{data: 1}
	child1 = &Node[int]{data: 2}
	child2 = &Node[int]{data: 3}
	child3 = &Node[int]{data: 4}
	grandchild = &Node[int]{data: 7}

	tree2.children = []*Node[int]{child1, child2, child3}
	child2.children = []*Node[int]{grandchild}

	// Compare the trees
	equal := tree1.Equals(tree2)
	fmt.Printf("Are the trees equal? %t\n", equal)

}
