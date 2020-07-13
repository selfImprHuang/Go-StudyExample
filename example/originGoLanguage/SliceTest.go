/*
 *  @Author : huangzj
 *  @Time : 2020/7/13 14:19
 *  @Description：切片的测试
 */

package originGoLanguage

import "fmt"

func TestSlice() {
	//定义切片
	var number []int
	fmt.Println(number)

	number = append(number, 1)
	number = append(number, 2)
	number = append(number, 3)
	number = append(number, 4)

	fmt.Println(len(number), cap(number), number)

	//切片截取
	fmt.Println(number[0:1])
	fmt.Println(number[:2])
	fmt.Println(number[1:])

	//初始化切片
	number1 := make([]int, len(number), 10)
	fmt.Println(len(number1), cap(number1), number1)

	//复制切片
	copy(number1, number)
	fmt.Println(len(number1), cap(number1), number1)

}
