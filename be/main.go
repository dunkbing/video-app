package main

import (
	"dunkbing/web-scrap/configs"
	"dunkbing/web-scrap/crawler"
	"dunkbing/web-scrap/db"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	configs.InitLog()
	configs.LoadEnv()
	db.ConnectDB()
	//db.SeedVideos()
	crawler := crawler.NewCrawler()
	go crawler.Start()
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	err := r.Run()
	if err != nil {
		fmt.Println("Error when running server: ", err.Error())
	}
}
