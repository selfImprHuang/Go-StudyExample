/*
 *  @Author : huangzj
 *  @Time : 2020/3/26 18:18
 *  @Description：
 */

package entity

import "time"

type UserType struct {
	UserTypeId   int
	UserTypeName string
}

type User struct {
	UserType  UserType  `jpath:"userType"`
	LoginName string    `jpath:"loginName"`
	T         time.Time `jpath:"t"`
}
