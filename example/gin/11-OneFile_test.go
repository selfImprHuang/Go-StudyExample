/*
 *  @Author : huangzj
 *  @Time : 2020/12/8 10:51
 *  @Description：单文件上传
 */

package gin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"testing"
)

func TestOneFile(t *testing.T) {
	router := gin.Default()
	// 为 multipart forms 设置较低的内存限制 (默认是 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.POST("/upload", func(c *gin.Context) {
		// 单文件
		fileHeader, _ := c.FormFile("file")
		log.Println(fileHeader.Filename)

		var content []byte
		file, _ := fileHeader.Open()
		defer file.Close()
		_, _ = file.Read(content)
		fmt.Println(string(content)) //输出文件内容

		//上传文件至指定目录
		err := c.SaveUploadedFile(fileHeader, "C://Users//admin//Desktop")
		if err != nil {
			panic(err)
		}
		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", fileHeader.Filename))
	})
	router.Run(":8080")
}
