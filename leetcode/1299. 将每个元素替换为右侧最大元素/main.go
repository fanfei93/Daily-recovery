package main

/**
给你一个数组 arr ，请你将每个元素用它右边最大的元素替换，如果是最后一个元素，用 -1 替换。

完成所有替换操作后，请你返回这个数组。

示例：

输入：arr = [17,18,5,4,6,1]
输出：[18,6,6,6,1,-1]
 

提示：

1 <= arr.length <= 10^4
1 <= arr[i] <= 10^5

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/replace-elements-with-greatest-element-on-right-side
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func replaceElements(arr []int) []int {
	//不改变原切片
	//res := make([]int,len(arr))
	//res[len(arr)-1] = -1
	//max := arr[len(arr)-1]
	//
	//for i := len(arr)-2; i >= 0; i-- {
	//	res[i] = max
	//	if arr[i] > max {
	//		max = arr[i]
	//	}
	//}
	//return res



	//改变原切片 空间复杂度更低
	max := arr[len(arr)-1]
	arr[len(arr)-1] = -1
	for i := len(arr)-2; i >= 0; i-- {
		arr[i],max = max,getMax(arr[i],max)
	}
	return arr
}

func getMax(a,b int) int {
	if a > b {
		return a
	}
	return b
}

//两种方法提交时耗时相同，空间复杂度也基本一致