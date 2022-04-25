package route

import (
	"fmt"
	"localFileServer/rely"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mattn/go-colorable"
	"go.uber.org/zap"
)

func SetUp() *gin.Engine {
	// 处理cmd乱码
	gin.ForceConsoleColor()
	gin.DefaultWriter = colorable.NewColorableStdout()
	// 开启服务
	r := gin.Default()
	// 日志中间件
	r.Use(rely.GinLogger(), rely.GinRecovery(true))
	// 加载静态文件
	r.Static("static", "./static")
	r.StaticFile("favicon.ico", "./favicon.ico")
	r.POST("/upload", func(c *gin.Context) {
		// 处理跨域请求，和预请求
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST")

		file, err := c.FormFile("filename")
		if err != nil {
			zap.L().Error("upload file error", zap.Error(err))
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "请求失败",
				"data":    "err",
			})
			fmt.Println(err)
			return
		}
		dst := "./File/" + file.Filename
		c.SaveUploadedFile(file, dst)
		zap.L().Info("upload file success", zap.String("filename", file.Filename))
		c.JSON(http.StatusOK, gin.H{
			"code":    "success",
			"status":  "ok",
			"message": dst + "上传成功",
		})
	})
	r.LoadHTMLFiles("./index.html")
	r.GET("/", func(c *gin.Context) {
		zap.L().Info("index.html is loading----")
		c.HTML(http.StatusOK, "index.html", gin.H{
			"result": "SUCCESS",
		})
	})
	return r
}
