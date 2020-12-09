/*
 *  @Author : huangzj
 *  @Time : 2020/12/8 9:46
 *  @Description：
 */

package gin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

type Test struct {
	Label *string
	Reps  []int64
}

func TestRender(t *testing.T) {
	r := gin.Default()

	byMap(r) //直接通过map返回

	byStructWithAlias(r) //通过结构体返回,并设置结构体别名

	byXml(r) //返回xml报文

	byYaml(r) //返回YAML报文

	byProtoBuf(r) //返回序列化的二进制数据

	// 监听并在 0.0.0.0:8080 上启动服务
	r.Run(":8080")
}

func byProtoBuf(r *gin.Engine) {
	r.GET("/someProtoBuf", func(c *gin.Context) {
		reps := []int64{int64(1), int64(2)}
		label := "test"
		// protobuf 的具体定义写在 testdata/protoexample 文件中。
		data := &Test{
			Label: &label,
			Reps:  reps,
		}
		// 请注意，数据在响应中变为二进制数据
		// 将输出被 protoexample.Test protobuf 序列化了的数据
		c.ProtoBuf(http.StatusOK, data)
	})
}

func byYaml(r *gin.Engine) {
	r.GET("/someYAML", func(c *gin.Context) {
		c.YAML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})
}

func byXml(r *gin.Engine) {
	r.GET("/someXML", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})
}

func byStructWithAlias(r *gin.Engine) {
	r.GET("/moreJSON", func(c *gin.Context) {
		// 你也可以使用一个结构体
		var msg struct {
			Name    string `json:"user"`
			Message string
			Number  int
		}
		msg.Name = "Lena"
		msg.Message = "hey"
		msg.Number = 123
		// 注意 msg.Name 在 JSON 中变成了 "user"
		// 将输出：{"user": "Lena", "Message": "hey", "Number": 123}
		c.JSON(http.StatusOK, msg)
	})
}

func byMap(r *gin.Engine) {
	// gin.H 是 map[string]interface{} 的一种快捷方式
	r.GET("/someJSON", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})
}
