package main

import (
	"PublicFileServer/model"
	"PublicFileServer/router"
	"PublicFileServer/util"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron"
	"github.com/spf13/viper"
)

func main() {
	util.InitDB()
	model.InitAutoMigrateDB()
	r := gin.Default()
	//r.Use(router.Cors())
	//gin.SetMode(gin.ReleaseMode)
	router.RegRouter(r)
	c := cron.New()
	c.AddFunc("@every 10m", model.DelFile)
	c.Start()
	r.Run(":" + viper.GetString("server.port"))
}
