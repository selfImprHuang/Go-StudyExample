/*
 *  @Author : huangzj
 *  @Time : 2020/4/28 17:17
 *  @Description：
 */

package main

import "Go-StudyExample/example"

func main() {
	//--------------------------测试mapStructure的功能-----
	//这边的四种使用方式差别感觉不是很大...
	example.MapStructureTestFunc()
	example.MapStructureTestFunc1()
	example.MapStructureTestFunc2()
	example.MapStructureTestFunc3()

	//--------------------------测试json包的转换功能
	example.JsonMarshalTest()
	example.JsonUnmarshalTest()
}
