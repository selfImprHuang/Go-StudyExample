/*
 *  @Author : huangzj
 *  @Time : 2020/12/30 14:57
 *  @Description：
 */

package go_cmp

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"testing"
)

type NetAddr1 struct {
	IP   string
	Port int
}

func compareNetAddr(a, b NetAddr1) bool {
	if a.Port != b.Port {
		return false
	}

	if a.IP != b.IP {
		if a.IP == "127.0.0.1" && b.IP == "localhost" {
			return true
		}

		if a.IP == "localhost" && b.IP == "127.0.0.1" {
			return true
		}

		return false
	}

	return true
}

//自定义比较器：
// 		这种方式与上面介绍的自定义Equal()方法有些类似，但更灵活。
// 		有时，我们要自定义比较操作的类型定义在第三方包中，这样就无法给它定义Equal方法。
// 		这时，我们就可以采用自定义Comparer的方式。
func TestDiyCompareMethod(t *testing.T) {
	a1 := NetAddr1{"127.0.0.1", 5000}
	a2 := NetAddr1{"localhost", 5000}

	fmt.Println("a1 equals a2?", cmp.Equal(a1, a2))
	fmt.Println("a1 equals a2 with comparer?", cmp.Equal(a1, a2, cmp.Comparer(compareNetAddr)))
}
