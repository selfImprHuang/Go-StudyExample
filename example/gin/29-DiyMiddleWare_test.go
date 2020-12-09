/*
 *  @Author : huangzj
 *  @Time : 2020/12/8 14:45
 *  @Description：
 */

package gin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"testing"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// 请求前
		fmt.Println(fmt.Sprintf("记录日志，本次请求的参数为：%v", c.Params))

		c.Next()

		// 请求后
		latency := time.Since(t)
		log.Print(latency)

		// 获取发送的 status
		status := c.Writer.Status()
		log.Println(status)
	}
}

//在请求前后者请求后自定义中间件处理
func TestMiddleWare(t *testing.T) {
	router := gin.Default() //Default会返回两个中间件 Logger(), Recovery().这接口调用之前和之后进行处理，所以我们也可以自定义中间件

	router.Use(Logger())

	router.GET("/test", func(c *gin.Context) {
		example := c.MustGet("example").(string)

		// 打印："12345"
		log.Println(example)
	})

	// 监听并在 0.0.0.0:8080 上启动服务
	router.Run(":8080")
}
