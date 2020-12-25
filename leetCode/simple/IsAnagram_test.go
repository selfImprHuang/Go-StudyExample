/*
 *  @Author : huangzj
 *  @Time : 2020/12/22 11:53
 *  @Description：
 */

package simple

import (
	"fmt"
	"testing"
)

type IsAnagram struct{}

func (IsAnagram) Description(string) string {
	return `
	给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。

		示例 1:
		
		输入: s = "anagram", t = "nagaram"
		输出: true
		示例 2:
		
		输入: s = "rat", t = "car"
		输出: false
		说明:
		你可以假设字符串只包含小写字母。
		
		进阶:
		如果输入字符串包含 unicode 字符怎么办？你能否调整你的解法来应对这种情况？
`
}

//思路分析:因为这边只提到了26个字母.所以声明长度为26的数组,在字符串t中出现的字符位 + 1 ,在字符串s中出现的字符位 -1，达到抵消的效果，最后如果数组都为0，就是异位数，否则不是
//这个题目还可以使用排序、前缀树   map 等方法处理
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	num := [26]int{}

	sb := []byte(s)
	tb := []byte(t)
	for i := 0; i < len(s); i++ {
		num[sb[i]-'a']++
		num[tb[i]-'a']--
	}

	for _, r := range num {
		if r != 0 {
			return false
		}
	}

	return true
}

func TestIsAnagram(t *testing.T) {
	fmt.Println(isAnagram("axmsdawd", "dwsadmxa"))
	fmt.Println(isAnagram("axmsdawd", "dwsadmxb"))
}
