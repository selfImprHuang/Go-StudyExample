/*
 *  @Author : huangzj
 *  @Time : 2020/4/28 17:17
 *  @Description：
 */

package main

import (
	"fmt"
	"math"
)

func main() {
	x := 0
	for i := 1; i < 31; i++ {
		x = x + 1<<i
	}
	fmt.Println(math.MaxInt32)
	fmt.Println(x)
}
