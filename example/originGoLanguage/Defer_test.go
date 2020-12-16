/*
 *  @Author : huangzj
 *  @Time : 2020/7/13 14:37
 *  @Description：测试defer关键字、recover
 */

package originGoLanguage

import (
	"fmt"
	"testing"
	"time"
)

func TestDefer(t *testing.T) {
	//defer会在所有函数执行完成之后才进行执行
	testDeferExe()
	//defer产生的值是在对应位置的值，后面的变化不会产生影响
	fmt.Println(testDeferValue())
	//defer的执行是先进后出
	testDeferExeOrder()
	testDeferRecover()    //defer对异常的捕获
	testDeferRecoverOut() //外层异常捕获
}

func TestCtripException(t *testing.T) {
	fmt.Println("测试不进行捕获的情况")
	testNotDefer()

	fmt.Println("测试进行捕获的情况")
	testDefer1()
}

func testNotDefer() {
	go sayHello()

	for i := 0; i < 10; i++ {
		time.Sleep(5000)
		fmt.Println(i)
	}
}

func testDefer1() {
	go sayHello1()

	for i := 0; i < 10; i++ {
		time.Sleep(5000)
		fmt.Println(i)
	}
}

func sayHello1() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("错误信息", err)
		}
	}()
}

func sayHello() {
	var map1 map[string]int
	//这里会发生异常
	fmt.Println(map1["say"])
}

func testDeferRecover() {
	//多个panic只会捕获最后一个
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("错误信息是：", err)
		}
	}()

	panic1()
	panic2()
	panic3()

	//一旦错误被捕获就不会继续运行
	fmt.Println("我会继续运行")

}

func panic3() {
	panic("panic3")
}

func panic2() {
	panic("panic2")
}

func panic1() {
	panic("panic1")
}

func testDeferRecoverOut() {
	testDeferRecover()

	fmt.Println("外层继续执行")
}

func testDeferExe() {
	i := 0
	defer func() {
		fmt.Println(i)
	}()
	fmt.Println("i am after")
}

func testDeferValue() int {
	i := 0
	defer func() {
		fmt.Println("first: ", i)
	}()

	i++
	defer func() {
		fmt.Println("second: ", i)
	}()

	i++
	return i
}

func testDeferExeOrder() {
	defer func() {
		fmt.Println("first")
	}()

	defer func() {
		fmt.Println("second")
	}()

	defer fmt.Println("third")
}
