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
	"net/http"
	"os"
)

func upload(c *gin.Context) {
	file, _ := c.FormFile("file")
	log.Println(file.Filename)
	err := c.SaveUploadedFile(file, "./files/"+file.Filename)
	if err != nil {
		c.String(http.StatusOK, fmt.Sprintf("'%s' upload fail!", file.Filename))
		return
	}
	pFile, err := os.Open("./files/" + file.Filename)
	if err != nil {
		c.String(http.StatusOK, fmt.Sprintf("打开文件失败：%s", err))
		return
	}
	defer pFile.Close()
	md5h := md5.New()
	io.Copy(md5h, pFile)
	fileMd5 := hex.EncodeToString(md5h.Sum(nil))
	exist, shareCode := model.FileExist(fileMd5)
	if exist {
		c.String(http.StatusOK, "文件已存在，提取码："+shareCode)
		return
	}
	success, code := GenerateCode()
	if !success {
		c.String(http.StatusOK, "提取码生成失败,请重试！")
		return
	}
	model.CreateFile(file.Filename, code, fileMd5, file.Size)
	c.String(http.StatusOK, fmt.Sprintf("'%s' 上传成功!提取码:%s", file.Filename, code))
}
func download(c *gin.Context) {

}

func File(e *gin.Engine) {
	g := e.Group("/file")
	{
		g.POST("/upload", upload)
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
