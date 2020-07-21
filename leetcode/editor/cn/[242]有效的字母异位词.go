package cn

//给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
//
// 示例 1: 
//
// 输入: s = "anagram", t = "nagaram"
//输出: true
// 
//
// 示例 2: 
//
// 输入: s = "rat", t = "car"
//输出: false 
//
// 说明: 
//你可以假设字符串只包含小写字母。 
//
// 进阶: 
//如果输入字符串包含 unicode 字符怎么办？你能否调整你的解法来应对这种情况？ 
// Related Topics 排序 哈希表


//leetcode submit region begin(Prohibit modification and deletion)
func isAnagram(s string, t string) bool {
	//smap:=make(map[int32]int,0)
	//tmap:=make(map[int32]int,0)
	//
	//for _,v:=range s{
	//	smap[v]=smap[v]+1
	//}
	//for _,v:=range t{
	//	tmap[v]=tmap[v]+1
	//}
	//if len(smap)!=len(tmap) {
	//	return false
	//}else {
	//	for k,v:=range smap{
	//		value,ok:=tmap[k]
	//		if !ok||value!=v {
	//			return false
	//		}
	//	}
	//}
	sArray:=[26]int{}
	tArray:=[26]int{}
	for _,v:=range s{
		sArray[v-'a']=sArray[v-'a']+1
	}
	for _,v:=range t{
		tArray[v-'a']=tArray[v-'a']+1
	}
	return sArray==tArray
}
//leetcode submit region end(Prohibit modification and deletion)
