/*
 *  @Author : huangzj
 *  @Time : 2020/12/16 10:33
 *  @Description：
 */

package gJson

import (
	"fmt"
	"strings"
	"testing"
)
import "github.com/tidwall/gjson"

//获取json中的某一个键的值
func TestCommonOperation(t *testing.T) {
	json := `{"name":{"first":"www.topgoer.com","last":"dj"},"age":18}`
	lastName := gjson.Get(json, "name.last")
	fmt.Println("last name:", lastName.String())

	age := gjson.Get(json, "age")
	fmt.Println("age:", age.Int())
}

//检验json字符串是否合法
func TestValid(t *testing.T) {
	var json = `{"name":dj,age:18}`
	if !gjson.Valid(json) {
		fmt.Println("error")
	} else {
		fmt.Println("ok")
	}
}

//测试一次性获取json的多个键值
func TestGetMany(t *testing.T) {
	var json = `
		{
		  "name":"dj",
		  "age":18,
		  "pets": ["cat", "dog"],
		  "contact": {
			"phone": "123456789",
			"email": "dj@example.com"
		  }
	}`

	// .#表示数组长度
	results := gjson.GetMany(json, "name", "age", "pets.#", "contact.phone")
	for _, result := range results {
		fmt.Println(result)
	}
}

//测试遍历 Json的数组元素或者结构体所有字段
func TestTraverse(t *testing.T) {
	var json = `
	{
	  "name":"dj",
	  "age":18,
	  "pets": ["cat", "dog"],
	  "contact": {
		"phone": "123456789",
		"email": "dj@example.com"
	  }
	}`

	pets := gjson.Get(json, "pets")
	pets.ForEach(func(_, pet gjson.Result) bool {
		fmt.Println(pet)
		return true
	})

	contact := gjson.Get(json, "contact")
	contact.ForEach(func(key, value gjson.Result) bool {
		fmt.Println(key, value)
		return true
	})
}

// 测试gjson中修饰符的作用
// @reverse：翻转一个数组；
// @ugly：移除 JSON 中的所有空白符；
// @pretty：使 JSON 更易用阅读；
// @this：返回当前的元素，可以用来返回根元素；
// @valid：校验 JSON 的合法性；
// @flatten：数组平坦化，即将["a", ["b", "c"]]转为["a","b","c"]；
// @join：将多个对象合并到一个对象中。
func TestModifier(t *testing.T) {
	var json = `{
	  "name":{"first":"Tom", "last": "Anderson"},
	  "age": 37,
	  "children": ["Sara", "Alex", "Jack"],
	  "fav.movie": "Dear Hunter",
	  "friends": [
		{"first": "Dale", "last":"Murphy", "age": 44, "nets": ["ig", "fb", "tw"]},
		{"first": "Roger", "last": "Craig", "age": 68, "nets": ["fb", "tw"]},
		{"first": "Jane", "last": "Murphy", "age": 47, "nets": ["ig", "tw"]}
	  ]
	}`

	fmt.Println(gjson.Get(json, "children|@reverse"))
	fmt.Println(gjson.Get(json, "children|@reverse|0"))
	fmt.Println(gjson.Get(json, "friends|@ugly"))
	fmt.Println(gjson.Get(json, "friends|@pretty"))
	fmt.Println(gjson.Get(json, "@this"))

	nestedJSON := `{"nested": ["one", "two", ["three", "four"]]}`
	fmt.Println(gjson.Get(nestedJSON, "nested|@flatten"))

	userJSON := `{"info":[{"name":"dj", "age":18},{"phone":"123456789","email":"dj@example.com"}]}`
	fmt.Println(gjson.Get(userJSON, "info|@join"))
}

//测试自定义修饰符
func TestDiyModifier(t *testing.T) {
	gjson.AddModifier("case", diyModifierCaseFunc)

	const json = `{"children": ["Sara", "Alex", "Jack"]}`
	fmt.Println(gjson.Get(json, "children|@case:upper"))
	fmt.Println(gjson.Get(json, "children|@case:lower"))
}

//json: 对应的字符串内容
//arg:  标签后面的内容
func diyModifierCaseFunc(json, arg string) string {
	if arg == "upper" {
		return strings.ToUpper(json)
	}

	if arg == "lower" {
		return strings.ToLower(json)
	}

	return json
}

// 测试键路径，有点类似通配符的标识
// children.#：返回数组children的长度；
// children.1：读取数组children的第 2 个元素（注意索引从 0 开始）；
// child*.2：首先child*匹配children，.2读取第 3 个元素；
// c?ildren.0：c?ildren匹配到children，.0读取第一个元素；
// fav.\moive：因为键名中含有.，故需要\转义；
// friends.#.first：如果数组后#后还有内容，则以后面的路径读取数组中的每个元素，返回一个新的数组。所以该查询返回的数组所有friends的first字段组成；
// friends.1.last：读取friends第 2 个元素的last字段。
func TestKeyPath(t *testing.T) {
	var json = `
	{
	  "name":{"first":"Tom", "last": "Anderson"},
	  "age": 37,
	  "children": ["Sara", "Alex", "Jack"],
	  "fav.movie": "Dear Hunter",
	  "friends": [
		{"first": "Dale", "last":"Murphy", "age": 44, "nets": ["ig", "fb", "tw"]},
		{"first": "Roger", "last": "Craig", "age": 68, "nets": ["fb", "tw"]},
		{"first": "Jane", "last": "Murphy", "age": 47, "nets": ["ig", "tw"]}
	  ]
	}
	`

	fmt.Println("last name:", gjson.Get(json, "name.last"))
	fmt.Println("age:", gjson.Get(json, "age"))
	fmt.Println("children:", gjson.Get(json, "children"))
	fmt.Println("children count:", gjson.Get(json, "children.#"))
	fmt.Println("second child:", gjson.Get(json, "children.1"))
	fmt.Println("third child*:", gjson.Get(json, "child*.2"))
	fmt.Println("first c?ild:", gjson.Get(json, "c?ildren.0"))
	fmt.Println("fav.moive", gjson.Get(json, `fav.\moive`))
	fmt.Println("first name of friends:", gjson.Get(json, "friends.#.first"))
	fmt.Println("last name of second friend:", gjson.Get(json, "friends.1.last"))
}

//遍历json的每一行.这个应该需要的是json是按照正常的格式进行输入的
func TestEachLine(t *testing.T) {
	var json = `
	{"name": "Gilbert", "age": 61}
	{"name": "Alexa", "age": 34}
	{"name": "May", "age": 57}
	{"name": "Deloise", "age": 44}`

	fmt.Println(gjson.Get(json, "..#"))
	fmt.Println(gjson.Get(json, "..1"))
	fmt.Println(gjson.Get(json, "..#.name"))
	fmt.Println(gjson.Get(json, `..#(name="May").age`))
}
