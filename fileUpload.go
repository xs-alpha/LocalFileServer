package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mattn/go-colorable"
)

func main() {
	gin.ForceConsoleColor()
	gin.DefaultWriter = colorable.NewColorableStdout()
	r := gin.Default()
	r.Static("static", "./static")
	r.StaticFile("favicon.ico", "./favicon.ico")
	r.POST("/upload", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST")

		file, err := c.FormFile("filename")
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"message": "请求失败",
				"data":    "err",
			})
			fmt.Println(err)
			return
		}
		dst := "./File/" + file.Filename
		c.SaveUploadedFile(file, dst)

		c.JSON(http.StatusOK, gin.H{
			"code":    "success",
			"status":  "ok",
			"message": dst + "上传成功",
		})
	})
	r.LoadHTMLFiles("./index.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"urluploadlocal": "http://192.168.3.5:8080/upload",
			"result":         "SUCCESS",
		})
	})

	r.Run(":8080")
}
