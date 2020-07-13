/*
 *  @Author : huangzj
 *  @Time : 2020/7/13 14:24
 *  @Description：go运行时错误测试
 */

package originGoLanguage

import (
	"errors"
	"fmt"
)

func TestError() {
	fmt.Println("测试创建一个新的错误")
	testNewError()

	fmt.Println("")
}

func testNewError() {
	var fl1 float32 = 0

	fl2, err := tError(fl1)
	fmt.Println(fl2, err)

	var flx float32 = 2.2
	fl3, err1 := tError(flx)
	fmt.Println(fl3, err1)
}

func tError(f float32) (float32, error) {
	if f == 0 {
		return 0, errors.New("错误")
	}
	return f, nil
}
