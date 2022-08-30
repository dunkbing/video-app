package main

import (
	"context"
	"dunkbing/web-scrap/configs"
	"dunkbing/web-scrap/crawler"
	"dunkbing/web-scrap/db"
	"dunkbing/web-scrap/routes"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

func main() {
	configs.InitLog()
	configs.LoadEnv()
	cfg := configs.GetConfig()
	db.ConnectDB()
	_, _ = db.GetCollection(db.VideoColl).
		Indexes().
		CreateOne(
			context.Background(),
			mongo.IndexModel{
				Keys: bson.D{
					{"title", "text"},
					{"names", "text"},
					{"tags", "text"},
				},
			},
		)
	_crawler := crawler.NewCrawler()
	fmt.Println(_crawler)
	go _crawler.Start()
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		origin := cfg.CorsOrigin
		c.Header("Access-Control-Allow-Origin", origin)
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})
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
