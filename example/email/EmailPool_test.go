/*
 *  @Author : huangzj
 *  @Time : 2020/12/29 11:32
 *  @Description：
 */

package email

import (
	"fmt"
	"github.com/jordan-wright/email"
	"log"
	"net/smtp"
	"os"
	"sync"
	"testing"
	"time"
)

func TestEmailPool(t *testing.T) {
	ch := make(chan *email.Email, 4)
	//创建2个连接的邮件池
	p, err := email.NewPool("smtp.qq.com:587", 2, smtp.PlainAuth("", "你的邮箱@qq.com", "你的stmp的密码", "smtp.qq.com"))
	if err != nil {
		log.Fatal("failed to create pool:", err)
	}

	//进行邮件发送
	var wg sync.WaitGroup
	wg.Add(2)
	for i := 0; i < 2; i++ {
		go func() {
			defer wg.Done()
			for e := range ch {
				err := p.Send(e, 1*time.Second)
				if err != nil {
					_, _ = fmt.Fprintf(os.Stderr, "email:%v sent error:%v\n", e, err)
				}
			}
		}()
	}

	//发送四个邮件
	for i := 0; i < 4; i++ {
		e := email.NewEmail()
		e.From = "hzj <你的邮箱@qq.com>"
		e.To = []string{"hzj <接受者的邮箱@qq.com>"}
		e.Subject = "Awesome web"
		e.Text = []byte(fmt.Sprintf("Awesome Web %d", i+1))
		ch <- e
	}

	close(ch)
	wg.Wait()
}
