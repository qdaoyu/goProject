package base

import (
	"github.com/gin-gonic/gin"
)

func LoadBase(e *gin.Engine) {

	e.POST("/upload", UploadHandler)

	e.POST("/login", Login)

	e.POST("/calDupName", CalDupName)

	e.GET("/home", HomeMenuHandler)

	// e.POST("/auth", AuthHandler)

	e.NoRoute(Page404)

}
