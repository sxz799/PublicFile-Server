package model

import (
	"PublicFileServer/util"
	"log"
)

func InitAutoMigrateDB() {
	err := util.DB.AutoMigrate(FileObj{})
	if err != nil {
		log.Println("FileObj表创建失败")
	}
	err2 := util.DB.AutoMigrate(SystemLog{})
	if err2 != nil {
		log.Println("SystemLog表创建失败")
	}
}
