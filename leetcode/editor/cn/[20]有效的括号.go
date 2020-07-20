package cn
//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
//
// 有效字符串需满足： 
//
// 
// 左括号必须用相同类型的右括号闭合。 
// 左括号必须以正确的顺序闭合。 
// 
//
// 注意空字符串可被认为是有效字符串。 
//
// 示例 1: 
//
// 输入: "()"
//输出: true
// 
//
// 示例 2: 
//
// 输入: "()[]{}"
//输出: true
// 
//
// 示例 3: 
//
// 输入: "(]"
//输出: false
// 
//
// 示例 4: 
//
// 输入: "([)]"
//输出: false
// 
//
// 示例 5: 
//
// 输入: "{[]}"
//输出: true 
// Related Topics 栈 字符串


//leetcode submit region begin(Prohibit modification and deletion)
func isValid(s string) bool {
	if len(s)==0{
		return true
	}
	if len(s)&1==1{
		return false
	}
	var stack []int32
	for _,v:=range s{
		if v=='('||v=='['||v=='{'{
			stack=append(stack,v)
		}
		if v==')'{
			if len(stack)==0 {
				return false
			}
			temp:=stack[len(stack)-1]
			stack=stack[:len(stack)-1]
			if temp!='('{
				return false
			}

		}
		if v==']'{
			if len(stack)==0 {
				return false
			}
			temp:=stack[len(stack)-1]
			stack=stack[:len(stack)-1]
			if temp!='['{
				return false
			}

		}
		if v=='}'{
			if len(stack)==0 {
				return false
			}
			temp:=stack[len(stack)-1]
			stack=stack[:len(stack)-1]
			if temp!='{'{
				return false
			}
		}
	}
	if len(stack)==0{
		return true
	}
	return false
}
//leetcode submit region end(Prohibit modification and deletion)
