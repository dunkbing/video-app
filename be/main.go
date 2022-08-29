package main

import (
    "dunkbing/web-scrap/model"
    "fmt"
    "github.com/gin-gonic/gin"
    "net/http"
)

func main() {
    model.InitLog()
    crawler := model.NewCrawler()
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
