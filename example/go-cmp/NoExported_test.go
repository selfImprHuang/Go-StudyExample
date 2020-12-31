/*
 *  @Author : huangzj
 *  @Time : 2020/12/30 15:21
 *  @Description：测试未导出字段
 */

package go_cmp

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"reflect"
	"testing"
)

//自定义是否比较未导出字段，这个方法和上面的很类似，都是必须制定相应的结构体
func TestDiyExportMethod(t *testing.T) {
	u1 := User2{"dj", 18, Address{}}
	u2 := User2{"dj", 18, Address{}}

	fmt.Println("u1 equals u2?", cmp.Equal(u1, u2, cmp.Exporter(allowUnExportedInType)))
}

//测试只忽略当前对象的未导出属性
func TestExportOne(t *testing.T) {
	c1 := Contactx{Phone: "123456789", Email: "dj@example.com"}
	c2 := Contactx{Phone: "123456789", Email: "dj@example.com"}

	u1 := Userx{"dj", 18, c1}
	u2 := Userx{"dj", 18, c2}

	//i think 通过 IgnoreUnexported 这个属性可以实现忽略具体类型的未导出字段，但是好像没办法忽略 【指针类型】 的未导出字段
	fmt.Println("u1 equals u2?", cmp.Equal(u1, u2, cmpopts.IgnoreUnexported(Userx{})))
}

//结构体中有未导出字段的时候,通过Equal进行比较的时候，会报错，除非设置忽略
func TestExportError(t *testing.T) {
	c1 := &Contact1{Phone: "123456789", Email: "dj@example.com"}
	c2 := &Contact1{Phone: "123456789", Email: "dj@example.com"}

	u1 := User1{"dj", 18, c1}
	u2 := User1{"dj", 18, c2}

	fmt.Println("u1 equals u2?", cmp.Equal(u1, u2))
}

//使用 【IgnoreUnexported】,但是在结构体下层还有其他非导出的字段也会报错，这个字段没办法深入到里面
func TestExportError2(t *testing.T) {
	u11 := User2{"dj", 18, Address{}}
	u22 := User2{"dj", 18, Address{}}

	fmt.Println("u1 equals u2?", cmp.Equal(u11, u22, cmpopts.IgnoreUnexported(User2{})))
}

func allowUnExportedInType(t reflect.Type) bool {
	if t.Name() == "Address" {
		return true
	}

	return false
}

type Contactx struct {
	Phone string
	Email string
}

type Userx struct {
	Name    string
	Age     int
	contact Contactx //注意这边是结构体
}

type Contact1 struct {
	Phone string
	Email string
}

type User1 struct {
	Name    string
	Age     int
	contact *Contact1
}

type Address struct {
	Province string
	city     string
}

type User2 struct {
	Name    string
	Age     int
	Address Address
}
