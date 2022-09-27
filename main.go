package main

import (
	"PublicFileServer/model"
	"PublicFileServer/router"
	"PublicFileServer/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron"
	"github.com/spf13/viper"
	"log"
	"os"
)

func init() {
	viper.SetConfigName("conf")
	viper.SetConfigType("yml")
	viper.AddConfigPath("conf")
	err := viper.ReadInConfig()
	if err != nil {
		log.Panicln("viper load fail ...")
		return
	}
	util.InitDB()
	model.InitAutoMigrateDB()
	_, err2 := os.Stat("files")
	if err2 != nil && os.IsNotExist(err2) {
		os.Mkdir("files", os.ModePerm)
	}
}

func main() {
	r := gin.Default()
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
