package db

import (
	"context"
	"dunkbing/web-scrap/model"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

func SeedVideos() {
	videoColl := GetCollection(VideoColl)
	index := 1
	for index <= 4000 {
		logrus.Info("Seeding ", index)
		ctx := context.Background()
		var vid model.Video
		err := videoColl.FindOne(ctx, bson.M{"index": index}).Decode(&vid)
		if err == mongo.ErrNoDocuments {
			vid = model.Video{
				Id:        "",
				Title:     "",
				Thumbnail: "",
				Duration:  "",
				Index:     index,
			}
			_, err := videoColl.InsertOne(ctx, vid)
			if err != nil {
				logrus.Error("Insert video error: ", err.Error())
			}
		}
		index++
		time.Sleep(time.Second / 2)
	}
}
