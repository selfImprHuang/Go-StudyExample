/*
 *  @Author : huangzj
 *  @Time : 2020/12/29 11:24
 *  @Description：
 */

package email

import (
	"github.com/jordan-wright/email"
	"log"
	"net/smtp"
	"testing"
)

func TestSendAnnex(t *testing.T) {
	e := email.NewEmail()
	e.From = "hzj <你的邮箱@qq.com>"
	e.To = []string{"hzj <接受者的邮箱@qq.com>"}
	e.Subject = "附件"
	e.Text = []byte("附件")
	_, _ = e.AttachFile("囚徒健身2翻译版 修改.pdf")
	err := e.Send("smtp.qq.com:587", smtp.PlainAuth("", "你的邮箱@qq.com", "你的stmp的密码", "smtp.qq.com"))
	if err != nil {
		log.Fatal(err)
	}
}
