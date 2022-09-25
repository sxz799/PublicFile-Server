package model

import (
	"PublicFileServer/util"
	"log"
)

type Result struct {
	Code    int     `json:"code"`
	Success bool    `json:"success"`
	Message string  `json:"message"`
	FileObj FileObj `json:"fileObj"`
}

func InitAutoMigrateDB() {
	err := util.DB.Set("gorm:table_options", "DEFAULT CHARSET=utf8").AutoMigrate(FileObj{})
	if err != nil {
		log.Println("FileObj表创建失败")
	}
	err2 := util.DB.Set("gorm:table_options", "DEFAULT CHARSET=utf8").AutoMigrate(SystemLog{})
	if err2 != nil {
		log.Println("SystemLog表创建失败")
	}
}
