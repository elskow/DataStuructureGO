package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	memo := make(map[int][]int)
	depth := 0

	var dfs func(node *TreeNode, depth int)
	dfs = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}
		memo[depth] = append(memo[depth], node.Val)
		dfs(node.Left, depth+1)
		dfs(node.Right, depth+1)
	}
	dfs(root, depth)

	var res [][]int
	for _, v := range memo {
		res = append(res, v)
	}
	for i := 0; i < len(res); i++ {
		if i%2 == 1 {
			for j := 0; j < len(res[i])/2; j++ {
				res[i][j], res[i][len(res[i])-j-1] = res[i][len(res[i])-j-1], res[i][j]
			}
		}
	}
	return res
}

func main() {
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 20}
	root.Right.Left = &TreeNode{Val: 15}
	root.Right.Right = &TreeNode{Val: 7}
	for _, v := range zigzagLevelOrder(root) {
		for _, v2 := range v {
			print(v2, " ")
		}
		println()
	}
}
