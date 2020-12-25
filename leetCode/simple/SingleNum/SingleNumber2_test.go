/*
 *  @Author : huangzj
 *  @Time : 2020/12/24 9:30
 *  @Description：
 */

package SingleNum

import (
	"fmt"
	"testing"
)

type SingleNumber2 struct{}

func (SingleNumber2) Description(string) string {
	return `
		给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现了三次。找出那个只出现了一次的元素。

		说明：
		
		你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？
		
		示例 1:
		
		输入: [2,2,3,2]
		输出: 3
		示例2:
		
		输入: [0,1,0,1,0,1,99]
		输出: 99
	`
}

//思路分析：最常规的方法是使用map统计或者是set去重，这边记录另外两种解法

//如果所有数字都出现了 3 次，那么每一列的 1 的个数就一定是 3 的倍数。之所以有的列不是 3 的倍数，就是因为只出现了 1 次的数贡献出了 1。所以所有不是 3 的倍数的列写 1，其他列写 0 ，就找到了这个出现 1 次的数。
//go这边要取64位，因为我的机器是64位的，如果我取32位就没办法得到正确的答案
func singleNumber2(nums []int) int {
	var result int
	//遍历每一位。这边只考虑64位的情况
	for i := 0; i < 64; i++ {
		var res int
		//遍历每一个数字
		for _, num := range nums {
			//对每个数字进行移位，然后 & 上1就可以得到移位后的最后一位是否为1的结果。进行该位置的1的个数的统计
			res += (num >> i) & 1
		}
		//该位置的1的个数%3，余数就是该位置是否为1的判断，然后再移动原来的位置，通过 | 来得到最终的结果
		result = result | ((res % 3) << i)
	}
	return result
}

//思路分析：
//func singleNumber3(nums []int) int {
//
//}

func TestSingleNumber2(t *testing.T) {
	fmt.Println(singleNumber2([]int{2, 2, 3, 2}))
	fmt.Println(singleNumber2([]int{0, 1, 0, 1, 0, 1, 99}))
	fmt.Println(singleNumber2([]int{-2, -2, 1, 1, 4, 1, 4, 4, -4, -2}))
}
