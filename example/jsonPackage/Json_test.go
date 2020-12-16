/*
 *  @Author : huangzj
 *  @Time : 2020/3/27 15:38
 *  @Description：
 */

package jsonPackage

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

type Time time.Time

const (
	timeFormat = "2006-01-02 15:04:05"
)

//-------------------结构体----------------

type Person struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Birthday Time   `json:"birthday"`
}

type json1 struct {
	Name      string    `json:"name"`       //测试一下别名
	Value     string    `json:"value,-"`    //测试一下,有别名也有 - 能被转换
	Value1    string    `json:"-"`          //测试一下,只有 - 能被转换
	method1   string    `json:"method1"`    //测试一下私有属性不能被转换
	Now       time.Time `json:"now"`        //测试一下时间字段可否被转换
	Now1      Time      `json:"now1"`       //自定义时间转换的格式
	Age       int       `json:",omitempty"` //测试一下 omitempty 值为空不能被转换
	AgeString string    `json:",omitempty"` //测试一下字符串 omitempty 值为空 不能被转换
	Length    int       `json:",String"`    //测试一下tag里面带有string的自动转换
	Person    Person    //测试一下多级的结构
}

//-----------------------------------------示例方法

/*
 * 测试序列化功能
 */
func TestJsonMarshalTest(t *testing.T) {
	j := json1{
		Name:    "name",
		Value:   "value",
		Value1:  "value1",
		method1: "method1",
		Now:     time.Now(),
		Now1:    Time(time.Now()),
		Age:     0,
		Length:  24,
		Person: Person{
			Id:       123,
			Name:     "xxxx",
			Birthday: Time(time.Now()),
		},
	}

	b, err := json.Marshal(j)
	if err != nil {
		panic(err.Error())
	}

	fmt.Print(string(b))

}

/*
 * 测试序列化功能
 */
func TestJsonUnmarshalTest(t *testing.T) {
	src := `{"id":5,"name":"xiaoming","birthday":"2016-06-30 16:09:51"}`
	p := new(Person)
	err := json.Unmarshal([]byte(src), &p)

	var m map[string]interface{}
	_ = json.Unmarshal([]byte(src), &m)
	if err != nil {
		panic(err.Error())
	}
}

//-------------------------接口实现

//实现该方法，实现对应的时间处理，json应该是没有支持时间处理的
func (t *Time) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+timeFormat+`"`, string(data), time.Local)
	*t = Time(now)
	return
}

/*
 * 实现该方法，确定时间格式的输出
 */
func (t Time) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(timeFormat)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, timeFormat)
	b = append(b, '"')
	return b, nil
}

func (t Time) String() string {
	return time.Time(t).Format(timeFormat)
}
