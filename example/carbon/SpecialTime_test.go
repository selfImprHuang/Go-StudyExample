/*
 *  @Author : huangzj
 *  @Time : 2020/12/25 15:38
 *  @Description：
 */

package carbon

import (
	"fmt"
	"github.com/uniplaces/carbon"
	"log"
	"testing"
	"time"
)

func TestSpecialTime(t *testing.T) {

	//2020 12 25 是周五
	timed, err := carbon.Create(2020, 12, 25, 0, 0, 0, 0, "Asia/Shanghai")
	if err != nil {
		log.Fatal(err)
	}
	//func description 设置(自定义工作日),如果在自定义工作日中的就返回true | 这边还设置了一周的起始和结束
	timed.SetWeekStartsAt(time.Sunday)
	timed.SetWeekEndsAt(time.Saturday)
	timed.SetWeekendDays([]time.Weekday{time.Monday, time.Tuesday, time.Wednesday})

	fmt.Printf("Today is %s, weekend? %t\n", timed.Weekday(), timed.IsWeekend())

	fmt.Println()

	//func description 增加多少时间减掉多少时间，这边是比较特殊的处理，就是说可以增加周数，月数，年数
	fmt.Printf("Right now is %s\n", carbon.Now().DateTimeString())

	today, _ := carbon.NowInLocation("Japan")
	fmt.Printf("Right now in Japan is %s\n", today)

	fmt.Printf("Tomorrow is %s\n", carbon.Now().AddDay())
	fmt.Printf("Last week is %s\n", carbon.Now().SubWeek())
	fmt.Printf("Last month is %s\n", carbon.Now().SubMonth())

	nextOlympics, _ := carbon.CreateFromDate(2016, time.August, 5, "Europe/London")
	nextOlympics = nextOlympics.AddYears(4)
	fmt.Printf("Next olympics are in %d\n", nextOlympics.Year())

	if carbon.Now().IsWeekend() {
		fmt.Printf("Happy time!")
	}
}
