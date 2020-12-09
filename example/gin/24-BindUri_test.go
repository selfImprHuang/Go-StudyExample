/*
 *  @Author : huangzj
 *  @Time : 2020/12/8 14:06
 *  @Description：直接绑定Uri参数
 */

package gin

import (
	"github.com/gin-gonic/gin"
	"testing"
)

type Person struct {
	ID   string `uri:"id" binding:"required,uuid"`
	Name string `uri:"name" binding:"required"`
}

func TestBindUri(t *testing.T) {
	route := gin.Default()
	//这边的【:name】和【:id】说明了对应的该Uri，传入的参数分别是叫做name和id.在下面的Bind直接进行映射
	route.GET("/uri/:name/:id", func(c *gin.Context) {
		var person Person
		if err := c.ShouldBindUri(&person); err != nil {
			c.JSON(400, gin.H{"msg": err})
			return
		}
		c.JSON(200, gin.H{"name": person.Name, "uuid": person.ID})
	})
	route.Run(":8088")

	//请求示例
	_ = `
        这边可以看到uri后面的两个参数就是 name 和 id,直接不用声明名称.
 		$ curl -v localhost:8088/uri/thinkerou/987fbc97-4bed-5078-9f07-9141ba07c9f3
 		$ curl -v localhost:8088/uri/thinkerou/not-uuid
	`
}
