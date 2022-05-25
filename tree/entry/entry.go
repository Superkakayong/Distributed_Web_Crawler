package main

import "Distributed_Web_Crawler/tree"

func main() {
	var root tree.Node

	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)

	//nodes := []treeNode{
	//	{value: 3},
	//	{},
	//	{6, nil, &root},
	//}
	//
	//fmt.Println(nodes)

	root.Print()

	root.Right.Left.SetValue(4)
	root.Right.Left.Print()
}
