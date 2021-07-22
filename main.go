package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recoverTree(root *TreeNode) {
	var prev *TreeNode

	// a, could be called first, since it is set once
	var a *TreeNode
	var b *TreeNode

	// use double pointers
	dfs(root, &prev, &a, &b)

	// this problem there is exactly one substition
	// the tree has exactly one error
	a.Val, b.Val = b.Val, a.Val
	return
}

// in order
func dfs(root *TreeNode, prev, a, b **TreeNode) {
	if root == nil {
		return
	}

	dfs(root.Left, prev, a, b)

	// Value at the address of prev is or *prev
	if *prev != nil {
		if root.Val <= (*prev).Val {
			// set the first, just once
			if *a == nil {
				*a = *prev
			}
			// the second pass, this will reset b
			*b = root
		}
	}

	// set the previous
	*prev = root

	dfs(root.Right, prev, a, b)

	return
}
