/*
 *  @Author : huangzj
 *  @Time : 2020/12/29 9:16
 *  @Description：
 */

package copier

import (
	"fmt"
	"github.com/jinzhu/copier"
	"testing"
)

type User struct {
	Name string
	Age  int
	Role string
}

func (u *User) DoubleAge() int {
	return u.Age * 2
}

type Employee struct {
	Name      string
	Age       int
	SuperRole string
	DoubleAge int
}

func (e *Employee) Role(role string) {
	e.SuperRole = "通过role得到superRole" + role
}

func TestCopier(t *testing.T) {
	user := User{Name: "dj", Age: 18}
	users := []User{
		{Name: "dj", Age: 18, Role: "Admin"},
		{Name: "dj2", Age: 18, Role: "Dev"},
	}
	employee := Employee{}
	var employees []Employee

	//通过 【Role】 方法转换 User的Role属性到 Employee 的SuperRole -- 这个方法名好像要和被转换的一致才行
	//通过 User 的 【DoubleAge】 方法转换到 Employee 的同名属性
	_ = copier.Copy(&employee, &user)
	fmt.Println(fmt.Sprintf("%#v\n", employee))

	//结构体转换到数组，相当于append操作，不过是类型不同，复制具体属性
	_ = copier.Copy(&employees, &user)
	fmt.Println(fmt.Sprintf("%#v\n", employees))

	//不同的结构体数组的转换
	_ = copier.Copy(&employees, &users)
	fmt.Println(fmt.Sprintf("%#v\n", employee))
}
