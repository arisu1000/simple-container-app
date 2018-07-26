package main

import "github.com/gin-gonic/gin"

func main() {
	appVersion := "v0.2"
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		resp := gin.H{
			"app version": appVersion,
			"header": c.Request.Header,
			"client ip": c.Request.RemoteAddr,
		}

		c.JSON(200, resp)
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
