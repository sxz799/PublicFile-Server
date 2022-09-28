package main

import (
	"PublicFileServer/gobalConfig"
	"PublicFileServer/model"
	"PublicFileServer/router"
	"PublicFileServer/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron"
	"log"
)

func main() {
	log.Println("正在连接数据库...")
	util.InitDB()
	log.Println("正在检查表结构...")
	model.InitAutoMigrateDB()
	r := gin.Default()
	//gin.SetMode(gin.ReleaseMode)
	if gobalConfig.FrontMode {
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
	r.Run(":" + gobalConfig.ServerPort)
}
