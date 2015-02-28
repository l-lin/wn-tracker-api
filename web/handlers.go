package web

import (
	"github.com/l-lin/wn-tracker-api/novel"
	"github.com/l-lin/wn-tracker-api/notification"
	"github.com/l-lin/wn-tracker-api/feed"
	"github.com/gorilla/mux"
	"net/http"
	"encoding/json"
	"log"
	"fmt"
	"io"
	"io/ioutil"
)

// Handler to fetch the list of novels
func Novels(w http.ResponseWriter, r *http.Request) {
	userC := make(chan string)
	go GetUserId(r, userC)
	userId := <- userC
	if !novel.Exists(userId) {
		log.Printf("[-] No novels found for user %s. Copy the default one...", userId)
		novel.CopyDefaultFor(userId)
	}

	write(w, http.StatusOK, novel.GetList(userId))
}

// Handler to fetch a novel
func Novel(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	novelId := vars["novelId"]

	userC := make(chan string)
	go GetUserId(r, userC)
	userId := <- userC

	n := novel.Get(novelId, userId)
	if n != nil && n.NovelId != "" {
		log.Printf("[-] Found the novel novelId %s", novelId)
		write(w, http.StatusOK, n)
		return
	}

	// If we didn't find it, 404
	log.Printf("[-] Could not find the novel novelId %s", novelId)
	write(w, http.StatusNotFound, JsonErr{Code: http.StatusNotFound, Text: fmt.Sprintf("Novel not Found for novelId %s", novelId)})
}

// Handler to save a novel
func SaveNovel(w http.ResponseWriter, r *http.Request)  {
	var n novel.Novel
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		log.Printf("[x] Could not read the body. Reason: %s", err.Error())
		write(w, http.StatusInternalServerError, JsonErr{Code: http.StatusInternalServerError, Text: "Could not read the body."})
		return
	}
	if err := r.Body.Close(); err != nil {
		log.Printf("[x] Could not close ready the body. Reason: %s", err.Error())
		write(w, http.StatusInternalServerError, JsonErr{Code: http.StatusInternalServerError, Text: "Could not close the body."})
		return
	}
	if err := json.Unmarshal(body, &n); err != nil {
		// 422: unprocessable entity
		write(w, 422, JsonErr{Code: 422, Text: "Could not parse the given parameter"})
		return
	}

	userC := make(chan string)
	rssC := make(chan string)
	go GetUserId(r, userC)
	go feed.FindRssFeedUrl(n.Url, rssC)
	n.UserId = <- userC
	n.FeedUrl = <- rssC

	if !n.IsValid() {
		write(w, http.StatusPreconditionFailed, JsonErr{
			Code: http.StatusPreconditionFailed, Text: "The title should not be empty!",
		})
		return
	}

	log.Printf("[-] Creating new novel with title %s", n.Title)
	n.Save()
	write(w, http.StatusCreated, n)
}

// Handler to update a novel
func UpdateNovel(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	novelId := vars["novelId"]

	var n novel.Novel
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		log.Fatalf("[x] Could not read the body. Reason: %s", err.Error())
	}
	if err := r.Body.Close(); err != nil {
		log.Fatalf("[x] Could not close ready the body. Reason: %s", err.Error())
	}
	if err := json.Unmarshal(body, &n); err != nil {
		// 422: unprocessable entity
		write(w, 422, JsonErr{Code: 422, Text: "Could not parse the given parameter"})
		return
	}
	n.NovelId = novelId

	userC := make(chan string)
	rssC := make(chan string)
	go GetUserId(r, userC)
	go feed.FindRssFeedUrl(n.Url, rssC)
	n.UserId = <- userC
	n.FeedUrl = <- rssC

	if !n.IsValid() || n.NovelId == "" {
		write(w, http.StatusPreconditionFailed, JsonErr{
			Code: http.StatusPreconditionFailed, Text: "The given novel has incorrect attributes",
		})
		return
	}
	log.Printf("[-] Updating novel novelId %s", novelId)
	n.Update()
	write(w, http.StatusOK, n)
}

// Handler to delete a novel
func DeleteNovel(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	novelId := vars["novelId"]
	n := novel.New()
	n.NovelId = novelId

	userC := make(chan string)
	go GetUserId(r, userC)
	n.UserId = <- userC

	log.Printf("[-] Deleting novel novelId %s", novelId)
	n.Delete()
	write(w, http.StatusNoContent, nil)
}

// Handler to search novels
func SearchNovels(w http.ResponseWriter, r *http.Request) {
	userC := make(chan string)
	go GetUserId(r, userC)
	userId := <- userC
	if !novel.Exists(userId) {
		log.Printf("[-] No novels found for user %s. Copy the default one...", userId)
		novel.CopyDefaultFor(userId)
	}

	query := r.URL.Query()
	title := query["title"][0]

	write(w, http.StatusOK, novel.Search(userId, title))
}

// Handler to sign in Google account
func SignIn(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You are now authenticated! You can close this tab.")
}

// Handler to fetch the list of notifications
func Notifications(w http.ResponseWriter, r *http.Request) {
	userC := make(chan string)
	go GetUserId(r, userC)
	userId := <- userC
	write(w, http.StatusOK, notification.GetList(userId))
}

// Handler to fetch a notification
func Notification(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	notificationId := vars["notificationId"]

	n := notification.Get(notificationId)
	if n != nil && n.NotificationId != "" {
		log.Printf("[-] Found the notification id %s", notificationId)
		write(w, http.StatusOK, n)
		return
	}

	// If we didn't find it, 404
	log.Printf("[-] Could not find the notification id %s", notificationId)
	write(w, http.StatusNotFound, JsonErr{Code: http.StatusNotFound, Text: fmt.Sprintf("Notification not Found for notificationId %s", notificationId)})
}

// Handler to delete a notification
func DeleteNotification(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	notificationId := vars["notificationId"]
	n := notification.New()
	n.NotificationId = notificationId

	log.Printf("[-] Deleting notification id %s", notificationId)
	n.Delete()
	write(w, http.StatusNoContent, nil)
}

// This Handler is used only to check if the user is indeed authenticated
func AuthTest(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "You are now authenticated! You can close this tab.")
}

// Write the response in JSON Content-type
func write(w http.ResponseWriter, status int, n interface {}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)

	if n != nil {
		if err := json.NewEncoder(w).Encode(n); err != nil {
			panic(err)
		}
	}
}
