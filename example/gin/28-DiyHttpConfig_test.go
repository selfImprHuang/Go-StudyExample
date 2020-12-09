/*
 *  @Author : huangzj
 *  @Time : 2020/12/8 14:14
 *  @Description：自定义Http配置
 */

package gin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
	"time"
)

func TestHttpConfig(t *testing.T) {
	router := gin.Default()

	//最重要的是可以通过自定义Handler来处理不同类型的请求，这边的请求是说监听的不同端口.
	//而不同的请求方法，可以处理对应的不同路径的请求，可以把一个类型的路径用一个相应的方法进行处理.
	s := &http.Server{
		Addr:           ":8080",          //监听地址
		Handler:        router,           //处理类可自定义
		ReadTimeout:    10 * time.Second, //读超时
		WriteTimeout:   10 * time.Second, //写超时
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
