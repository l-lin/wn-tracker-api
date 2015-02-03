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

func SaveNovel(w http.ResponseWriter, r *http.Request)  {
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

func SignIn(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You are now authenticated! You can close this tab.")
}

func Notifications(w http.ResponseWriter, r *http.Request) {
	userC := make(chan string)
	go GetUserId(r, userC)
	userId := <- userC
	write(w, http.StatusOK, notification.GetList(userId))
}

func write(w http.ResponseWriter, status int, n interface {}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)

	if n != nil {
		if err := json.NewEncoder(w).Encode(n); err != nil {
			panic(err)
		}
	}
}
