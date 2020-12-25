/*
 *  @Author : huangzj
 *  @Time : 2020/12/25 16:01
 *  @Description：测试两个时间之间相差多少对应的时间
 */

package carbon

import (
	"fmt"
	"github.com/uniplaces/carbon"
	"testing"
)

func TestTimeDiff(t *testing.T) {
	vancouver, _ := carbon.Today("Asia/Shanghai")
	london, _ := carbon.Today("Asia/Hong_Kong")
	fmt.Println(vancouver.DiffInSeconds(london, true)) // 0

	//func description 测试时间相差几个小时
	ottawa, _ := carbon.CreateFromDate(2000, 1, 1, "America/Toronto")
	vancouver, _ = carbon.CreateFromDate(2000, 1, 1, "America/Vancouver")
	fmt.Println(ottawa.DiffInHours(vancouver, true)) // 3

	fmt.Println(ottawa.DiffInHours(vancouver, false)) // 3
	fmt.Println(vancouver.DiffInHours(ottawa, false)) // -3

	//func description 测试时间相差多少天
	timed, _ := carbon.CreateFromDate(2012, 1, 31, "UTC")
	fmt.Println(timed.DiffInDays(timed.AddMonth(), true))  // 31
	fmt.Println(timed.DiffInDays(timed.SubMonth(), false)) // -31

	timed, _ = carbon.CreateFromDate(2012, 4, 30, "UTC")
	fmt.Println(timed.DiffInDays(timed.AddMonth(), true)) // 30
	fmt.Println(timed.DiffInDays(timed.AddWeek(), true))  // 7

	//func description 测试时间相差多少分钟
	timed, _ = carbon.CreateFromTime(10, 1, 1, 0, "UTC")
	fmt.Println(timed.DiffInMinutes(timed.AddSeconds(59), true))  // 0
	fmt.Println(timed.DiffInMinutes(timed.AddSeconds(60), true))  // 1
	fmt.Println(timed.DiffInMinutes(timed.AddSeconds(119), true)) // 1
	fmt.Println(timed.DiffInMinutes(timed.AddSeconds(120), true)) // 2
}
