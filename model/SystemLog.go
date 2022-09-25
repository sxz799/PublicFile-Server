package model

import (
	"PublicFileServer/util"
	"time"
)

type SystemLog struct {
	Id   int    `json:"id" gorm:"autoIncrement"`
	Type string `json:"type"`
	Msg  string `json:"msg"`
	Date string `json:"date"`
}

func AddSystemLog(msg, Type string) {
	sLog := SystemLog{
		Msg:  msg,
		Type: Type,
		Date: time.Now().Format("2006-01-02 15:04:05"),
	}
	util.DB.Create(&sLog)
}
