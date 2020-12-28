/*
 *  @Author : huangzj
 *  @Time : 2020/12/28 17:32
 *  @Description：
 */

package sJson

import (
	"fmt"
	"github.com/tidwall/sjson"
	"testing"
)

func TestAdvanceSJson(t *testing.T) {
	//通过通配符的方式进行匹配找到对应字段，这个是没办法做到的所以这边会报错
	user := `{"name":"dj","age":18}`
	newValue, err := sjson.Set(user, "na?e", "dajun")
	fmt.Println(err, newValue)

}

func TestErrorJson(t *testing.T) {
	//SJson不会检查json串是否正确，只会把设置的值进行返回，所以在使用的时候需要进行json串的正确性校验
	user := `{"name":dj,age:18}`
	newValue, err := sjson.Set(user, "name", "dajun")
	fmt.Println(err, newValue)
}
