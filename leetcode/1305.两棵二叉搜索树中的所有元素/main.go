package main

import "log"

/**
给你 root1 和 root2 这两棵二叉搜索树。

请你返回一个列表，其中包含 两棵树 中的所有整数并按 升序 排序。.

 

示例 1：



输入：root1 = [2,1,4], root2 = [1,0,3]
输出：[0,1,1,2,3,4]
示例 2：

输入：root1 = [0,-10,10], root2 = [5,1,7,0,2]
输出：[-10,0,0,1,2,5,7,10]
示例 3：

输入：root1 = [], root2 = [5,1,7,0,2]
输出：[0,1,2,5,7]
示例 4：

输入：root1 = [0,-10,10], root2 = []
输出：[-10,0,10]
示例 5：



输入：root1 = [1,null,8], root2 = [8,1]
输出：[1,1,8,8]
 

提示：

每棵树最多有 5000 个节点。
每个节点的值在 [-10^5, 10^5] 之间。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/all-elements-in-two-binary-search-trees
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
* Definition for a binary tree node.
* type TreeNode struct {
*     Val int
*     Left *TreeNode
*     Right *TreeNode
* }
*/
func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	s1 := make([]int, 0)
	s2 := make([]int, 0)
	inorderTraversal(root1, &s1)
	inorderTraversal(root2, &s2)
	res := make([]int, len(s1)+len(s2))
	var i, j int
	for i < len(s1) || j < len(s2) {
		if i == len(s1) || (j < len(s2) && s1[i] > s2[j]) {
			res[i+j] = s2[j]
			j++
		} else {
			res[i+j] = s1[i]
			i++
		}
	}
	return res
}

//二叉树的中序遍历
func inorderTraversal(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}

	inorderTraversal(root.Left, res)
	*res = append(*res, root.Val)
	inorderTraversal(root.Right, res)

}

func main() {
	//root := &TreeNode{Val: 1, Left: &TreeNode{Val: 0, Left: nil, Right: nil}, Right: &TreeNode{Val: 2, Left: nil, Right: nil}}
	//res := make([]int, 0)
	//inorderTraversal(root, &res)
	//log.Println(res)

	list := []int{1}

	list = list[1:]
	log.Println(list)
}
