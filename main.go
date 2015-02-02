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
	"net/http"
)

func main() {
	secure := negroni.Classic()
	secure.Use(oauth2.LoginRequired())
	secure.UseHandler(web.NewRouter())

	router := http.NewServeMux()
	router.Handle("/", secure)
	router.HandleFunc("/notification", web.Notification)

	feed.NewCronRss()

	app := negroni.New()
	app.Use(sessions.Sessions("my_session", cookiestore.New([]byte("secret123"))))
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
