package web

import (
	"github.com/codegangsta/negroni"
	oauth2 "github.com/goincremental/negroni-oauth2"
	"os"
)

func NewOAuth() negroni.Handler {
	return oauth2.Google(&oauth2.Config{
		ClientID: 		os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret:	os.Getenv("GOOGLE_CLIENT_SECRET"),
		RedirectURL:  	os.Getenv("GOOGLE_REDIRECT_URI"),
		Scopes:       	[]string{"https://www.googleapis.com/auth/plus.login"},
	})
}
