package crawler

import (
	"context"
	"dunkbing/web-scrap/db"
	"dunkbing/web-scrap/model"
	"fmt"
	"github.com/gocolly/colly"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"strconv"
	"strings"
	"time"
)

type crawler struct {
	collector  *colly.Collector
	page       int
	totalPage  int
	videoIndex int
}

func NewCrawler() *crawler {
	return &crawler{
		collector: colly.NewCollector(
			colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.5112.102 Safari/537.36 Edg/104.0.1293.70"),
		),
		page:       1,
		totalPage:  60,
		videoIndex: 1,
	}
}

func (c *crawler) Crawl() {
	logrus.Info("Start crawling....")
	start := time.Now()
	baseUrl := "https://spankbang.com"
	videoColl := db.GetCollection(db.VideoColl)
	c.collector.OnRequest(func(r *colly.Request) {
		r.Headers.Set("X-Requested-With", "XMLHttpRequest")
		r.Headers.Set("Referer", baseUrl)
		logrus.Info(fmt.Sprintf("Preparing request for page %v: %s", c.page, r.URL))
	})
	_ = c.collector.Limit(&colly.LimitRule{
		Delay: time.Second * 4,
	})
	c.collector.OnError(func(r *colly.Response, err error) {
		logrus.Error("Request URL:", r.Request.URL, "failed with response:", r, "Error:", err)
	})
	c.collector.OnResponse(func(r *colly.Response) {
		fmt.Printf("Received response %v\n", r.StatusCode)
	})

	vidListSel := "#browse_new > div > div > div.video-list.video-rotate.video-list-with-ads"
	lastPageSel := "#browse_new > div > div > div.pagination > ul > li:nth-child(6) > a"
	c.collector.OnHTML(lastPageSel, func(lastPageEl *colly.HTMLElement) {
		if c.page == 1 {
			totalPage, err := strconv.Atoi(lastPageEl.Text)
			logrus.Info("Total page: ", totalPage)
			if err == nil {
				c.totalPage = totalPage
			}
		}
	})

	c.collector.OnHTML(vidListSel, func(vidListEl *colly.HTMLElement) {
		vidItemSel := ".video-item"
		vidListEl.ForEach(vidItemSel, func(index int, vidItemEl *colly.HTMLElement) {
			video := model.Video{
				Id:        "--",
				Title:     "--",
				Thumbnail: "--",
				Duration:  "--",
				Index:     c.videoIndex,
			}
			duration := vidItemEl.ChildText("p > span.l")
			video.Duration = duration
			videoPage := c.collector.Clone()
			videoPage.OnRequest(func(r *colly.Request) {
				logrus.Info("Preparing request for video: ", r.URL)
			})
			videoPage.OnHTML(".left", func(e *colly.HTMLElement) {
				title := e.ChildText("h1")
				thumbnail := e.ChildAttr(".play_cover > img", "src")
				video.Title = title
				video.Thumbnail = thumbnail
			})
			starSel := "#video > div.left > div.info > section.details > div > div:nth-child(4)"
			videoPage.OnHTML(starSel, func(starEl *colly.HTMLElement) {
				starText := starEl.ChildText("span")
				logrus.Info("Span starText: ", starText)
				if starText == "Pornstar:" {
					var names []string
					starEl.ForEach("div a", func(i int, aEl *colly.HTMLElement) {
						names = append(names, aEl.Text)
					})
					video.Names = strings.Join(names, "|")

				} else if starText == "Tags:" {
					var tags []string
					starEl.ForEach("div a", func(i int, aEl *colly.HTMLElement) {
						tags = append(tags, aEl.Text)
					})
					video.Tags = strings.Join(tags, "|")
				}
			})
			tagSel := "#video > div.left > div.info > section.details > div > div:nth-child(5)"
			videoPage.OnHTML(tagSel, func(tagEl *colly.HTMLElement) {
				if tagEl == nil {
					return
				}
				var tags []string
				tagEl.ForEach("div a", func(i int, aEl *colly.HTMLElement) {
					tags = append(tags, aEl.Text)
				})
				video.Tags = strings.Join(tags, "|")
			})
			videoLink := vidItemEl.ChildAttr("a.thumb", "href")
			linkChunks := strings.Split(videoLink, "/")
			if len(linkChunks) > 1 {
				video.Id = linkChunks[1]
			}
			videoLink = fmt.Sprintf("https://spankbang.com%s", videoLink)
			err := videoPage.Visit(videoLink)
			if err != nil {
				logrus.Error("Visiting %s error: %s\n", videoLink, err.Error())
			}
			videoPage.Wait()
			if video.Id != "--" {
				logrus.Info("Crawled video: ", video)

				var existedVid model.Video
				ctx := context.Background()
				err = videoColl.FindOne(ctx, bson.M{"index": index}).Decode(&existedVid)
				if err == mongo.ErrNoDocuments {
					existedVid = model.Video{
						Id:        "",
						Title:     "",
						Thumbnail: "",
						Duration:  "",
						Index:     c.videoIndex,
					}
					_, err := videoColl.InsertOne(ctx, existedVid)
					if err != nil {
						logrus.Error("Insert video error: ", err.Error())
					}
				}
				update := bson.M{
					"id":        video.Id,
					"title":     video.Title,
					"thumbnail": video.Thumbnail,
					"duration":  video.Duration,
					"names":     video.Names,
					"tags":      video.Tags,
				}
				_, err := videoColl.UpdateOne(context.Background(), bson.M{"index": c.videoIndex}, bson.M{"$set": update})
				if err != nil {
					logrus.Error("Error updating video ", c.videoIndex, video.Id)
				} else {
					c.videoIndex++
				}
			}
			time.Sleep(time.Second * 8)
		})
	})
	c.collector.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
	})
	for c.page <= c.totalPage {
		logrus.Info("Crawling page: ", c.page)
		url := fmt.Sprintf("%s/trending_videos/%v", baseUrl, c.page)
		err := c.collector.Visit(url)
		if err != nil {
			_ = fmt.Errorf("%s", err.Error())
		}
		c.collector.Wait()
		c.page++
		logrus.Info("Crawled page: ", c.page)
	}
	elapsed := time.Since(start)
	logrus.Info("Done crawling. Took ", elapsed.Minutes())
}

func (c *crawler) Start() {
	c.Crawl()
	job := cron.New()
	crawlInterval := "0 */4 * * *"
	_, _ = job.AddFunc(crawlInterval, func() {
		c.Crawl()
	})
	job.Entries()
	job.Start()
}
