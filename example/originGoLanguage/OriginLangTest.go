/*
 *  @Author : huangzj
 *  @Time : 2020/7/13 14:35
 *  @Description：
 */

package originGoLanguage

func TestOriginLang() {

	//网络及端口测试
	NetTest()

	//传值和传址的比较
	TestQuote()

	//切片测试
	TestSlice()

	//测试Map
	TestMap()

	//通道测试
	TestChannel()

	//测试init，只要在当前结构体中有引入对应的包，就会调用相应的init方法，具体可以查看 上面的import

	//测试error的生成
	TestError()

	//测试defer关键字
	TestDefer()

	//测试携程的异常捕获 --- 好像有没有捕获，主函数都可以正常运行
	TestCtripException()

}
