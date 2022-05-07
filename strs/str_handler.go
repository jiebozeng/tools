package strs

import (
	"github.com/gin-gonic/gin"
	"github.com/jiebozeng/golangutils/convert"
	"net/http"
	"strings"
)

func StrToLower(c *gin.Context) {
	inStr := c.PostForm("inStr")
	outStr := strings.ToLower(convert.ToString(inStr))
	c.HTML(http.StatusOK,"strToLower.html",gin.H{
		"inStr":inStr,
		"outStr":outStr,
	})
}

func StrToLowerIndex(c *gin.Context)  {
	c.HTML(http.StatusOK,"strToLower.html",gin.H{
	})
}