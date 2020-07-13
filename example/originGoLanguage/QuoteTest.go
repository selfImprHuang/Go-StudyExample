/*
 *  @Author : huangzj
 *  @Time : 2020/7/13 13:43
 *  @Description：传值和传址
 */

package originGoLanguage

import "fmt"

func TestQuote() {
	x, y := "i am x", "i am y"

	fmt.Println("回参的值2")
	x2, y2 := byValue(x, y)
	fmt.Println(x2)
	fmt.Println(y2)

	fmt.Println("传值的方式(原始值)")
	fmt.Println(x)
	fmt.Println(y)

	fmt.Println("回参的值")
	x1, y1 := byAddress(&x, &y)
	fmt.Println(x1)
	fmt.Println(y1)

	fmt.Println("传引用的方式(原始值)")
	fmt.Println(x)
	fmt.Println(y)

}

func byAddress(i *string, i2 *string) (string, string) {
	*i2, *i = *i, *i2
	return *i, *i2
}

func byValue(s string, s2 string) (string, string) {
	s, s2 = s2, s
	return s, s2
}
