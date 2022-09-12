package base

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UploadHandler(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		fmt.Println(err)
	} else {
		log.Println(file.Filename)

		dst := "./" + file.Filename
		// 上传文件至指定的完整文件路径
		c.SaveUploadedFile(file, dst)

		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	}
}

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if username == "qiudaoyu" && password == "123" {
		c.JSON(http.StatusOK, gin.H{
			"errCode":  "0",
			"username": username,
			"password": password,
		})
	}

}

func Page404(c *gin.Context) {
	// c.HTML(http.StatusNotFound, "views/404.html", nil)
	c.Redirect(http.StatusMovedPermanently, "http://localhost:8080/page404")
}
