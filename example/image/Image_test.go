/*
 *  @Author : huangzj
 *  @Time : 2020/7/6 10:42
 *  @Description：参考地址：https://www.cnblogs.com/dfsxh/articles/11082359.html
 */

package example

import (
	"fmt"
	"github.com/golang/freetype"
	"github.com/nfnt/resize"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

const (
	filePath = "example/image/image.jpg" //原始logo图片路径
	logoPath = "example/image/log.png"   //原始图片路径

	resizeImagePath  = "example/image/resize.png"  //重置尺寸图片的路径
	resizeImagePath2 = "example/image/resize2.png" //重置尺寸图片的路径
	resizeWidth      = 1000                        //重置尺寸图片的宽度
	resizeHeight     = 300                         //重置尺寸图片的高度

	waterMarkPath = "example/image/waterMakerImage.png" //携带水印的图片
	wordPath      = "example/image/wordPath.png"        //携带文字的图片
)

func TestImage(t *testing.T) {

	AddWordToImg() //设置文字到图片

	//图片压缩的两种方式，压缩率好像差不多，对应的尺寸比例不一样
	ResizeImage()
	ResizeImage2()

	AddWaterMark() //添加水印
}

//图片压缩
func ResizeImage() {
	// 打开图片并解码
	fileOrigin, _ := os.Open(filePath)
	originFile, _ := os.Stat(filePath)
	origin, err := jpeg.Decode(fileOrigin) //这边这个方法之前尝试使用截图的数据来做，会报错：missing SOI marker，似乎是因为base64编码的问题
	defer fileOrigin.Close()

	canvas := resize.Resize(resizeWidth, resizeHeight, origin, resize.Lanczos3) //尺寸重置的画布对象
	fileOut, err := os.Create(resizeImagePath)                                  //创建重置大小后的图片文件
	if err != nil {
		log.Fatal(err)
	}
	defer fileOut.Close()
	_ = jpeg.Encode(fileOut, canvas, &jpeg.Options{Quality: 80}) //编码对应的重置后图片文件（Options中的参数代表图片质量）
	targetFile, _ := os.Stat(resizeImagePath)
	fmt.Println(fmt.Sprintf("原始图片长度%d,当前图片长度%d", originFile.Size(), targetFile.Size()))
}

//图片压缩2（这种方式似乎会按照图片的比例进行缩放-取较小的长或宽的比例进行缩放，可以尝试调整一下resize的长度来对比结果）
func ResizeImage2() {
	// 打开图片并解码
	fileOrigin, _ := os.Open(filePath)
	originFile, _ := os.Stat(filePath)
	origin, err := jpeg.Decode(fileOrigin) //这边这个方法之前尝试使用截图的数据来做，会报错：missing SOI marker，似乎是因为base64编码的问题
	defer fileOrigin.Close()

	canvas := resize.Thumbnail(resizeWidth, resizeHeight, origin, resize.Lanczos3) //尺寸重置的画布对象
	fileOut, err := os.Create(resizeImagePath2)                                    //创建重置大小后的图片文件
	if err != nil {
		log.Fatal(err)
	}
	defer fileOut.Close()
	_ = jpeg.Encode(fileOut, canvas, &jpeg.Options{Quality: 100}) //编码对应的重置后图片文件（Options中的参数代表图片质量）
	targetFile, _ := os.Stat(resizeImagePath2)
	fmt.Println(fmt.Sprintf("原始图片长度%d,当前图片长度%d", originFile.Size(), targetFile.Size()))
}

//在图片上添加水印
func AddWaterMark() {
	// 打开图片并解码
	fileOrigin, _ := os.Open(filePath)
	origin, _ := jpeg.Decode(fileOrigin)
	defer fileOrigin.Close()

	// 打开水印图并解码
	fileWatermark, _ := os.Open(logoPath)
	watermark, _ := png.Decode(fileWatermark)
	defer fileWatermark.Close()
	originSize := origin.Bounds() //原始图界限

	canvas := image.NewNRGBA(originSize) //创建新图层

	draw.Draw(canvas, originSize, origin, image.ZP, draw.Src) // 贴原始图

	draw.Draw(canvas, watermark.Bounds().Add(image.Pt(originSize.Dx()/2, originSize.Dy()/2)), watermark, image.ZP, draw.Over) // 贴水印图

	//生成新图片
	createImage, _ := os.Create(waterMarkPath)
	_ = jpeg.Encode(createImage, canvas, &jpeg.Options{95})
	defer createImage.Close()
}

type RGBA struct { // Pix保管图像的像素色彩信息，顺序为R, G, B, A
	// 像素(x, y)起始位置是Pix[(y-Rect.Min.Y)*Stride + (x-Rect.Min.X)*4]
	Pix []uint8
	// Stride是Pix中每行像素占用的字节数
	Stride int
	// Rect是图像的范围
	Rect image.Rectangle
}

const (
	fontPath = "F://Go_BySelf//src//Go-StudyExample//example//image//微软vista黑体.ttf" //字体的存储位置
	content  = "i love 许小baby"
)

//在图片上添加文字
//使用到的包：go get github.com/golang/freetype
//参考地址;http://www.sunaloe.cn/d/36.html
func AddWordToImg() {
	img := copyBackgroundImage() //从图片中复制出图层
	generateContentFont(img)     //设置字体和Content
	createImage, _ := os.Create(wordPath)
	_ = jpeg.Encode(createImage, img, &jpeg.Options{Quality: 95})
	defer createImage.Close()
}

func generateContentFont(img *image.NRGBA) {
	//读取字体数据
	fontBytes, err := ioutil.ReadFile(fontPath)
	checkErr(err)

	//载入字体数据
	font, err := freetype.ParseFont(fontBytes)
	checkErr(err)

	f := freetype.NewContext()
	//设置分辨率
	f.SetDPI(72)
	//设置字体
	f.SetFont(font)
	//设置尺寸
	f.SetFontSize(100)
	f.SetClip(img.Bounds())
	//设置输出的图片
	f.SetDst(img)
	//设置字体颜色
	f.SetSrc(image.NewUniform(color.RGBA{R: 12, G: 225, B: 12, A: 99}))

	//设置字体的位置
	pt := freetype.Pt(500, 500+int(f.PointToFixed(26))>>8)

	_, err = f.DrawString(content, pt)
	checkErr(err)
}

func copyBackgroundImage() *image.NRGBA {
	imgFile, _ := os.Open(filePath)
	defer imgFile.Close()
	//获取到底部的照片作为背景，也可以理解为是在底部的照片上面写入文字
	backGround, _, _ := image.Decode(imgFile)
	dx := backGround.Bounds().Size().X
	dy := backGround.Bounds().Size().Y

	img := image.NewNRGBA(image.Rect(0, 0, dx, dy)) //初始化新生成图片

	//通过循环像素点把图片复制到新的图层.作为新图层的背景，这边其实也可以直接设置非图片的背景图
	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			img.Set(x, y, backGround.At(x, y))
		}
	}

	return img
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
