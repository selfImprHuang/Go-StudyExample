/*
 *  @Author : huangzj
 *  @Time : 2020/12/21 23:01
 *  @Description：
 */

package simple

import (
	"fmt"
	"testing"
)

type LongestCommonPrefix struct{}

func (LongestCommonPrefix) Description(string) string {
	return `
		编写一个函数来查找字符串数组中的最长公共前缀。
		
		如果不存在公共前缀，返回空字符串 ""。
		
		示例 1:
		
		输入: ["flower","flow","flight"]
		输出: "fl"
		示例 2:
		
		输入: ["dog","racecar","car"]
		输出: ""
		解释: 输入不存在公共前缀。
`
}

//思路分析：最长前缀其实就是求最长字符串长度.所以第一个字符串作为公共字符串，然后循环所有的字符串，进行公共字符串的缩减
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	var match []byte
	for index, str := range strs {
		if index == 0 {
			match = []byte(str)
			continue
		}
		position := getPosition(match, str)
		match = match[0:position]
	}

	return string(match)
}

func getPosition(match []byte, str string) int {
	var position int
	for i, m := range []byte(str) {
		if len(match)-1 < i {
			break
		}
		if m == match[i] {
			position++
			continue
		}
		break
	}

	return position
}

func TestLongestCommonPrefix(t *testing.T) {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Println(longestCommonPrefix([]string{"flow", "flower", "flight"}))
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))
	fmt.Println(longestCommonPrefix([]string{"cir", "car"}))
}
