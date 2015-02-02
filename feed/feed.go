package feed

import (
	"github.com/l-lin/wn-tracker-api/db"
	_ "github.com/lib/pq"
	"log"
	"time"
)

// The feed
type Feed struct {
	FeedId        string        `json:"feedId"`
	FeedUrl       string        `json:"feedUrl"`
	LastUpdated   time.Time     `json:"-"`
}

// Instanciate a new feed
func New() *Feed {
	return &Feed{}
}

// Fetch all feeds from the db
func GetList() []*Feed {
	feeds := make([]*Feed, 0)
	database := db.Connect()
	defer database.Close()

	rows, err := database.Query("SELECT f.feed_id, f.feed_url, f.last_updated FROM feeds f")
	if err != nil {
		log.Fatalf("[x] Error when getting the list of feeds. Reason: %s", err.Error())
	}
	for rows.Next() {
		feeds = append(feeds, toFeed(rows))
	}
	if err := rows.Err(); err != nil {
		log.Fatalf("[x] Error when getting the list of feeds. Reason: %s", err.Error())
	}
	return feeds
}

// Fetch the content of the rows and build a new default feed
func toFeed(rows db.RowMapper) *Feed {
	feed := New()
	err := rows.Scan(
		&feed.FeedId,
		&feed.FeedUrl,
		&feed.LastUpdated,
	)
	if err != nil {
		log.Printf("[-] Could not scan the feed. Reason: %s", err.Error())
	}
	return feed
}
