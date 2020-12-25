/*
 *  @Author : huangzj
 *  @Time : 2020/12/22 14:23
 *  @Description：
 */

package simple

import (
	"fmt"
	"testing"
)

type FindDisappearedNumbers struct{}

func (FindDisappearedNumbers) Description(string) string {
	return `
			给定一个范围在  1 ≤ a[i] ≤ n ( n = 数组大小 ) 的 整型数组，数组中的元素一些出现了两次，另一些只出现一次。
			
			找到所有在 [1, n] 范围之间没有出现在数组中的数字。
			
			您能在不使用额外空间且时间复杂度为O(n)的情况下完成这个任务吗? 你可以假定返回的数组不算在额外空间内。
			
			示例:
			
			输入:
			[4,3,2,7,8,2,3,1]
			
			输出:
			[5,6]

		`
}

//思路分析：
//     *
//     * 找出 1 - n 中没有出现的数字。不能使用额外的空间，两次循环时间复杂度为 2O(n)，即为 O(n)。
//     *
//     * 解题思路：使用数组的下标来标记数字的出现于否，通过一遍遍历即可标记出全部已经出现的数组
//     *
//     * [4,3,2,7,8,2,3,1] 初始数据
//     *
//     * [4,3,2,-7,8,2,3,1] 第一个数据 4 出现，将数组的第四个也就是下标 3 的数据修改为负数。-7 计算时，通过绝对值处理一下即可不影响数据的计算
//     * [4,3,-2,-7,8,2,3,1]
//     * [4,-3,-2,-7,8,2,3,1]
//     * [4,-3,-2,-7,8,2,-3,1]
//     * [4,-3,-2,-7,8,2,-3,-1]
//     * [4,-3,-2,-7,8,2,-3,-1]
//     * [4,-3,-2,-7,8,2,-3,-1]
//     * [-4,-3,-2,-7,8,2,-3,-1]
//     *
//     * 计算结束，数组的第五个，第六个依然为整数，证明 5,6 没有出现
//     *
func findDisappearedNumbers(nums []int) []int {
	list := make([]int, 0)
	for _, num := range nums {
		nums[toPositive(num)-1] = -toPositive(nums[toPositive(num)-1])
	}

	for index, num := range nums {
		if num >= 1 {
			list = append(list, index+1)
		}
	}

	return list
}

func toPositive(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func TestFindDisappearedNumbers(t *testing.T) {
	fmt.Println(findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1}))
}
