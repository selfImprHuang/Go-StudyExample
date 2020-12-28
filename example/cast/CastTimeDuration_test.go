/*
 *  @Author : huangzj
 *  @Time : 2020/12/28 9:52
 *  @Description：
 */

package cast

import (
	"fmt"
	"github.com/spf13/cast"
	"testing"
	"time"
)

func TestTimeDuration(t *testing.T) {
	now := time.Now()
	timestamp := 1579615973
	timeStr := "2020-01-21 22:13:48"

	fmt.Println(cast.ToTime(now))       // 2020-01-22 06:31:50.5068465 +0800 CST m=+0.000997701(时间是当前时间，格式如此)
	fmt.Println(cast.ToTime(timestamp)) // 2020-01-21 22:12:53 +0800 CST
	fmt.Println(cast.ToTime(timeStr))   // 2020-01-21 22:13:48 +0000 UTC

	d, _ := time.ParseDuration("1m30s")
	ns := 30000
	strWithUnit := "130s"
	strWithoutUnit := "130"

	fmt.Println(cast.ToDuration(d))              // 1m30s
	fmt.Println(cast.ToDuration(ns))             // 30µs
	fmt.Println(cast.ToDuration(strWithUnit))    // 2m10s
	fmt.Println(cast.ToDuration(strWithoutUnit)) // 130ns
}
