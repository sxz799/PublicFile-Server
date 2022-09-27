package main

import (
	"PublicFileServer/model"
	"PublicFileServer/router"
	"PublicFileServer/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron"
	"github.com/spf13/viper"
)

func main() {
	util.InitDB()
	model.InitAutoMigrateDB()
	r := gin.Default()
	//r.Use(router.Cors())
	gin.SetMode(gin.ReleaseMode)
	if viper.GetBool("server.frontMode") {
		fmt.Println("已开启前后端整合模式！")
		r.LoadHTMLGlob("static/index.html")
		r.Static("/static", "static")
		r.GET("/", func(context *gin.Context) {
			context.HTML(200, "index.html", "")
		})
	}
	router.RegRouter(r)
	c := cron.New()
	c.AddFunc("@every 10m", model.DelFile)
	c.Start()
	r.Run(":" + viper.GetString("server.port"))
}
