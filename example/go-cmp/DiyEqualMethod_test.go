/*
 *  @Author : huangzj
 *  @Time : 2020/12/30 14:49
 *  @Description：自定义Equal方法...
 */

package go_cmp

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"testing"
)

type NetAddr struct {
	IP   string
	Port int
}

func (a NetAddr) Equal(b NetAddr) bool {
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

func TestDiyEqualMethod(t *testing.T) {
	a1 := NetAddr{"127.0.0.1", 5000}
	a2 := NetAddr{"localhost", 5000}
	a3 := NetAddr{"192.168.1.1", 5000}

	fmt.Println("a1 equals a2?", cmp.Equal(a1, a2))
	fmt.Println("a1 equals a3?", cmp.Equal(a1, a3))
}
