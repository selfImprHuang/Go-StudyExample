/*
 *  @Author : huangzj
 *  @Time : 2020/12/8 15:27
 *  @Description：
 */

package gin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"testing"
)

func TestReadCookie(t *testing.T) {

	router := gin.Default()

	router.GET("/cookie", func(c *gin.Context) {

		//read note 这边需要web前端发送对应的cookie
		cookie, err := c.Cookie("gin_cookie")

		//read note 设置对应值的cookie
		if err != nil {
			cookie = "NotSet"
			c.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)
		}
		fmt.Printf("Cookie value: %s \n", cookie)
	})
	_ = router.Run()
}
