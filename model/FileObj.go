package model

import (
	"PublicFileServer/gobalConfig"
	"PublicFileServer/util"
	"os"
	"time"
)

type FileObj struct {
	Id           int    `json:"id" gorm:"autoIncrement"`
	FileName     string `json:"fileName"`
	PathName     string `json:"pathName"`
	FileSize     int64  `json:"fileSize"`
	FileMd5      string `json:"fileMd5"`
	UploadDate   string `json:"uploadDate"`
	ShareCode    string `json:"shareCode"`
	FileLocation string `json:"fileLocation"`
}

func FileExist(fileName, fileMd5 string) (bool, string) {
	var fileObj FileObj
	util.DB.Where("file_md5=? and file_name=?", fileMd5, fileName).First(&fileObj)
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

func CreateFile(fileName, fileCode, fileMd5, pathName string, fileSize int64) {
	fileObj := FileObj{
		FileName:     fileName,
		PathName:     pathName,
		FileSize:     fileSize,
		FileMd5:      fileMd5,
		UploadDate:   time.Now().Format("2006-01-02 15:04:05"),
		ShareCode:    fileCode,
		FileLocation: "file/" + pathName,
	}
	util.DB.Create(&fileObj)
}

func GetFile(fileCode string) (FileObj, error) {
	var fileObj FileObj
	err := util.DB.Where("share_code=?", fileCode).First(&fileObj).Error
	return fileObj, err
}

func DelFile() {
	var files []FileObj

	util.DB.Where("upload_date < ?", time.Now().Add(-time.Hour*time.Duration(gobalConfig.FileLife)).Format("2006-01-02 15:04:05")).Find(&files)
	for _, file := range files {
		AddSystemLog("删除了文件："+file.PathName, "deleteFile")
		util.DB.Delete(&file)
		os.Remove("files/" + file.PathName)
	}
}
