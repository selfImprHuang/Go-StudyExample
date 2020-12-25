/*
 *  @Author : huangzj
 *  @Time : 2020/12/23 16:37
 *  @Description：
 */

package simple

import (
	"fmt"
	"testing"
)

type MajorityElement2 struct{}

func (MajorityElement2) Description(string) string {
	return `
		给定一个大小为 n 的整数数组，找出其中所有出现超过 ⌊ n/3 ⌋ 次的元素。
	
		进阶：尝试设计时间复杂度为 O(n)、空间复杂度为 O(1)的算法解决此问题。

		示例 1：
		
		输入：[3,2,3]
		输出：[3]

		示例 2：
		
		输入：nums = [1]
		输出：[1]

		示例 3：
		
		输入：[1,1,1,3,3,2,2,2]
		输出：[1,2]

		`
}

type NumCount struct {
	num   int
	count int
}

//思路说明：如果在一个数组中，有数字超过数组长度的1/3,那么对于整个数组来说，最多有两个数字能够超过1/3
//多数投票升级版：
//
//超过n/3的数最多只能有两个。先选出两个候选人A,B。 遍历数组，分三种情况：
//	1.如果投A（当前元素等于A），则A的票数++;
//	2.如果投B（当前元素等于B），B的票数++；
//	3.如果A,B都不投（即当前与A，B都不相等）,那么检查此时A或B的票数是否减为0：
//		3.1 如果为0,则当前元素成为新的候选人；
//		3.2 如果A,B两个人的票数都不为0，那么A,B两个候选人的票数均减一；
//遍历结束后选出了两个候选人，但是这两个候选人是否满足>n/3，还需要再遍历一遍数组，找出两个候选人的具体票数。

func majorityElement2(nums []int) []int {
	result := make([]int, 0)
	var numCount1, numCount2 NumCount
	for _, num := range nums {
		switch num {
		case numCount1.num:
			numCount1.count++
			continue
		case numCount2.num:
			numCount2.count++
		default:
			if numCount1.count == 0 {
				numCount1.num = num
				numCount1.count++
				continue
			}

			if numCount2.count == 0 {
				numCount2.num = num
				numCount2.count++
				continue
			}

			numCount1.count--
			numCount2.count--
		}
	}

	//遍历结束后选出了两个候选人，但是这两个候选人是否满足>n/3，还需要再遍历一遍数组，找出两个候选人的具体票数。
	var count1, count2 int
	for _, num := range nums {
		if numCount1.num == num {
			count1++
		}
		if numCount2.num == num && numCount1.num != numCount2.num {
			count2++
		}
	}

	if count1 > len(nums)/3 {
		result = append(result, numCount1.num)
	}
	if count2 > len(nums)/3 {
		result = append(result, numCount2.num)
	}

	return result
}

func TestMajorityElement2(t *testing.T) {
	fmt.Println(majorityElement2([]int{3, 2, 3}))
	fmt.Println(majorityElement2([]int{1, 1, 1, 3, 3, 2, 2, 2}))
	fmt.Println(majorityElement2([]int{1}))
}
