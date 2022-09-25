package router

import (
	"PublicFileServer/model"
	"PublicFileServer/util"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/url"
	"os"
)

func upload(c *gin.Context) {
	file, _ := c.FormFile("file")
	err := c.SaveUploadedFile(file, "./files/"+file.Filename)
	if err != nil {
		c.JSON(200, model.Result{
			Code:    -1,
			Success: false,
			Message: fmt.Sprintf("%s 保存失败!", file.Filename),
		})
		return
	}
	pFile, err := os.Open("./files/" + file.Filename)
	if err != nil {
		c.JSON(200, model.Result{
			Code:    0,
			Success: false,
			Message: fmt.Sprintf("打开文件失败：%s", err),
		})
		return
	}
	defer pFile.Close()
	md5h := md5.New()
	io.Copy(md5h, pFile)
	fileMd5 := hex.EncodeToString(md5h.Sum(nil))
	fileExist, shareCode := model.FileExist(file.Filename, fileMd5)
	if fileExist {
		c.JSON(200, model.Result{
			Code:    1,
			Success: true,
			Message: "文件：" + file.Filename + " 已存在，提取码：" + shareCode,
		})
		return
	}
	success, code := GenerateCode()
	if !success {
		c.JSON(200, model.Result{
			Code:    1,
			Success: true,
			Message: "提取码生成失败,请重试！",
		})
		return
	}
	model.CreateFile(file.Filename, code, fileMd5, file.Size)
	model.AddSystemLog("...上传了文件："+file.Filename, "upload")
	c.JSON(200, model.Result{
		Code:    1,
		Success: true,
		Message: "文件：" + file.Filename + " 上传成功!提取码：" + code,
	})
}

func download(c *gin.Context) {
	code := c.Query("code")
	file, err := model.GetFile(code)
	if err != nil {
		return
	} else {
		model.AddSystemLog("...下载了文件："+file.FileName, "download")
		c.Header("Content-Type", "application/octet-stream")
		c.Header("Content-Disposition", "attachment; filename="+file.FileName)
		c.Header("filename", url.QueryEscape(file.FileName))
		c.Header("Content-Transfer-Encoding", "binary")
		c.File("./files/" + file.FileName)
	}
}

func exist(c *gin.Context) {
	log.Println("remote IP : ", c.RemoteIP())
	code := c.Query("code")
	file, err := model.GetFile(code)
	if err != nil {
		c.JSON(200, model.Result{
			Code:    1,
			Success: false,
		})
	} else {
		c.JSON(200, model.Result{
			Code:    1,
			Success: true,
			FileObj: file,
		})
	}
}

func File(e *gin.Engine) {
	g := e.Group("/file")
	{
		g.POST("/upload", upload)
		g.GET("/exist", exist)
		g.GET("/download", download)
	}
}

func GenerateCode() (bool, string) {
	time := 0
	for {
		if time > 100 {
			return false, ""
		}
		code := util.RandAllString(6)
		if model.CodeExist(code) {
			time++
			continue
		} else {
			return true, code
		}
	}
}
