/*
 *  @Author : huangzj
 *  @Time : 2020/12/8 9:18
 *  @Description：
 */

package gin

import (
	"fmt"
	"net/http"
	"testing"
)
import "github.com/gin-gonic/gin"

func TestAsciiJson(t *testing.T) {
	r := gin.Default()

	r.GET("/someJSON", func(c *gin.Context) {
		data := map[string]interface{}{
			"lang": "GO语言",
			"tag":  "<br>",
		}

		// 输出 : {"lang":"GO\u8bed\u8a00","tag":"\u003cbr\u003e"}
		c.AsciiJSON(http.StatusOK, data)
	})

	// 监听并在 0.0.0.0:8080 上启动服务
	err := r.Run(":8080")
	if err != nil {
		fmt.Println(err.Error())
	}
}
