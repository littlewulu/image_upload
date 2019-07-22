package main

import (
	"github.com/gin-gonic/gin"
	"image_upload/config"
	"image_upload/controller"
	"math/rand"
	"time"
)

func init(){
	// 初始化随机数的种子
	rand.Seed(time.Now().Unix())

	if !config.ConfigInstance.IsDebug{
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()

	if config.ConfigInstance.IsDebug{
		router.Use(gin.Logger())
	}

	router.Use(gin.Recovery())

	router.POST("/api/tomato/image/upload_image", controller.UploadImageFileFunc)


	router.Run(config.ConfigInstance.RestServerAddr)
}


func main(){


}

