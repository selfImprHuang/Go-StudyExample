/*
 *  @Author : huangzj
 *  @Time : 2020/7/13 14:27
 *  @Description：通道测试
 */

package originGoLanguage

import (
	"fmt"
	"testing"
)

func TestChannel(t *testing.T) {

	list := []int{0, 1, 45, -12, 33, 90, -22, 100}
	//通道的初始化
	c := make(chan int)

	go sum(list[len(list)/2:], c)
	go sum(list[:len(list)/2], c)

	//通道获取值的赋值
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)
	//通道关闭
	close(c)
}

func sum(list []int, c chan int) {
	sum := 0
	for _, i := range list {
		sum += i
	}
	c <- sum
}
