package routes

import (
	"dunkbing/web-scrap/db"
	"dunkbing/web-scrap/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"strconv"
)

func GetVideos(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
	}
	limit := 40

	videoColl := db.GetCollection(db.VideoColl)
	pageOptions := options.Find()
	pageOptions.SetSkip(int64(limit * (page - 1)))
	pageOptions.SetLimit(int64(limit))
	ctx := c.Request.Context()
	cur, err := videoColl.Find(ctx, bson.D{{}}, pageOptions)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
	}
	defer cur.Close(ctx)

	var videos []model.Video
	for cur.Next(ctx) {
		var video model.Video
		if err := cur.Decode(&video); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"error":   err.Error(),
			})
			return
		}
		videos = append(videos, video)
	}
	if err := cur.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    videos,
	})
}
