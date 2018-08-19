package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"os"
	"io"
)

const (
	appVersion = "v0.2"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		resp := gin.H{
			"app version": appVersion,
			"header":      c.Request.Header,
			"client ip":   c.Request.RemoteAddr,
		}
		c.JSON(200, resp)
	})
	r.GET("/env", func(c *gin.Context) {
		resp := gin.H{
			"env": os.Environ(),
		}
		c.JSON(200, resp)
	})

	r.GET("/volume-config", func(c *gin.Context) {
		filepath := c.Query("path")
		b, err := ioutil.ReadFile(filepath)
		if err != nil {
			fmt.Print(err)
		}
		configStr := string(b)

		resp := gin.H{
			"volume-content": configStr,
		}
		c.JSON(200, resp)
	})
	return r
}

func main() {
	// Logging to a file.
	f, _ := os.Create("/tmp/app.log")
	gin.DefaultWriter = io.MultiWriter(f)

	r := setupRouter()
	r.Run() // listen and serve on 0.0.0.0:8080
}
