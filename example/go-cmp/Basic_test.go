/*
 *  @Author : huangzj
 *  @Time : 2020/12/30 14:43
 *  @Description：
 */

package go_cmp

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"testing"
)

//测试指向同一个结构体的对象
func TestStruct(t *testing.T) {
	u1 := User{Name: "dj", Age: 18}
	u2 := User{Name: "dj", Age: 18}

	fmt.Println("u1 == u2?", u1 == u2)
	fmt.Println("u1 equals u2?", cmp.Equal(u1, u2))

	c1 := &Contact{Phone: "123456789", Email: "dj@example.com"}

	u1.Contact = c1
	u2.Contact = c1
	fmt.Println("u1 == u2 with same pointer?", u1 == u2)
	fmt.Println("u1 equals u2 with same pointer?", cmp.Equal(u1, u2))

}

//测试指向不同结构体，但是结构体值一致的对象
func TestPoint(t *testing.T) {
	u1 := User{Name: "dj", Age: 18}
	u2 := User{Name: "dj", Age: 18}

	fmt.Println("u1 == u2?", u1 == u2)
	fmt.Println("u1 equals u2?", cmp.Equal(u1, u2))

	c1 := &Contact{Phone: "123456789", Email: "dj@example.com"}
	c2 := &Contact{Phone: "123456789", Email: "dj@example.com"}

	u1.Contact = c1
	u2.Contact = c2
	fmt.Println("u1 == u2 with different pointer?", u1 == u2)
	fmt.Println("u1 equals u2 with different pointer?", cmp.Equal(u1, u2))
}
