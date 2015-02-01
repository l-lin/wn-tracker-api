package rss

import (
	"testing"
	"fmt"
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

func testWithUrl(t *testing.T, url string) {
	c := make(chan string)
	go FindRssFeedUrl(url, c)
	rssFeedUrl := <-c
	if rssFeedUrl == "" {
		t.Fatalf("[x] The rss feed of %s should not be empty!", url)
	}
	fmt.Printf("[-] RSS Feed URL: %s\n", rssFeedUrl)
}
