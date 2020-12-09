/*
 *  @Author : huangzj
 *  @Time : 2020/12/8 15:12
 *  @Description：
 */

package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"net/http"
	"testing"
	"time"
)

// Booking 包含绑定和验证的数据。
type Booking struct {
	CheckIn  time.Time `form:"check_in" binding:"required,diyValid" time_format:"2006-01-02"`
	CheckOut time.Time `form:"check_out" binding:"required,gtfield=CheckIn" time_format:"2006-01-02"`
}

//read note 自定义的校验类，需要实现 validator.Func方法(参数是validator.FieldLevel,这里把所有的参数都集合进去了，之前的版本是多个参数)
func bookDataValid(level validator.FieldLevel) bool {
	if date, ok := level.Field().Interface().(time.Time); ok {
		today := time.Now()
		if today.Year() > date.Year() || today.YearDay() > date.YearDay() {
			return false
		}
	}

	return true
}

func TestDiyStructValid(t *testing.T) {
	route := gin.Default()

	//read note 注册自定义的校验类,这边设置的key就是需要在结构体上绑定数据时候书写的tag标签，注意看上面Booking 的CheckIn标签(只有在结构体上书写了该标签，这个校验器才会去校验对应的字段)
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("diyValid", bookDataValid)
	}

	route.GET("/bookable", getBookable)
	_ = route.Run(":8085")

	//请求示例
	_ = `
 		$ curl "localhost:8085/bookable?check_in=2018-04-16&check_out=2018-04-17"
		{"message":"Booking dates are valid!"}
		
		$ curl "localhost:8085/bookable?check_in=2018-03-08&check_out=2018-03-09"
		{"error":"Key: 'Booking.CheckIn' Error:Field validation for 'CheckIn' failed on the 'diyValid' tag"}
	`
}

func getBookable(c *gin.Context) {
	var b Booking
	if err := c.ShouldBindWith(&b, binding.Query); err == nil {
		c.JSON(http.StatusOK, gin.H{"message": "Booking dates are valid!"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
