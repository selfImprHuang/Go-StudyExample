package example

import (
	entity2 "Go-StudyExample/entity"
	"encoding/json"
	"fmt"
	"github.com/goinggo/mapstructure"
)

//-----------------------json数据---------------------
var document = `{"loginName":"sptest1","userType":{"userTypeId":1,"userTypeName":"normal_user","t":"2026-01-02 15:04:05"}}`

var document1 = `{"cobrandId":10010352,"channelId":-1,"locale":"en_US","tncVersion":2,"people":[{"name":"jack","age":{"birth":10,"year":2000,"animals":[{"barks":"yes","tail":"yes"},{"barks":"no","tail":"yes"}]}},{"name":"jill","age":{"birth":11,"year":2001}}]}`

var document2 = `[{"name":"bill"},{"name":"lisa"}]`

//-----------------------------结构体-----------------------
type Animal struct {
	Barks string `jpath:"barks"`
}

type People struct {
	Age     int      `jpath:"age.birth"` // jpath is relative to the array
	Animals []Animal `jpath:"age.animals"`
}

type Items struct {
	Categories []string `jpath:"categories"`
	Peoples    []People `jpath:"people"` // Specify the location of the array
}

//---------------------------------测试方法------------------------

func MapStructureTestFunc() {
	var te entity2.Entity
	m := make(map[string]interface{})
	m["Num"] = 1
	m["S"] = "test"
	m["T"] = map[string]string{"1": "1", "2": "2"}

	err := mapstructure.Decode(m, &te)
	if err != nil {
		panic(err.Error())
	}

	fmt.Print(te.Num, " ", te.S, " ", te.T)
}

func MapStructureTestFunc1() {
	var docMap map[string]interface{}
	_ = json.Unmarshal([]byte(document), &docMap)

	var user entity2.User
	err := mapstructure.Decode(docMap, &user)

	if err != nil {
		panic(err.Error())
	}
	fmt.Println(user.T.Format("2006-01-02 15:04:05"))
	fmt.Println(user, " ", user.UserType.UserTypeId, " ", user.UserType.UserTypeName)
}

type NameDoc struct {
	Name string `jpath:"name"`
}

func MapStructureTestFunc2() {

	sliceScript := []byte(document2)
	var sliceMap []map[string]interface{}
	_ = json.Unmarshal(sliceScript, &sliceMap)

	var myslice []NameDoc
	err := mapstructure.DecodeSlicePath(sliceMap, &myslice)

	if err != nil {
		panic(err.Error())
	}

	fmt.Println(myslice[0], " ", myslice[1])
}

func MapStructureTestFunc3() {
	docScript := []byte(document1)
	var docMap map[string]interface{}
	_ = json.Unmarshal(docScript, &docMap)

	var items Items
	err := mapstructure.DecodePath(docMap, &items)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(items.Peoples[0], items.Peoples[1])
}
