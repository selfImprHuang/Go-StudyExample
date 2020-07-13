/*
 *  @Author : huangzj
 *  @Time : 2020/7/13 14:22
 *  @Description：
 */

package originGoLanguage

import "fmt"

func TestMap() {
	//没有初始化的是nil,不能进行赋值
	var countryMap1 map[string]string
	fmt.Println(countryMap1 == nil)

	countryMap := make(map[string]string)
	countryMap["a"] = "aCountry"
	countryMap["b"] = "bCountry"
	countryMap["c"] = "cCountry"
	countryMap["d"] = "dCountry"

	for key, value := range countryMap {
		fmt.Println(key + " is " + value)
	}

	//获取Map的值
	mValue, ok := countryMap["a"]
	fmt.Println(ok, mValue)

	mValue1, ok1 := countryMap["aaaa"]
	fmt.Println(ok1, mValue1)

	//删除map中的元素
	delete(countryMap, "a")
	delete(countryMap, "b")
	for key, value := range countryMap {
		fmt.Println(key + " is " + value)
	}
}
