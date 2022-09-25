package model

import (
	"PublicFileServer/util"
	"time"
)

type FileObj struct {
	Id           int    `json:"id" gorm:"autoIncrement"`
	FileName     string `json:"fileName"`
	FileSize     int64  `json:"fileSize"`
	FileMd5      string `json:"fileMd5"`
	UploadDate   string `json:"uploadDate"`
	ShareCode    string `json:"shareCode"`
	FileLocation string `json:"fileLocation"`
}

func FileExist(fileName string) (bool, string) {
	var fileObj FileObj
	util.DB.Where("file_md5=?", fileName).First(&fileObj)
	if fileObj.FileSize > 0 {
		return true, fileObj.ShareCode
	} else {
		return false, ""
	}
}

func CodeExist(code string) bool {
	var fileObj FileObj
	return util.DB.Where("share_code=?", code).First(&fileObj) == nil
}

func CreateFile(fileName, fileCode, fileMd5 string, fileSize int64) {
	fileObj := FileObj{
		FileName:     fileName,
		FileSize:     fileSize,
		FileMd5:      fileMd5,
		UploadDate:   time.Now().String(),
		ShareCode:    fileCode,
		FileLocation: "file/" + fileName,
	}
	util.DB.Create(&fileObj)
}
func GetFile(fileCode string) (FileObj, error) {
	var fileObj FileObj
	err := util.DB.Where("share_code=?", fileCode).First(&fileObj).Error
	return fileObj, err
}
