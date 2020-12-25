/*
 *  @Author : huangzj
 *  @Time : 2020/12/21 11:18
 *  @Description：
 */

package simple

import "testing"

type TwoNum struct{}

func (TwoNum) Description(string) string {
	return `
		给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
		
		你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。

		示例:
		
		给定 nums = [2, 7, 11, 15], target = 9
		
		因为 nums[0] + nums[1] = 2 + 7 = 9
		所以返回 [0, 1]
	
`
}

//思路分析：map用来缓存比target小的数据, 比如说A + B =target.先遍历到A,则存储B -> A的下标，当找到B的时候，就可以把A的下标一起返回
func twoSum(nums []int, target int) []int {
	deal := make(map[int]int, 0)
	for i, row := range nums {
		if index, ok := deal[row]; ok {
			return []int{index, i}
		}
		deal[target-row] = i
	}

	return []int{}
}

func TestTwoSum(t *testing.T) {

}
