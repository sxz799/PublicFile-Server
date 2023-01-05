package gobalConfig

import (
	"github.com/spf13/viper"
	"log"
	"os"
)

var (
	LimitFileLife   int
	LimitFileSize   int
	ShareCodeType   int
	ShareCodeLength int
	ServerPort      string
	FrontMode       bool
	GinMode         string
)

func init() {
	log.Println("正在检查files文件夹")
	_, err := os.Stat("files")
	if err != nil && os.IsNotExist(err) {
		os.Mkdir("files", os.ModePerm)
	}
}

func init() {
	log.Println("正在应用配置文件...")
	viper.SetConfigName("conf")
	viper.SetConfigType("yml")
	viper.AddConfigPath("conf")
	err := viper.ReadInConfig()
	if err != nil {
		log.Panicln("viper load fail ...")
		return
	}

	ServerPort = viper.GetString("server.port")
	GinMode = viper.GetString("server.ginMode")
	LimitFileLife = viper.GetInt("config.limitFileLife")
	LimitFileSize = viper.GetInt("config.limitFileSize")
	ShareCodeType = viper.GetInt("config.shareCodeType")
	ShareCodeLength = viper.GetInt("config.shareCodeLength")
	FrontMode = viper.GetBool("config.frontMode")
}
