package main

import "fmt"

type TreeNode1 struct {
	Val   int
	Left  *TreeNode1
	Right *TreeNode1
}

func invertTree(root *TreeNode1) *TreeNode1 {
	if root != nil {
		temp := root.Left
		root.Left = invertTree(root.Right)
		root.Right = invertTree(temp)
	}
	return root
}

func main() {
	root := &TreeNode1{
		4,
		&TreeNode1{
			2,
			&TreeNode1{
				Val:   1,
				Left:  nil,
				Right: nil,
			},
			&TreeNode1{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
		},
		&TreeNode1{
			7,
			&TreeNode1{
				Val:   6,
				Left:  nil,
				Right: nil,
			},
			&TreeNode1{
				Val:   9,
				Left:  nil,
				Right: nil,
			},
		},
	}
	invertTree(root)
	fmt.Println(111)
}
