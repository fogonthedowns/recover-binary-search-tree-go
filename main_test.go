package main

import (
	"testing"
)

func TestFun(t *testing.T) {
	a := &TreeNode{Val: 2}
	b := &TreeNode{Val: 4, Left: a}
	c := &TreeNode{Val: 1}
	root := &TreeNode{Val: 3, Left: c, Right: b}
	recoverTree(root)
	want := 2
	if root.Val != want {
		t.Errorf("got %d, want %d", root.Val, want)
	}
}
