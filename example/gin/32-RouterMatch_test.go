/*
 *  @Author : huangzj
 *  @Time : 2020/12/8 15:54
 *  @Description：
 */

package gin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

func TestRouterMatch(t *testing.T) {
	router := gin.Default()

	OnlyMatch(router) //只能匹配到指定格式

	MathMany(router) //*匹配多重字段

	_ = router.Run(":8080")
}

func MathMany(router *gin.Engine) {
	// 此 handler 将匹配 /user/john/ 和 /user/john/send
	// 如果没有其他路由匹配 /user/john，它将重定向到 /user/john/
	router.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})
}

func OnlyMatch(router *gin.Engine) {
	// 此 handler 将匹配 /user/john 但不会匹配 /user/ 或者 /user
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})
}
