package main

import (
	"fmt"
	"qiudaoyu/middleWare"
	"qiudaoyu/routers/base"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	base.LoadBase(router)

	// 为 multipart forms 设置较低的内存限制 (默认是 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.Use(middleWare.CORSMiddleware())
	router.Use(middleWare.StatCost())

	if err := router.Run(":8082"); err != nil {
		fmt.Printf("startup service failed, err:%v\n", err)
	}
}
