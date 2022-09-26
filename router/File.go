package router

import (
	"PublicFileServer/model"
	"PublicFileServer/util"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"io"
	"net/url"
	"strconv"
	"time"
)

func upload(c *gin.Context) {
	file, _ := c.FormFile("file")
	pFile, err := file.Open()
	if err != nil {
		c.JSON(200, model.Result{
			Status:  "error",
			Success: false,
			Message: fmt.Sprintf("校验文件失败：%s", err),
		})
		return
	}
	defer pFile.Close()
	md5h := md5.New()
	io.Copy(md5h, pFile)
	fileMd5 := hex.EncodeToString(md5h.Sum(nil))
	pathName := time.Now().Format("20060102150405") + "_" + file.Filename
	fileExist, shareCode := model.FileExist(file.Filename, fileMd5)
	if fileExist {
		c.JSON(200, model.Result{
			Status:  "info",
			Success: true,
			Message: "文件：" + file.Filename + " 已存在，提取码：" + shareCode,
		})
		return
	} else {
		err2 := c.SaveUploadedFile(file, "./files/"+pathName)
		if err2 != nil {
			c.JSON(200, model.Result{
				Status:  "error",
				Success: false,
				Message: fmt.Sprintf("%s 保存失败!", file.Filename),
			})
			return
		}
	}
	success, code := GenerateCode()
	if !success {
		c.JSON(200, model.Result{
			Status:  "error",
			Success: true,
			Message: "提取码生成失败,请重试！",
		})
		return
	}
	model.CreateFile(file.Filename, code, fileMd5, pathName, file.Size)
	model.AddSystemLog("...上传了文件："+file.Filename, "upload")
	c.JSON(200, model.Result{
		Status:  "success",
		Success: true,
		Message: "文件：" + file.Filename + " 上传成功!提取码：" + code,
	})
}

func download(c *gin.Context) {
	code := c.Param("code")
	file, err := model.GetFile(code)
	if err != nil {
		return
	} else {
		model.AddSystemLog("...下载了文件："+file.FileName, "download")
		c.Header("Content-Type", "application/octet-stream")
		c.Header("Content-Disposition", "attachment; filename="+file.FileName)
		c.Header("filename", url.QueryEscape(file.FileName))
		c.Header("Content-Transfer-Encoding", "binary")
		c.File("./files/" + file.PathName)
	}
}

func exist(c *gin.Context) {
	code := c.Param("code")
	file, err := model.GetFile(code)
	if err != nil {
		c.JSON(200, model.Result{
			Status:  "error",
			Success: false,
		})
	} else {
		c.JSON(200, model.Result{
			Status:  "success",
			Success: true,
			FileObj: file,
		})
	}
}
func config(c *gin.Context) {
	fileLife, _ := strconv.Atoi(viper.GetString("config.fileLife"))
	fileSize, _ := strconv.Atoi(viper.GetString("config.fileSize"))
	c.JSON(200, gin.H{
		"fileLife": fileLife,
		"fileSize": fileSize,
	})
}
func File(e *gin.Engine) {
	g := e.Group("/file")
	{
		g.POST("/upload", upload)
		g.GET("/exist/:code", exist)
		g.GET("/download/:code", download)
		g.GET("/config", config)
	}
}

func GenerateCode() (bool, string) {
	t := 0
	for {
		if t > 100 {
			return false, ""
		}
		code := util.RandAllString(6)
		if model.CodeExist(code) {
			t++
			continue
		} else {
			return true, code
		}
	}
}
