/*
 *  @Author : huangzj
 *  @Time : 2020/12/28 9:18
 *  @Description：测试cast，这个包是把interface，根据类型进行转换的处理
 */

package cast

import (
	"fmt"
	"github.com/spf13/cast"
	"testing"
)

func TestCast(t *testing.T) {
	// ToString
	fmt.Println(cast.ToString("leedarjun"))        // leedarjun
	fmt.Println(cast.ToString(8))                  // 8
	fmt.Println(cast.ToString(8.31))               // 8.31
	fmt.Println(cast.ToString([]byte("one time"))) // one time
	fmt.Println(cast.ToString(nil))                // ""

	var foo interface{} = "one more time"
	fmt.Println(cast.ToString(foo)) // one more time

	// ToInt
	fmt.Println(cast.ToInt(8))     // 8
	fmt.Println(cast.ToInt(8.31))  // 8
	fmt.Println(cast.ToInt("8"))   // 8
	fmt.Println(cast.ToInt(true))  // 1
	fmt.Println(cast.ToInt(false)) // 0

	var eight interface{} = 8
	fmt.Println(cast.ToInt(eight)) // 8
	fmt.Println(cast.ToInt(nil))   // 0
	//To..E方法会返回转换错误的报错信息
	_, e := cast.ToIntE("asda")
	if e != nil {
		fmt.Println(e)
	}
}
