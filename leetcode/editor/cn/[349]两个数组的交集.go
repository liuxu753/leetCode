package cn
//给定两个数组，编写一个函数来计算它们的交集。
//
// 
//
// 示例 1： 
//
// 输入：nums1 = [1,2,2,1], nums2 = [2,2]
//输出：[2]
// 
//
// 示例 2： 
//
// 输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
//输出：[9,4] 
//
// 
//
// 说明： 
//
// 
// 输出结果中的每个元素一定是唯一的。 
// 我们可以不考虑输出结果的顺序。 
// 
// Related Topics 排序 哈希表 双指针 二分查找


//leetcode submit region begin(Prohibit modification and deletion)
func intersection(nums1 []int, nums2 []int) []int {
	var res []int
	map1:=make(map[int]int)
	map2:=make(map[int]int)
	for _,v:=range nums1{
		map1[v]=1
	}
	for _,v:=range nums2{
		map2[v]=1
	}
	for k,_:=range map1{
		_,ok:=map2[k]
		if ok {
			res=append(res,k)
		}
	}
	return	res
}
//leetcode submit region end(Prohibit modification and deletion)
