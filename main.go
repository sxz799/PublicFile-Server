package main

import (
	"PublicFileServer/gobalConfig"
	"PublicFileServer/model"
	"PublicFileServer/router"
	"PublicFileServer/util"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron"
	"log"
)

func main() {
	log.Println("正在连接数据库...")
	util.InitDB()
	log.Println("正在检查表结构...")
	model.InitAutoMigrateDB()
	gin.SetMode(gobalConfig.GinMode)
	r := gin.Default()
	if gobalConfig.FrontMode {
		log.Println("已开启前后端整合模式！")
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
	log.Println("定时任务启动成功,服务启动成功,当前使用端口：", gobalConfig.ServerPort)
	r.Run(":" + gobalConfig.ServerPort)
}
