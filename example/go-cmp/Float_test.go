/*
 *  @Author : huangzj
 *  @Time : 2020/12/30 15:11
 *  @Description：
 */

package go_cmp

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"math"
	"testing"
)

type FloatPair struct {
	X float64
	Y float64
}

func TestBasicFloat(t *testing.T) {
	//特殊的浮点数NaN（Not a Number），它与任何浮点数都不等，包括它自己
	p1 := FloatPair{X: math.NaN()}
	p2 := FloatPair{X: math.NaN()}
	fmt.Println("p1 equals p2?", cmp.Equal(p1, p2))

	//受限于变量的存储空间，就会存在误差
	f1 := 0.1
	f2 := 0.2
	f3 := 0.3
	p3 := FloatPair{X: f1 + f2}
	p4 := FloatPair{X: f3}
	fmt.Println("p3 equals p4?", cmp.Equal(p3, p4))

	//Go 语言中这些字面量的运算直接是在编译器完成的，所以同样的值虽然受限于变量的存储空间，但是结果是一样的.
	p5 := FloatPair{X: 0.1 + 0.2}
	p6 := FloatPair{X: 0.3}
	fmt.Println("p5 equals p6?", cmp.Equal(p5, p6))
}

//通过equal方法测试NaN返回结果一致
func TestNanEqual(t *testing.T) {
	p1 := FloatPair{X: math.NaN()}
	p2 := FloatPair{X: math.NaN()}
	fmt.Println("p1 equals p2?", cmp.Equal(p1, p2, cmpopts.EquateNaNs()))
}

//测试误差在一定的范围内  具体表示是：|x-y| ≤ max(fraction*min(|x|, |y|), margin)
func TestCountError(t *testing.T) {
	f1 := 0.1
	f2 := 0.2
	f3 := 0.3
	p3 := FloatPair{X: f1 + f2}
	p4 := FloatPair{X: f3}
	fmt.Println(fmt.Sprintf("实际存储的值分别是：\n     %v \n     %v", p3.X, p4.X))
	fmt.Println("p3 equals p4?", cmp.Equal(p3, p4, cmpopts.EquateApprox(0.1, 0.001)))
}
