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
)

func TestCastPoint(t *testing.T) {
	p := new(int)
	*p = 8
	fmt.Println(cast.ToInt(p)) // 8

	pp := &p
	fmt.Println(cast.ToInt(pp)) // 8
}
