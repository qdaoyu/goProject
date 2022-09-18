package main

import (
	"fmt"
	"qiudaoyu/middleWare"
	"qiudaoyu/models"
	"qiudaoyu/routers/base"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	// 为 multipart forms 设置较低的内存限制 (默认是 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	//初始化数据库连接
	models.InitDb()
	router.Use(
		middleWare.JWTAuthMiddleware(),
		middleWare.CORSMiddleware(),
		middleWare.StatCost(),
	)
	base.LoadBase(router)

	// router.Use(middleWare.CORSMiddleware())
	// router.Use(middleWare.StatCost())

	if err := router.Run(":8082"); err != nil {
		fmt.Printf("startup service failed, err:%v\n", err)
	}
}
