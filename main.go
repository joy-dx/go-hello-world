package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joy-dx/go-hello-world/utils"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello joydx",
		})
	})
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"secret": "AIzaSyDaGmWKa4JsXZ-HjGw7ISLn_3namBGewQe",
		})
	})
	r.GET("/a-new-route", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": utils.Capitalize("nice to meet you"),
		})
	})
	r.GET("/test-credential-leak", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"secret": "BIzaSyDaGmWKa4JsXZ-HjGw7ISLn_3namBGfuQe",
		})
	})
	return r
}

func main() {
	r := setupRouter()
	port, present := os.LookupEnv("GOHW_PORT")
	if !present {
		port = "8080"
	}
	if err := r.Run(fmt.Sprintf(":%s", port)); err != nil {
		log.Fatal(err)
	}
}
