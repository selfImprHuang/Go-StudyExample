/*
 *  @Author : huangzj
 *  @Time : 2020/12/29 11:19
 *  @Description： 测试发送html，包含抄送对象
 */

package email

import (
	"github.com/jordan-wright/email"
	"log"
	"net/smtp"
	"testing"
)

func TestSendHtml(t *testing.T) {
	e := email.NewEmail()
	e.From = "hzj <你的邮箱@qq.com>"
	e.To = []string{"hzj <接受者的邮箱@qq.com>"}
	//设置抄送目标
	e.Cc = []string{"hzj <抄送邮箱@qq.com>"}
	e.Subject = "参考：Go 每日一库"
	e.HTML = []byte(`
<H1>参考：Go 每日一库 的email库</H1>
	<ul>
<li><a "https://darjun.github.io/2020/01/10/godailylib/flag/">Go 每日一库之 flag</a></li>
<li><a "https://darjun.github.io/2020/01/10/godailylib/go-flags/">Go 每日一库之 go-flags</a></li>
<li><a "https://darjun.github.io/2020/01/14/godailylib/go-homedir/">Go 每日一库之 go-homedir</a></li>
<li><a "https://darjun.github.io/2020/01/15/godailylib/go-ini/">Go 每日一库之 go-ini</a></li>
<li><a "https://darjun.github.io/2020/01/17/godailylib/cobra/">Go 每日一库之 cobra</a></li>
<li><a "https://darjun.github.io/2020/01/18/godailylib/viper/">Go 每日一库之 viper</a></li>
<li><a "https://darjun.github.io/2020/01/19/godailylib/fsnotify/">Go 每日一库之 fsnotify</a></li>
<li><a "https://darjun.github.io/2020/01/20/godailylib/cast/">Go 每日一库之 cast</a></li>
</ul>
	`)
	err := e.Send("smtp.qq.com:587", smtp.PlainAuth("", "你的邮箱@qq.com", "你的stmp的密码", "smtp.qq.com"))
	if err != nil {
		log.Fatal(err)
	}
}
