/*
 *  @Author : huangzj
 *  @Time : 2020/12/8 11:10
 *  @Description：
 */

package gin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"testing"
)

//上传命令，指定多文件数组
//  curl -X POST http://localhost:8080/upload \
//  	-F "upload[]=@/Users/appleboy/test1.zip" \
//  	-F "upload[]=@/Users/appleboy/test2.zip" \
//  	-H "Content-Type: multipart/form-data"
func TestMultiFile(t *testing.T) {
	router := gin.Default()
	// 为 multipart forms 设置较低的内存限制 (默认是 32 MiB)
	// router.MaxMultipartMemory = 8 << 20  // 8 MiB
	router.POST("/upload", func(c *gin.Context) {
		// Multipart form
		form, _ := c.MultipartForm()
		fileHeaders := form.File["upload[]"]

		for _, header := range fileHeaders {
			log.Println(header.Filename)
			file, _ := header.Open()
			var content []byte
			file.Read(content) //读取文件

			// c.SaveUploadedFile(file, dst) 	//上传文件至指定目录
		}
		c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(fileHeaders)))
	})
	router.Run(":8080")
}
