/*
 *  @Author : huangzj
 *  @Time : 2020/12/8 16:15
 *  @Description：
 */

package gin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"testing"
)

//路有组就是把相同前缀的路由路径分在一起。
func TestRouteGroup(t *testing.T) {
	router := gin.Default()

	// 简单的路由组: v1
	v1 := router.Group("/v1")
	{
		v1.POST("/login", func(context *gin.Context) {
			fmt.Println(context.Params)
		})
		v1.POST("/submit", func(context *gin.Context) {
			fmt.Println(context.Params)
		})
		v1.POST("/read", func(context *gin.Context) {
			fmt.Println(context.Params)
		})
	}

	// 简单的路由组: v2
	v2 := router.Group("/v2")
	{
		v2.POST("/login", func(context *gin.Context) {
			fmt.Println(context.Params)
		})
		v2.POST("/submit", func(context *gin.Context) {
			fmt.Println(context.Params)
		})
		v2.POST("/read", func(context *gin.Context) {
			fmt.Println(context.Params)
		})
	}

	_ = router.Run(":8080")
}
