/*
 *  @Author : huangzj
 *  @Time : 2020/7/3 12:00
 *  @Description：参考文章：https://zhuanlan.zhihu.com/p/32035735
 */

package example

import (
	"Go-StudyExample/example"
	"fmt"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	qrCode "github.com/skip2/go-qrcode"
	"image/color"
	"image/png"
	"log"
	"os"
	"testing"
)

const (
	url1 = "https://image.baidu.com/search/detail?ct=503316480&z=undefined&tn=baiduimagedetail&ipn=d&word=%E7%BE%8E%E5%A5%B3%E5%9B%BE%E7%89%87&step_word=&ie=utf-8&in=&cl=2&lm=-1&st=undefined&hd=undefined&latest=undefined&copyright=undefined&cs=1848592761,1307412614&os=1322849054,135797925&simid=0,0&pn=644&rn=1&di=41830&ln=2454&fr=&fmq=1593763755230_R&fm=&ic=undefined&s=undefined&se=&sme=&tab=0&width=undefined&height=undefined&face=undefined&is=0,0&istype=0&ist=&jit=&bdtype=11&spn=0&pi=0&gsm=258&hs=2&objurl=http%3A%2F%2Fup.enterdesk.com%2Fedpic%2Fb2%2F9d%2Fac%2Fb29dacffee2403e8b3bec67c9abcd647.jpg&rpstart=0&rpnum=0&adpicid=0&force=undefined"
	url2 = "https://www.baidu.com/link?url=oH5FRzYYHiZNm84fp0CaQtGZgw2-WRZgIj7B3Mmmzy0iM_EcVHkHmSHQwU6LuPBi6o24N2ObdLJEL1l6aIY5KeXia3S-9ysVL5K7Qe_SnHeXNOxqbI3UbpSBktjFfbIunH9Y9OteFNO-YYsx2Nh1YbOVrZhzh0Ne26W3_73xgsX7ys418xQJOGiGteVwk6UaEqPaQgqlUtSHBCSjH_FZuJmxswzP-5X5Hg894diQSm8136aulR3oC7AB0LFSmCZoIviAdOGagItOAqIdTP_P7NWn7rzebKAFoS0j-q8fMLvoZJYK1xTir1-scoBZ1aTSdEsrhTcpKzhHyG5CarDTj3BdmD96JzBppy9rDMXQ8XstX8gzJb4LrTCdg5t9RMK11wzjmgJhkj-9deb7iZqsdVDBgUpur-UxR6vsa5OvXkgLd4Tzv1fTm3-UjdYgfPDUQ9pGGLzLQQKO8i4eEEFsrSJZCx4eAdD94GBbOOVpDZtwJSwVqLJKvPtu8h0gt-lOmktaHxJffIXZfBKtfMfCnc_fqomMQAGoaaIkehD9X57zeK6OZg1f3kM7Azv5pzs2ITgaErADeN0vdQbZ8fWBT0tqTmmDNRJLgzVnBx9V9FMhgD6eAOHunHxHAT5g_DMr&timg=https%3A%2F%2Fss0.bdstatic.com%2F94oJfD_bAAcT8t7mm9GUKT-xh_%2Ftimg%3Fimage%26quality%3D100%26size%3Db4000_4000%26sec%3D1593763752%26di%3D3a1f64455c44b927fc36a894140e279e%26src%3Dhttp%3A%2F%2F00.minipic.eastday.com%2F20170420%2F20170420105628_ea6da92abc46098d8e03ad2ee55abeb7_9.jpeg&click_t=1593763828401&s_info=2543_1335&wd=&eqid=916eb61300000776000000065efee7a7"
	url3 = "http://note.youdao.com/noteshare?id=de3af43fd227f9a23e23f50d0dab105f"

	qrcode1 = "example/image/log_qrcode.png"
	qrcode2 = "example/image/log_qrcode1.png"
)

func TestQrCode(t *testing.T) {

	fmt.Printf("通过github.com/skips/go-qrcode 实现生成二维码")

	//生成黑白二维码
	err := qrCode.WriteFile(url1, qrCode.High, 255, qrcode1)
	if err != nil {
		panic(err)
	}

	//生成字节数组信息
	qrBytes, err := qrCode.Encode(url3, qrCode.High, 255)
	for _, r := range qrBytes {
		fmt.Println(r)
	}

	//自定义二维码信息
	qrCode1, err := qrCode.New(url2, qrCode.Medium)
	if err != nil {
		log.Fatal(err)
	} else {
		qrCode1.BackgroundColor = color.RGBA{250, 250, 50, 255} //设置背景色
		qrCode1.ForegroundColor = color.Black                   //设置二维码内背景色
		_ = qrCode1.WriteFile(256, qrcode2)
	}

	//------------------------------------------------------------------------------------------------

	fmt.Printf("通过 github.com/boombuler/barcode 实现生成二维码")
	content := "https://zhuanlan.zhihu.com/p/59125443" //二维码内容信息

	code, err2 := qr.Encode(content, qr.L, qr.Unicode) //对二维码进行编码
	example.Assert(err2)

	img, err3 := barcode.Scale(code, 300, 300) //图片大小设置
	example.Assert(err3)

	file, err4 := os.Create("F:/Go_BySelf/src/Go-StudyExample/example/image/qr_code.png") //创建文件
	example.Assert(err4)

	err5 := png.Encode(file, img) //图片编码生成
	//err6 := jpeg.Encode(file, img, &jpeg.Options{100}) //图像质量值为100，是最好的图像显示
	example.Assert(err5)

	err6 := file.Close()
	example.Assert(err6)

	//------------------------------------------------------------------------------------------------

}
