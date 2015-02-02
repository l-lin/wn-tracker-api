package web

import (
	"github.com/codegangsta/negroni"
	oauth2 "github.com/goincremental/negroni-oauth2"
	"os"
	"log"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

const googleUserInfoEndPoint = "https://www.googleapis.com/oauth2/v1/userinfo"

type UserInfo struct {
	Id string
}

func NewOAuth() negroni.Handler {
	return oauth2.Google(&oauth2.Config{
		ClientID: 		os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret:	os.Getenv("GOOGLE_CLIENT_SECRET"),
		RedirectURL:  	os.Getenv("GOOGLE_REDIRECT_URI"),
		Scopes:       	[]string{"https://www.googleapis.com/auth/userinfo.profile"},
	})
}

func GetUserId(r *http.Request, c chan string) {
	token := oauth2.GetToken(r)
	if token == nil || !token.Valid() {
		log.Fatal("[x] The user is not authenticated yet!")
	}
	accessToken := token.Access()

	log.Printf("[-] Getting the user id from token %s", accessToken)
	endPoint := googleUserInfoEndPoint + "?access_token=" + accessToken
	resp, err := http.Get(endPoint)
	if err != nil {
		log.Fatalf("[x] Could not find the user info with token %s. Reason: %s", accessToken, err.Error())
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("[x] Error reading content of %s. Reson: %s", endPoint, err.Error())
	}
	var userInfo UserInfo
	err = json.Unmarshal(body, &userInfo)
	if err != nil {
		log.Fatalf("[x] Could not unmarshal the user info. Reason: %s", err.Error())
	}
	c <- userInfo.Id
}
