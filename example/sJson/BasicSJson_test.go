/*
 *  @Author : huangzj
 *  @Time : 2020/12/28 17:29
 *  @Description：
 */

package sJson

import (
	"fmt"
	"github.com/tidwall/sjson"
	"testing"
)

func TestSJsonSet(t *testing.T) {
	//测试设值
	const json = `{"name":{"first":"li","last":"dj"},"age":18}`

	value, _ := sjson.Set(json, "name.last", "dajun")
	fmt.Println(value)

}

func TestSJsonDelete(t *testing.T) {
	//进行字段删除
	var newValue string
	user := `{"name":{"first":"li","last":"dj"},"age":18}`

	newValue, _ = sjson.Delete(user, "name.first")
	fmt.Println(newValue)

	newValue, _ = sjson.Delete(user, "name.full")
	fmt.Println(newValue)

	fruits := `{"fruits":["apple", "orange", "banana"]}`

	newValue, _ = sjson.Delete(fruits, "fruits.1")
	fmt.Println(newValue)

	newValue, _ = sjson.Delete(fruits, "fruits.-1")
	fmt.Println(newValue)

	newValue, _ = sjson.Delete(fruits, "fruits.5")
	fmt.Println(newValue)
}

//测试操作数组,【.下标】表示替代的是哪个位置的元素，-1比较特殊是在数组后面进行添加，如果超过数组长度会自动填充 null上去，这个好像有点尴尬
func TestSJsonOperateList(t *testing.T) {
	fruits := `{"fruits":["apple", "orange", "banana"]}`

	var newValue string
	newValue, _ = sjson.Set(fruits, "fruits.1", "grape")
	fmt.Println(newValue)

	newValue, _ = sjson.Set(fruits, "fruits.3", "pear")
	fmt.Println(newValue)

	newValue, _ = sjson.Set(fruits, "fruits.-1", "strawberry")
	fmt.Println(newValue)

	newValue, _ = sjson.Set(fruits, "fruits.5", "watermelon")
	fmt.Println(newValue)

}

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

//测试SJson支持设置的所有类型
func TestSJsonAllType(t *testing.T) {
	nilJSON, _ := sjson.Set("", "key", nil)
	fmt.Println(nilJSON)

	boolJSON, _ := sjson.Set("", "key", false)
	fmt.Println(boolJSON)

	intJSON, _ := sjson.Set("", "key", 1)
	fmt.Println(intJSON)

	floatJSON, _ := sjson.Set("", "key", 10.5)
	fmt.Println(floatJSON)

	strJSON, _ := sjson.Set("", "key", "hello")
	fmt.Println(strJSON)

	mapJSON, _ := sjson.Set("", "key", map[string]interface{}{"hello": "world"})
	fmt.Println(mapJSON)

	u := User{Name: "dj", Age: 18}
	structJSON, _ := sjson.Set("", "key", u)
	fmt.Println(structJSON)
}
