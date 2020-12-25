/*
 *  @Author : huangzj
 *  @Time : 2020/12/25 17:10
 *  @Description：
 */

package carbon

import (
	"fmt"
	"github.com/uniplaces/carbon"
	"testing"
	"time"
)

func TestTimeModifier(t *testing.T) {

	//测试时间修饰词,获取相应修饰词对应的日期时间
	timed := carbon.Now()
	fmt.Printf("Start of day:%s\n", timed.StartOfDay())
	fmt.Printf("End of day:%s\n", timed.EndOfDay())
	fmt.Printf("Start of month:%s\n", timed.StartOfMonth())
	fmt.Printf("End of month:%s\n", timed.EndOfMonth())
	fmt.Printf("Start of year:%s\n", timed.StartOfYear())
	fmt.Printf("End of year:%s\n", timed.EndOfYear())
	fmt.Printf("Start of decade:%s\n", timed.StartOfDecade())
	fmt.Printf("End of decade:%s\n", timed.EndOfDecade())
	fmt.Printf("Start of century:%s\n", timed.StartOfCentury())
	fmt.Printf("End of century:%s\n", timed.EndOfCentury())
	fmt.Printf("Start of week:%s\n", timed.StartOfWeek())
	fmt.Printf("End of week:%s\n", timed.EndOfWeek())

	fmt.Printf("Next:%s\n", timed.Next(time.Wednesday))
	fmt.Printf("Previous:%s\n", timed.Previous(time.Wednesday))
}
