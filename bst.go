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

func printinorder(p *TreeNode) {

	if p == nil {
		return
	}

	printinorder(p.Left)
	fmt.Println(p.Val)
	printinorder(p.Right)
}

func topview(p *TreeNode) {

	if p == nil {
		return
	}

	printleft(p.Left)
	fmt.Println(p.Val)
	printright(p.Right)
}
func printleft(p *TreeNode) {

	if p == nil {
		return
	}

	printleft(p.Left)
	fmt.Println(p.Val)
}
func printright(p *TreeNode) {

	if p == nil {
		return
	}

	fmt.Println(p.Val)
	printright(p.Right)
}

func insertnode(Val int, p *TreeNode) *TreeNode {

	if p == nil {
		p = &TreeNode{Val, nil, nil}
	} else if Val > p.Val {
		p.Right = insertnode(Val, p.Right)
	} else if Val < p.Val {
		p.Left = insertnode(Val, p.Left)
	}
	return p
}
func treemin(p *TreeNode) *TreeNode {
	for p.Left != nil {
		p = p.Left
	}
	return p
}

func successor(Val int, p *TreeNode) *TreeNode {

	if p == nil {
		return nil
	}

	var succ *TreeNode = nil
	for p != nil {
		if Val > p.Val {
			p = p.Right
		} else if Val < p.Val {
			succ = p
			p = p.Left
		} else if Val == p.Val {
			if p.Right != nil {
				succ = treemin(p.Right)
			}
			break
		}
	}
	return succ
}
func isvalidbst(p *TreeNode) bool {

	lower := math.MinInt32
	upper := math.MaxInt32
	return checkbst(p, lower, upper)
}
func checkbst(p *TreeNode, lower int, upper int) bool {

	if p == nil {
		return true
	}

	if p.Val > lower && p.Val < upper {
		if checkbst(p.Left, lower, p.Val) && checkbst(p.Right, p.Val, upper) {
			return true
		}
	}
	return false
}
func lca(p *TreeNode, v1 int, v2 int) *TreeNode {

	if p == nil {
		return nil
	}

	for p != nil {
		if p.Val > v1 && p.Val > v2 && p.Left != nil {
			p = p.Left
		} else if p.Val < v1 && p.Val < v2 && p.Right != nil {
			p = p.Right
		} else if p.Val > v1 && p.Val < v2 {
			return p
		}
	}
	return nil
}
func levelorder(p *TreeNode) {
	n := height(p)
	for i := 1; i <= n; i += 1 {
		printlevel(p, i)
		fmt.Println()
	}
}
func printlevel(p *TreeNode, level int) {

	if p == nil {
		return
	}
	if level == 1 {
		fmt.Printf("%d\t", p.Val)
	} else if level > 1 {
		printlevel(p.Left, level-1)
		printlevel(p.Right, level-1)
	}
}
func height(p *TreeNode) int {

	if p == nil {
		return 0
	}
	left := height(p.Left)
	right := height(p.Right)

	if left > right {
		return 1 + left
	} else {
		return 1 + right
	}
	return 0
}

var maxPathVar = math.MinInt32

func maxsumpath(p *TreeNode) int {
	pathsum(p)
	return maxPathVar
}
func max(x int, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
func pathsum(p *TreeNode) int {
	if p == nil {
		return 0
	}

	left := max(0, pathsum(p.Left))
	right := max(0, pathsum(p.Right))

	maxPathVar = max(maxPathVar, (left + right + p.Val))
	return (max(left, right) + p.Val)
}
func main() {

	var treelist1 *TreeNode

	treefeed := []int{-10, 9, 20, 15, 7}

	for i := 0; i < len(treefeed); i++ {
		treelist1 = insertnode(treefeed[i], treelist1)
	}

	/*fmt.Println("Printing tree...")
	printinorder(treelist1)*/

	/*treelist1 = &TreeNode{8,nil,nil}
	treelist1.Left = &TreeNode{12,nil,nil}
	treelist1.Right = &TreeNode{13,nil,nil}
	fmt.Println(isvalidbst(treelist1))*/

	/*succ := successor(int(5), treelist1)
	fmt.Println("Successor of 5:", succ.Val)*/

	/*lcanode := lca(treelist1, 1, 6)
	fmt.Println("Lowest common ancestor: ", lcanode.Val)*/

	//topview(treelist1)

	//fmt.Println(height(treelist1))
	//levelorder(treelist1)

	fmt.Println(maxsumpath(treelist1))
}
