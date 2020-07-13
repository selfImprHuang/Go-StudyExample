/*
 *  @Author : huangzj
 *  @Time : 2020/4/28 17:17
 *  @Description：
 */

package main

import (
	"Go-StudyExample/example"
	"Go-StudyExample/example/originGoLanguage"
	_ "Go-StudyExample/example/originGoLanguage/init"
	"Go-StudyExample/example/problem"
)

func main() {
	//--------------------------测试mapStructure的功能-----
	//这边的四种使用方式差别感觉不是很大...
	example.MapStructureTestFunc()
	example.MapStructureTestFunc1()
	example.MapStructureTestFunc2()
	example.MapStructureTestFunc3()

	////--------------------------测试json包的转换功能
	example.JsonMarshalTest()
	example.JsonUnmarshalTest()

	////克隆测试
	example.TestClone()

	////二维码测试
	example.QrCodeTest()

	//图片测试
	example.ImageTest()

	//go原生语言测试
	originGoLanguage.TestOriginLang()

	//矩阵旋转
	problem.MatrixRotationTest()
}
