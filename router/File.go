package router

import (
	"PublicFileServer/gobalConfig"
	"PublicFileServer/model"
	"PublicFileServer/util"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/url"
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
			Message: "文件已存在，提取码：" + shareCode,
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
		Message: "文件上传成功!提取码：" + code,
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
	log.Println(gobalConfig.FileLife, gobalConfig.FileSize)
	c.JSON(200, gin.H{
		"fileLife": gobalConfig.FileLife,
		"fileSize": gobalConfig.FileSize,
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
		codeLength := gobalConfig.ShareCodeLength
		var code string
		switch gobalConfig.ShareCodeType {
		case 1:
			code = util.RandStringAll(codeLength)
		case 2:
			code = util.RandStringLarge(codeLength)
		case 3:
			code = util.RandStringSmall(codeLength)
		case 4:
			code = util.RandStringNum(codeLength)
		case 5:
			code = util.RandStringLargeSmall(codeLength)
		case 6:
			code = util.RandStringLargeNum(codeLength)
		case 7:
			code = util.RandStringSmallNum(codeLength)
		default:
			code = util.RandStringAll(codeLength)
		}

		if model.CodeExist(code) {
			t++
			continue
		} else {
			return true, code
		}
	}
}
