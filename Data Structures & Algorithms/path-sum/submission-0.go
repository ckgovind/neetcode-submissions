/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func hasPathSum(root *TreeNode, targetSum int) bool {

	if root == nil {
		return false
	}

	if root.Right == nil && root.Left == nil {
		return targetSum-root.Val == 0
	}

	if hasPathSum(root.Right, targetSum-root.Val) {
		return true
	}

	if hasPathSum(root.Left, targetSum-root.Val) {
		return true
	}

	return false

}
