/*
 *  @Author : huangzj
 *  @Time : 2020/12/24 14:56
 *  @Description：
 */

package SingleNum

import "testing"

type SimpleNumber3 struct{}

func (SimpleNumber3) Description(string) string {
	return `
			给定一个整数数组 nums，其中恰好有两个元素只出现一次，其余所有元素均出现两次。 找出只出现一次的那两个元素。
			
			示例 :
			
			输入: [1,2,1,3,2,5]
			输出: [3,5]
			注意：
			
			结果输出的顺序并不重要，对于上面的例子， [5, 3] 也是正确答案。
			你的算法应该具有线性时间复杂度。你能否仅使用常数空间复杂度来实现？
		`

}

func singleNumber3(nums []int) int {
	//todo 待解答：https://leetcode-cn.com/problems/single-number-iii/solution/zhi-chu-xian-yi-ci-de-shu-zi-iii-by-leetcode/
	//todo https://leetcode-cn.com/problems/single-number-iii/solution/java-yi-huo-100-yao-shi-kan-bu-dong-wo-jiu-qu-chu-/
	return 1
}

func TestSimpleNum3(t *testing.T) {

}
