/*
 *  @Author : huangzj
 *  @Time : 2020/12/24 15:18
 *  @Description：
 */

package simple

import (
	"fmt"
	"testing"
)

type MajorityElement1 struct{}

func (MajorityElement1) Description(string) string {
	return `
		给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。

		你可以假设数组是非空的，并且给定的数组总是存在多数元素。

		示例 1:	
		输入: [3,2,3]
		输出: 3

		示例 2:	
		输入: [2,2,1,1,1,2,2]
		输出: 2
		`
}

//思路说明：
func majorityElement1(nums []int) int {
	var result, count int
	for _, num := range nums {
		if count == 0 {
			result = num
			count++
			continue
		}
		if num == result {
			count++
			continue
		}
		count--
	}
	return result
}

func TestMajorityElement1(t *testing.T) {
	fmt.Println(majorityElement1([]int{3, 2, 3}))
	fmt.Println(majorityElement1([]int{2, 2, 1, 1, 1, 2, 2}))
	fmt.Println(majorityElement1([]int{3, 3, 4}))
}
