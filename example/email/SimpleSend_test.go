/*
 *  @Author : huangzj
 *  @Time : 2020/12/29 11:10
 *  @Description：
 */

package email

import (
	"log"
	"net/smtp"
	"testing"

	"github.com/jordan-wright/email"
)

func TestSimpleSend(t *testing.T) {
	e := email.NewEmail()
	e.From = "hzj <你的邮箱@qq.com>"
	e.To = []string{"接受者的邮箱@qq.com"}
	e.Subject = "测试"
	e.Text = []byte("Text Body is, of course, supported!")
	err := e.Send("smtp.qq.com:587", smtp.PlainAuth("", "你的邮箱@qq.com", "你的stmp的密码", "smtp.qq.com"))
	if err != nil {
		log.Fatal(err)
	}
}
