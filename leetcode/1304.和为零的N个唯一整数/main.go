package main

/**
给你一个整数 n，请你返回 任意 一个由 n 个 各不相同 的整数组成的数组，并且这 n 个数相加和为 0 。

示例 1：

输入：n = 5
输出：[-7,-1,1,3,4]
解释：这些数组也是正确的 [-5,-1,1,2,3]，[-3,-1,2,-2,4]。
示例 2：

输入：n = 3
输出：[-1,0,1]
示例 3：

输入：n = 1
输出：[0]
 

提示：

1 <= n <= 1000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/find-n-unique-integers-sum-up-to-zero
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func sumZero(n int) []int {
	//res := make([]int,n)
	//index := 0
	//if n % 2 != 0 {
	//	res[index] = 0
	//	n--
	//	index++
	//}
	//for i := 1; i <= n / 2; i++ {
	//	res[index] = i
	//	res[index+1] = -i
	//	index += 2
	//}
	//return res

	//思路2
	res := make([]int,n)
	var sum int
	for i:= 0; i < n-1; i++ {
		res[i] = i
		sum += i
	}
	res[n-1] = -sum
	return res
}