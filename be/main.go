package main

import (
	"dunkbing/web-scrap/configs"
	"dunkbing/web-scrap/crawler"
	"dunkbing/web-scrap/db"
	"dunkbing/web-scrap/routes"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	configs.InitLog()
	configs.LoadEnv()
	db.ConnectDB()
	_crawler := crawler.NewCrawler()
	fmt.Println(_crawler)
	//go _crawler.Start()
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/videos", routes.GetVideos)
	err := r.Run()
	if err != nil {
		fmt.Println("Error when running server: ", err.Error())
	}
}
