/*
 *  @Author : huangzj
 *  @Time : 2020/12/8 11:58
 *  @Description：模型绑定和验证
 */

package gin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

func doc() string {
	return `
	要将请求体绑定到结构体中，使用模型绑定。 Gin 目前支持 JSON、XML、YAML 和标准表单值的绑定（foo=bar＆boo=baz）。
	
	Gin 使用 go-playground/validator.v8 进行验证。 查看标签用法的全部文档(https://github.com/go-playground/validator)

	使用时，需要在要绑定的所有字段上，设置相应的 tag。 例如，使用 JSON 绑定时，设置字段标签为 json:"fieldname"。
	Gin 提供了两类绑定方法：
		Type - Must bind
			Methods - Bind, BindJSON, BindXML, BindQuery, BindYAML
			Behavior - 这些方法属于 MustBindWith 的具体调用。 如果发生绑定错误，则请求终止，并触发 c.AbortWithError(400, err).SetType(ErrorTypeBind)。响应状态码被设置为 400 并且 Content-Type 被设置为 text/plain; charset=utf-8。 如果您在此之后尝试设置响应状态码，Gin 会输出日志 [GIN-debug] [WARNING] Headers were already written. Wanted to override status code 400 with 422。 如果您希望更好地控制绑定，考虑使用 ShouldBind 等效方法。
		
		Type - Should bind
			Methods - ShouldBind, ShouldBindJSON, ShouldBindXML, ShouldBindQuery, ShouldBindYAML
			Behavior - 这些方法属于 ShouldBindWith 的具体调用。 如果发生绑定错误，Gin 会返回错误并由开发者处理错误和请求。
		

	使用 Bind 方法时，Gin 会尝试根据 Content-Type 推断如何绑定。 如果你明确知道要绑定什么，可以使用 MustBindWith 或 ShouldBindWith。
	你也可以指定必须绑定的字段。 如果一个字段的 tag 加上了 binding:"required"，但绑定时是空值，Gin 会报错。
`
}

type Login struct {
	User     string `form:"user" json:"user" xml:"user"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func TestBindValid(t *testing.T) {
	router := gin.Default()
	BindJson(router) //对结构体中的json设置字段名进行绑定，如果没有对应字段则会报错

	BindXml(router) //对结构体中的Xml设置字段名进行绑定，如果没有对应字段则会报错

	BindHtml(router) //对html输入的地址栏参数进行绑定

	//请求示例
	_ = `$ curl -v -X POST \
		 http://localhost:8080/loginJSON \
			-H 'content-type: application/json' \
			-d '{ "user": "manu" }'
		> POST /loginJSON HTTP/1.1
		> Host: localhost:8080
		> User-Agent: curl/7.51.0
		> Accept: */*
		> content-type: application/json
		> Content-Length: 18
		>
		* upload completely sent off: 18 out of 18 bytes
		< HTTP/1.1 400 Bad Request
		< Content-Type: application/json; charset=utf-8
		< Date: Fri, 04 Aug 2017 03:51:31 GMT
		< Content-Length: 100
		<
		{"error":"Key: 'Login.Password' Error:Field validation for 'Password' failed on the 'required' tag"}
   `
	fmt.Println("上述为请求示例，因为在结构体的tag上设置了必须校验，所以当一个字段没有进行输入的时候，就会报错，除非是去掉对应的tag设置或者是传输正确的字段")
}

func BindHtml(router *gin.Engine) {
	// 绑定 HTML 表单 (user=manu&password=123)
	router.POST("/loginForm", func(c *gin.Context) {
		var form Login
		// 根据 Content-Type Header 推断使用哪个绑定器。
		if err := c.ShouldBind(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if form.User != "manu" || form.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})
}

func BindXml(router *gin.Engine) {
	// 绑定 XML (
	//  <?xml version="1.0" encoding="UTF-8"?>
	//  <root>
	//      <user>user</user>
	//      <password>123</password>
	//  </root>)
	router.POST("/loginXML", func(c *gin.Context) {
		var xml Login
		if err := c.ShouldBindXML(&xml); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if xml.User != "manu" || xml.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})
}

func BindJson(router *gin.Engine) {
	// 绑定 JSON ({"user": "manu", "password": "123"})
	router.POST("/loginJSON", func(c *gin.Context) {
		var json Login
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if json.User != "manu" || json.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})

}
