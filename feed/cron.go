package feed

import (
	"github.com/l-lin/wn-tracker-api/notification"
	"github.com/robfig/cron"
	"log"
	"time"
)

// Cron to fetch all the rss content
func NewCronRss() *cron.Cron {
	c := cron.New()
	c.AddFunc("0 */1 * * * *", fillNotifications)
	return c
}

func fillNotifications() {
	log.Printf("[-] CRON - Starting to fill the table notifications...")
	feeds := GetList()
	c := make(chan []notification.Notification)
	go getNotifications(feeds[:len(feeds)/2], c)
	go getNotifications(feeds[len(feeds)/2:], c)
	n1, n2 := <-c, <-c
	notifications := append(n1, n2...)
	if len(notifications) > 0 {
		count := 0
		for _, n := range notifications {
			if !notification.Exists(n.FeedId, n.Link) {
				n.Save()
				count++
			}
		}
		log.Printf("[-] CRON - Saved %d notifications", count)
	} else {
		log.Printf("[-] CRON - There are no notifications to save")
	}
	log.Printf("[-] CRON - Finished filling the table notifications...")
}

func getNotifications(feeds []*Feed, c chan []notification.Notification) {
	notifications := make([]notification.Notification, 0)
	for _, f := range feeds {
		if f.FeedUrl == "" {
			continue
		}

		// Fetch the content of the rss
		fc := make(chan *RSS)
		go ParseRssFeed(f.FeedUrl, fc)
		rss := <- fc

		// If last build date, then no need to see this feed
		lastBuildDate, _ := time.Parse(time.RFC1123Z, rss.Items.LastBuildDate)
		if lastBuildDate.Before(f.LastUpdated) {
			continue
		}

		hasNews := false
		for _, item := range rss.Items.ItemList {
			pubDate, err := time.Parse(time.RFC1123Z, item.PubDate)
			if err != nil {
				log.Printf("[x] Cannot read the pub date from feed URL %s", f.FeedUrl)
				break;
			}
			if pubDate.Before(f.LastUpdated) {
				break;
			}
			notifications = append(notifications, notification.Notification{
					FeedId: f.FeedId,
					Title: item.Title,
					Link: item.Link,
					PubDate: pubDate,
				})
			hasNews = true
		}
		if hasNews {
			f.LastUpdated = time.Now()
			defer f.Update()
		}
	}
	c <- notifications
}
