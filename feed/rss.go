package feed

import (
	"regexp"
	"log"
	"io/ioutil"
	"net/http"
	"encoding/xml"
	//	"time"
)

// The RSS representation
type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Items   Items `xml:"channel"`
}
type Items struct {
	XMLName        xml.Name `xml:"channel"`
	LastBuildDate  string `xml:"lastBuildDate"`
	Title          string `xml:"title"`
	ItemList       []Item `xml:"item"`
}
type Item struct {
	Title          string `xml:"title"`
	Link           string `xml:"link"`
	PubDate        string `xml:"pubDate"`
}

// Find the rss feed url from a given url
func FindRssFeedUrl(url string, c chan string) {
	log.Printf("[-] Searching a RSS feed for %s", url)

	result := ""
	// If the url is empty, then do not fetch it
	if url == "" {
		c <- result
		return
	}

	resp, err := http.Get(url)
	if err != nil {
		log.Printf("[x] Could not fetch content of %s. Reason: %s", url, err.Error())
		c <- result
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("[x] Error reading content of %s. Reason: %s", url, err.Error())
		c <- result
		return
	}

	// TODO: Improve the regexp so that http://www.wuxiaworld.com/cdindex-html does not return an incorrect result
	re := regexp.MustCompile(`<link rel="alternate" type=\"application/rss\+xml\" title=".+" href="(.+)" />`)
	if !re.MatchString(string(body)) {
		log.Printf("[-] No RSS feed found for %s", url)
	} else {
		link := re.FindString(string(body))
		re = regexp.MustCompile(`href="(.*)"`)
		result = re.FindStringSubmatch(link)[1]
		if len(result) > 256 {
			log.Printf("[x] The regex went wrong... Could not find the rss feed for %s. Got instead: %s", url, result)
			result = ""
		} else {
			log.Printf("[-] RSS feed found for %s: %s", url, result)
		}
	}
	c <- result
}

// Fetch the content of the url and unmarshal the body to a struct
func ParseRssFeed(url string, c chan *RSS) {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("[x] Could not fetch content of %s. Reason: %s", url, err.Error())
		c <- &RSS{}
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("[x] Error reading content of %s. Reason: %s", url, err.Error())
		c <- &RSS{}
		return
	}

	var rss RSS
	err = xml.Unmarshal(body, &rss)
	if err != nil {
		log.Printf("[x] Error when trying to unmarshal the rss feed %s. Reason: %s", url, err.Error())
	}
	c <- &rss
}
