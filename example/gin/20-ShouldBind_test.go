/*
 *  @Author : huangzj
 *  @Time : 2020/12/8 11:33
 *  @Description：
 */

package gin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"testing"
)

type formA struct {
	Foo string `json:"foo" xml:"foo" binding:"required"`
}

type formB struct {
	Bar string `json:"bar" xml:"bar" binding:"required"`
}

func TestBind(t *testing.T) {
	router := gin.Default()

	bindOnce(router) //只能绑定一次
	bindMore(router) //可以多次绑定
}

func bindMore(router *gin.Engine) {
	objA := formA{}
	objB := formB{}

	// 读取 c.Request.Body 并将结果存入上下文,所以下一次的读取可以从上下文中拿到，达到重复读取的目的
	// 会对性能造成轻微影响，如果调用一次就能完成绑定的话，那就不要用这个方法。
	router.GET("bindMany", func(context *gin.Context) {
		if errA := context.ShouldBindBodyWith(&objA, binding.JSON); errA == nil {
			fmt.Println(objA.Foo)
			//绑定不同的格式，这边演示xml的格式进行绑定
		} else if errB := context.ShouldBindBodyWith(&objB, binding.XML); errB == nil {
			fmt.Println(objB.Bar)
		}
	})
}

func bindOnce(router *gin.Engine) {
	objA := formA{}
	objB := formB{}

	//c.ShouldBind 使用了 c.Request.Body，不可重用。请求是数据流，没有存储下来则只能使用一次
	router.GET("bind", func(context *gin.Context) {
		//绑定Body数据到结构体.
		if errA := context.ShouldBind(&objA); errA == nil {
			fmt.Println(objA.Foo)
		} else if errB := context.ShouldBind(&objB); errB == nil {
			fmt.Println(objB.Bar)
		}
	})
}
