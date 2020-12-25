/*
 *  @Author : huangzj
 *  @Time : 2020/12/21 22:26
 *  @Description：
 */

package simple

import (
	"fmt"
	"testing"
)

type GroupAnagrams struct{}

func (GroupAnagrams) Description(string) string {
	return `
			给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。
			
			示例:
			
			输入: ["eat", "tea", "tan", "ate", "nat", "bat"]
			输出:
			[
			  ["ate","eat","tea"],
			  ["nat","tan"],
			  ["bat"]
			]
			说明：
			
			所有输入均为小写字母。
			不考虑答案输出的顺序。

`
}

//思路分析：定义26个质数作为字母的哈希，因为题目要求都是小写字母，并且质数相乘不会出现相同（三个不同的质数相乘的结果一定不一样）
//如果单词中包含的是是完全不同的字母。还可以使用二进制32位的26位来表示，性能更好
var prime = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103}

func groupAnagrams(strs []string) [][]string {
	m := make(map[int][]string, 0)
	for _, str := range strs {
		num := 1
		for _, b := range []byte(str) {
			num = num * prime[int(b-97)]
		}
		if _, ok := m[num]; ok {
			m[num] = append(m[num], str)
			continue
		}
		m[num] = []string{str}
	}

	list := make([][]string, 0)
	for _, value := range m {
		list = append(list, value)
	}
	return list
}

func TestGroupAnagrams(t *testing.T) {
	for _, row := range groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}) {
		fmt.Println(row)
	}
}
