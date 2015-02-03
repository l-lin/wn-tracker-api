package feed

import (
	"testing"
)

func Test_FindRssFeedUrl(t *testing.T) {
	testWithUrl(t, "http://skythewood.blogspot.sg/")
	testWithUrl(t, "https://defiring.wordpress.com")
//	testWithUrl(t, "http://www.wuxiaworld.com/cdindex-html")
	url := "http://google.com"
	c := make(chan string)
	go FindRssFeedUrl(url, c)
	rssFeedUrl := <-c
	if rssFeedUrl != "" {
		t.Fatalf("[x] There should be no RSS feed for %s", url)
	}
}

func Test_ParseRssFeed(t *testing.T) {
	url := "http://www.wuxiaworld.com/feed/"
	c := make(chan *RSS)
	go ParseRssFeed(url, c)
	rss := <- c
	if rss == nil {
		t.Fatalf("[x] The RSS feed should not be nil!")
	}
}

func testWithUrl(t *testing.T, url string) {
	c := make(chan string)
	go FindRssFeedUrl(url, c)
	rssFeedUrl := <-c
	if rssFeedUrl == "" {
		t.Fatalf("[x] The rss feed of %s should not be empty!", url)
	}
}
