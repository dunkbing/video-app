package model

import (
    "fmt"
    "github.com/gocolly/colly"
    "github.com/robfig/cron/v3"
    "github.com/sirupsen/logrus"
    "strings"
    "time"
)

const totalPage = 10

type crawler struct {
    collector *colly.Collector
    page      int
}

func NewCrawler() *crawler {
    return &crawler{
        collector: colly.NewCollector(
            colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.5112.102 Safari/537.36 Edg/104.0.1293.70"),
        ),
        page: 1,
    }
}

func (c *crawler) Crawl() {
    baseUrl := "https://spankbang.com"
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
    c.collector.OnHTML(vidListSel, func(vidListEl *colly.HTMLElement) {
        vidItemSel := ".video-item"
        vidListEl.ForEach(vidItemSel, func(index int, vidItemEl *colly.HTMLElement) {
            video := Video{
                Id:        "--",
                Title:     "--",
                Thumbnail: "--",
                Duration:  "--",
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
            logrus.Info("Crawled video: ", video)
            time.Sleep(time.Second * 4)
        })
    })
    c.collector.OnScraped(func(r *colly.Response) {
        fmt.Println("Finished", r.Request.URL)
    })
    for c.page <= totalPage {
        url := fmt.Sprintf("%s/trending_videos/%v", baseUrl, c.page)
        err := c.collector.Visit(url)
        if err != nil {
            _ = fmt.Errorf("%s", err.Error())
        }
        c.collector.Wait()
        c.page++
        fmt.Println("Done...!")
    }
}

func (c *crawler) Start() {
    c.Crawl()
    job := cron.New()
    _, _ = job.AddFunc("* * * * *", func() {
        fmt.Println("Job is running")
    })
    crawlInterval := "*/30 * * * *"
    _, _ = job.AddFunc(crawlInterval, func() {
        c.Crawl()
    })
    job.Start()
}
