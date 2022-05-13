package main

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"tools/configs"
	"tools/strs"
)

func main() {
	configs.InitConfig()
	if !configs.GetIsDebug() {
		gin.SetMode(gin.ReleaseMode)
		gin.DefaultWriter = ioutil.Discard
	}
	ser := gin.Default()
	ser.LoadHTMLGlob("templates/*")
	ser.Use(Cors())
	ser.SetTrustedProxies([]string{"127.0.0.1"})
	strsGroup := ser.Group("strs")
	strsGroup.GET("/", strs.StrToLowerIndex)
	strsGroup.GET("/index", strs.StrToLowerIndex)
	strsGroup.POST("/strtolower", strs.StrToLower)
	ser.Run(":" + configs.GetPort())
}

// 处理跨域请求,支持options访问
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		c.Header("Access-Control-Allow-Origin", origin)
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token,X-Token,X-User-Id")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS,DELETE,PUT")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		// 放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}
