package main

import (
	"PublicFileServer/model"
	"PublicFileServer/router"
	"PublicFileServer/util"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	util.InitDB()
	model.InitAutoMigrateDB()
	r := gin.Default()
	//gin.SetMode(gin.ReleaseMode)
	router.RegRouter(r)
	r.Run(":" + viper.GetString("server.port"))
}
