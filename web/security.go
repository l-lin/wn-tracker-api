package web

import (
	"github.com/l-lin/wn-tracker-api/token"
	"github.com/codegangsta/negroni"
	oauth2 "github.com/goincremental/negroni-oauth2"
	sessions "github.com/goincremental/negroni-sessions"
	"os"
	"log"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"fmt"
)

const (
	SESSION_USER_ID = "user_id"
	googleUserInfoEndPoint = "https://www.googleapis.com/oauth2/v1/userinfo"
)

// The user info for Google account
type UserInfo struct {
	Id string
}

// Returns a new Negroni middleware using Google OAuth2
func NewOAuth() negroni.Handler {
	return oauth2.Google(&oauth2.Config{
	ClientID: os.Getenv("GOOGLE_CLIENT_ID"),
	ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
	RedirectURL: os.Getenv("GOOGLE_REDIRECT_URI"),
	Scopes: []string{"https://www.googleapis.com/auth/userinfo.profile"},
})
}

// Get the user ID from a given token.
// It will make a GET request to https://www.googleapis.com/oauth2/v1/userinfo?access_token=...
func GetUserId(r *http.Request, c chan string) {
	s := sessions.GetSession(r)
	sessionUserId := s.Get(SESSION_USER_ID)
	if sessionUserId != nil {
		userId := fmt.Sprintf("%v", sessionUserId)
		log.Printf("[-] Found the userId %s from the session", userId)
		c <- userId
		return
	}

	oauthT := oauth2.GetToken(r)
	if oauthT == nil || !oauthT.Valid() {
		log.Fatal("[x] The user is not authenticated yet!")
	}
	accessToken := oauthT.Access()

	log.Printf("[-] Getting the user id from access token %s", accessToken)
	endPoint := googleUserInfoEndPoint + "?access_token=" + accessToken
	resp, err := http.Get(endPoint)
	if err != nil {
		log.Printf("[x] Could not find the user info with token %s. Reason: %s", accessToken, err.Error())
		c <- ""
		return
	}
	defer resp.Body.Close()
	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("[x] Error reading content of %s. Reson: %s", endPoint, err.Error())
		c <- ""
		return
	}
	var userInfo UserInfo
	err = json.Unmarshal(response, &userInfo)
	if err != nil {
		log.Printf("[x] Could not unmarshal the user info. Reason: %s", err.Error())
		c <- ""
		return
	}

	// Setting the userId to the session
	s.Set(SESSION_USER_ID, userInfo.Id)
	saveTokenIfNotExists(oauthT, userInfo.Id)

	c <- userInfo.Id
}

// Save or update the token
func saveTokenIfNotExists(oauthT oauth2.Tokens, userId string) {
	if !token.Exists(userId) {
		t := token.New()
		t.UserId = userId
		t.RefreshToken = oauthT.Refresh()
		log.Printf("[-] Saving token %v", t)
		t.Save()
	} else if oauthT.Refresh() != "" {
		// If the refresh is not empty => the user had revoked the permissions => we have to update the token
		t := token.Get(userId)
		t.RefreshToken = oauthT.Refresh()
		log.Printf("[-] Updating the token %v", t)
		t.Update()
	}
}
