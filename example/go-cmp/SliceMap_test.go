/*
 *  @Author : huangzj
 *  @Time : 2020/12/30 14:08
 *  @Description：map和切片的比较方法
 */

package go_cmp

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"testing"
)

func TestNilEqualEmpty(t *testing.T) {
	var s1 []int
	var s2 = make([]int, 0)

	var m1 map[int]int
	var m2 = make(map[int]int)

	//Go原生的" == "符号是没办法比较 切片和map的.
	//fmt.Println(s1 == s2)
	//fmt.Println(m1 == m2)

	//这边可以通过Equal来比较 切片和map
	fmt.Println("s1 equals s2?", cmp.Equal(s1, s2))
	fmt.Println("m1 equals m2?", cmp.Equal(m1, m2))

	//空的切片和nil如果要比较相等，需要加上对应的参数
	fmt.Println("s1 equals s2 with option?", cmp.Equal(s1, s2, cmpopts.EquateEmpty()))
	fmt.Println("m1 equals m2 with option?", cmp.Equal(m1, m2, cmpopts.EquateEmpty()))

}

//比较无序的切片，和本来就没有顺序的map
func TestSliceOutOfOrder(t *testing.T) {
	s1 := []int{1, 2, 3, 4}
	s2 := []int{4, 3, 2, 1}
	fmt.Println("s1 equals s2?", cmp.Equal(s1, s2))
	fmt.Println("s1 equals s2 with option?", cmp.Equal(s1, s2, cmpopts.SortSlices(func(i, j int) bool { return i < j })))

	fmt.Println()

	m1 := map[int]int{1: 10, 2: 20, 3: 30}
	m2 := map[int]int{1: 10, 2: 20, 3: 30}
	fmt.Println("m1 equals m2?", cmp.Equal(m1, m2))
	fmt.Println("m1 equals m2 with option?", cmp.Equal(m1, m2, cmpopts.SortMaps(func(i, j int) bool { return i < j })))
}
