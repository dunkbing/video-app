package routes

import (
	"dunkbing/web-scrap/db"
	"dunkbing/web-scrap/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"math"
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
	query := c.DefaultQuery("query", "")

	videoColl := db.GetCollection(db.VideoColl)
	pageOptions := options.Find()
	pageOptions.SetSkip(int64(limit * (page - 1)))
	pageOptions.SetLimit(int64(limit))

	filterOptions := bson.D{{}}
	if query != "" {
		filterOptions = bson.D{{"$text", bson.D{{"$search", query}}}}
	}

	ctx := c.Request.Context()
	cur, err := videoColl.Find(ctx, filterOptions, pageOptions)
	count, err := videoColl.CountDocuments(ctx, filterOptions)
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
		"data": gin.H{
			"items":       videos,
			"totalPages":  math.Ceil(float64(count) / float64(limit)),
			"total":       count,
			"currentPage": page,
		},
	})
}
