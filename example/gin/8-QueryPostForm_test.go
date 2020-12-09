/*
 *  @Author : huangzj
 *  @Time : 2020/12/8 9:39
 *  @Description：
 */

package gin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"testing"
)

//read note POST /post?id=1234&page=1 HTTP/1.1
// Content-Type: application/x-www-form-urlencoded
// name=manu&message=this_is_great

//read note Query是用来解析地址栏中的参数
// PostForm是用来解析报文体中的参数
func TestQueryPostForm(t *testing.T) {
	router := gin.Default()

	router.POST("/post", func(c *gin.Context) {

		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		name := c.PostForm("name")
		message := c.PostForm("message")

		fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
	})
	router.Run(":8080")
}
