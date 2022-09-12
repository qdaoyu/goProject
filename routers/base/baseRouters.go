package base

import (
	"github.com/gin-gonic/gin"
)

func LoadBase(e *gin.Engine) {

	e.POST("/upload", UploadHandler)

	e.POST("/login", Login)

	e.NoRoute(Page404)

}
