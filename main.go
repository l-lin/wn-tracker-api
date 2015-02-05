package main

import (
	"github.com/codegangsta/negroni"
	"github.com/l-lin/wn-tracker-api/web"
	"github.com/l-lin/wn-tracker-api/feed"
	oauth2 "github.com/goincremental/negroni-oauth2"
	sessions "github.com/goincremental/negroni-sessions"
	"github.com/goincremental/negroni-sessions/cookiestore"
	"os"
	"log"
)

func main() {
	secure := negroni.New()
	secure.Use(oauth2.LoginRequired())
	secure.UseHandler(web.NewSignInRouter())

	// Only redirect to Google login page for path /signin
	router := web.NewRouter()
	router.Handle("/signin", secure)

	// Start the cron to fetch the notifications from the RSS feeds
	feed.NewCronRss().Start()

	// Start the app
	app := negroni.Classic()
	app.Use(sessions.Sessions("wn_tracker", cookiestore.New([]byte(os.Getenv("SESSION_SECRET")))))
	app.Use(web.NewOAuth())
	app.UseHandler(router)
	app.Run(port())
}

func port() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "4747"
		log.Println("[-] No PORT environment variable detected. Setting to ", port)
	}
	return ":" + port
}
