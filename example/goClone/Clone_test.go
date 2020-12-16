/*
 *  @Author : huangzj
 *  @Time : 2020/5/13 17:52
 *  @Description：测试深拷贝和浅拷贝的代码类
 */

package example

import (
	"Go-StudyExample/entity"
	"Go-StudyExample/util"
	"fmt"
	"testing"
)

func TestClone(t *testing.T) {
	fmt.Println("由此可见其实Go语言对于内存的分配，相对对象是独一份，对于对象的基本类型和值类型（结构体），应该是在其内部开辟了内存空间进行存储，而对于" +
		"指针类型来说，直接使用它的地址（指针），不对结构体进行复制")
}

//json包处理深度克隆
func TestDeepCloneByJsonPackage(t *testing.T) {
	rule := entity.SynthesisRule{
		SynthesisId:    "112",
		SynthesisNum:   0,
		CostItems:      nil,
		SynthesisItems: nil,
	}
	//这边必须是指定的类型，不然没办法转回去
	rule1, err := util.DeepCopyByJson(rule)
	if err == nil && rule1 != &rule {
		fmt.Println("")
		fmt.Println("通过json包的方式可以实现深度克隆")
	}
}

//序列化方式进行深度克隆
func TestDeepCloneByEncode(t *testing.T) {
	rule := entity.SynthesisRule{
		SynthesisId:    "112",
		SynthesisNum:   0,
		CostItems:      nil,
		SynthesisItems: nil,
	}
	rule1 := new(entity.SynthesisRule)
	err := util.DeepCopyByGob(rule1, rule)
	if err != nil {
		panic(err)
	}
	rule1.SynthesisId = "123"
	if rule1 != &rule {
		fmt.Println("通过序列化的方式可以实现深度克隆")
	}
}

//结构体带基本类型
func TestJustBasicType(t *testing.T) {
	rule := entity.SynthesisRule{
		SynthesisId:    "112",
		SynthesisNum:   0,
		CostItems:      nil,
		SynthesisItems: nil,
	}
	rule1 := rule
	rule.SynthesisId = "222"
	rule.SynthesisItems = make([][]int, 0)
	rule.SynthesisItems = append(rule.SynthesisItems, []int{1})

	if &rule1 != &rule {
		fmt.Println("结构体中不存在指针，直接用 = 是")
		fmt.Println("深克隆")
		fmt.Println("")
	}

}

//结构体带指针
func TestHasPointType(t *testing.T) {
	cmp := entity.SynthesisRuleCmp{
		SynthesisRule: &entity.SynthesisRule{
			SynthesisId:    "112",
			SynthesisNum:   0,
			CostItems:      nil,
			SynthesisItems: nil,
		},
		RemainingSynthesisTime: 0,
		Sort:                   0,
	}

	cmp1 := cmp
	cmp.SynthesisId = "123"
	if &cmp != &cmp1 && cmp.SynthesisRule == cmp1.SynthesisRule {
		fmt.Println("结构体中存在指针，直接用 = 是深克隆")
		fmt.Println("但是对于结构体中的指针来说是浅克隆,直接使用地址")
		fmt.Println("")
	}
}

//结构体带切片
func TestHasPointTypeSlice(t *testing.T) {

	cmp := entity.SynthesisRuleCmp{
		SynthesisRule: &entity.SynthesisRule{
			SynthesisId:    "112",
			SynthesisNum:   0,
			CostItems:      nil,
			SynthesisItems: nil,
		},
		RemainingSynthesisTime: 0,
		Sort:                   0,
	}
	cmp1 := entity.SynthesisRuleCmp{
		SynthesisRule: &entity.SynthesisRule{
			SynthesisId:    "112111",
			SynthesisNum:   0,
			CostItems:      nil,
			SynthesisItems: nil,
		},
		RemainingSynthesisTime: 0,
		Sort:                   0,
	}
	list := make([]entity.SynthesisRuleCmp, 0)
	list = append(list, cmp1)
	list = append(list, cmp)

	list1 := make([]entity.SynthesisRuleCmp, len(list))
	cmp.RemainingSynthesisTime = 1000
	fmt.Println(cmp1.RemainingSynthesisTime)
	copy(list1, list)
	list1[0].SynthesisId = "12333"
	if list[0].SynthesisId == list1[0].SynthesisId {
		fmt.Println("非深度克隆")
	}

	list2 := list[0:]
	list2[0].SynthesisId = "54323"
	if list[0].SynthesisId == list2[0].SynthesisId {
		fmt.Println("非深度克隆")
	}
}
