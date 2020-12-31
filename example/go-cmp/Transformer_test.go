/*
 *  @Author : huangzj
 *  @Time : 2020/12/31 14:33
 *  @Description： 比较的时候转换对应的属性，改变结构体对比的规则
 */

package go_cmp

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"testing"
)

type UserO struct {
	Name string
	Age  int
}

func omitAge(u UserO) string {
	return u.Name
}

type UserO2 struct {
	Name    string
	Age     int
	Email   string
	Address string
}

func omitAge2(u UserO2) UserO2 {
	return UserO2{u.Name, 0, u.Email, u.Address}
}

//对某个字段忽略比较，也就是返回字段默认值
func TestOmit(t *testing.T) {
	u1 := UserO{Name: "dj", Age: 18}
	u2 := UserO{Name: "dj", Age: 28}

	//i think 这边omitAge只返回name属性应该是值比较UserO的name属性即可
	fmt.Println("u1 equals u2?", cmp.Equal(u1, u2, cmp.Transformer("omitAge", omitAge)))

	u3 := UserO2{Name: "dj", Age: 18, Email: "dj@example.com"}
	u4 := UserO2{Name: "dj", Age: 28, Email: "dj@example.com"}

	u5 := UserO2{Name: "dj", Age: 18, Email: "dj@example.com"}
	u6 := UserO2{Name: "dj1", Age: 28, Email: "dj@example.com"}

	//i think 这边的omitAge2返回修改之后的UserO2,只把age设置为0，其他属性不变，说明还是会参与到比较
	fmt.Println("u3 equals u4?", cmp.Equal(u3, u4, cmp.Transformer("omitAge", omitAge2)))

	fmt.Println("u5 equals u5?", cmp.Equal(u5, u6, cmp.Transformer("omitAge", omitAge2)))
}

type NetAddrO struct {
	IP   string
	Port int
}

func transformLocalhost(a NetAddrO) NetAddrO {
	if a.IP == "localhost" {
		return NetAddrO{IP: "127.0.0.1", Port: a.Port}
	}

	return a
}

//对结构体进行对比的时候，修改某字段的值
func TestTransformValue(t *testing.T) {
	a1 := NetAddrO{"127.0.0.1", 5000}
	a2 := NetAddrO{"localhost", 5000}

	//这边是把NetAddr对应的IP进行转换，localhost == 127.0.0.1 所以直接进行转换，比较的结果就会一致
	fmt.Println("a1 equals a2?", cmp.Equal(a1, a2, cmp.Transformer("localhost", transformLocalhost)))
}
