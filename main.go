package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "test",
		})
	})

	// 此规则能够匹配/user/john这种格式，但不能匹配/user/ 或 /user这种格式
	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "user name:%s", name)
	})

	// 但是，这个规则既能匹配/user/name/格式也能匹配/user/name/action这种格式
	// 如果没有其他路由器匹配/user/name，它将重定向到/user/name/
	r.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})

	r.Run()
}
