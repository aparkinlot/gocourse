package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	res := math.MinInt
	var propogate func(*TreeNode) int
	//fmt.Println(root.Val)
	propogate = func(rootDFS *TreeNode) int {
		if rootDFS == nil {
			return 0
		}
		lsum := max(0, propogate(rootDFS.Left))
		rsum := max(0, propogate(rootDFS.Right))

		res = max(res, rootDFS.Val+lsum+rsum)
		//fmt.Println(res)
		//fmt.Println(rootDFS.Val + max(lsum, rsum))
		return rootDFS.Val + max(lsum, rsum)
	}

	_ = propogate(root)
	return res
}

func main() {
	root := &TreeNode{
		Val: -10,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	res := maxPathSum(root)
	fmt.Println("ANS:", res)
	root = &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 11,
				Left: &TreeNode{
					Val: 7,
				},
				Right: &TreeNode{
					Val: 2,
				},
			},
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val: 13,
			},
			Right: &TreeNode{
				Val: 4,
				Right: &TreeNode{
					Val: 1,
				},
			},
		},
	}
	res = maxPathSum(root)
	fmt.Println("ANS:", res)
}
