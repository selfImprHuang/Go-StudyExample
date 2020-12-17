/*
 *  @Author : huangzj
 *  @Time : 2020/12/17 10:14
 *  @Description： 参考地址：http://www.topgoer.com/%E5%B8%B8%E7%94%A8%E6%A0%87%E5%87%86%E5%BA%93/log.html
 */

package originGoLanguage

import (
	"fmt"
	"log"
	"os"
	"testing"
)

//日志的Flag说明：
//const (
//	// 控制输出日志信息的细节，不能控制输出的顺序和格式。
//	// 输出的日志在每一项后会有一个冒号分隔：例如2009/01/23 01:23:23.123123 /a/b/c/d.go:23: message
//	Ldate         = 1 << iota     // 日期：2009/01/23
//	Ltime                         // 时间：01:23:23
//	Lmicroseconds                 // 微秒级别的时间：01:23:23.123123（用于增强Ltime位）
//	Llongfile                     // 文件全路径名+行号： /a/b/c/d.go:23
//	Lshortfile                    // 文件名+行号：d.go:23（会覆盖掉Llongfile）
//	LUTC                          // 使用UTC时间
//	LstdFlags     = Ldate | Ltime // 标准logger的初始值
//)
func TestLogOut(t *testing.T) {
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.SetPrefix("[我是日志的前缀]") //配置日志的前缀
	log.Println("这是一条很普通的日志。")
}

func TestLogOutInFile(t *testing.T) {
	logFile, err := os.OpenFile("test.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open log file failed, err:", err)
		return
	}
	log.SetOutput(logFile)
	log.SetPrefix("[我是日志的前缀]") //配置日志的前缀
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.Println("我是一个写到文件中的日志")
}

func TestNewLogger(t *testing.T) {
	logger := log.New(os.Stdout, "<自定义Logger对象>", log.Lshortfile|log.Ldate|log.Ltime) //第二个参数是前缀，第三个参数是
	logger.Println("这是自定义的logger记录的日志。")

	logFile, err := os.OpenFile("test.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open log file failed, err:", err)
		return
	}
	logger.SetOutput(logFile)
	logger.Println("自定义Logger对象输出到文件")
}
