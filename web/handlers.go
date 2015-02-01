package web

import (
	"github.com/l-lin/wn-tracker-api/novel"
	"github.com/gorilla/mux"
	oauth2 "github.com/goincremental/negroni-oauth2"
	"net/http"
	"encoding/json"
	"log"
	"fmt"
	"io"
	"io/ioutil"
)

func Novels(w http.ResponseWriter, r *http.Request) {
	token := getToken(r)
	if !novel.Exists(token) {
		log.Printf("[-] No novels found for user %s. Copy the default one...", token)
		novel.CopyDefaultFor(token)
	}

	write(w, http.StatusOK, novel.GetList(token))
}

func Novel(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	n := novel.Get(id, getToken(r))
	if n != nil && n.Id != "" {
		log.Printf("[-] Found the novel id %s", id)
		write(w, http.StatusOK, n)
		return
	}

	// If we didn't find it, 404
	log.Printf("[-] Could not find the novel id %s", id)
	write(w, http.StatusNotFound, JsonErr{Code: http.StatusNotFound, Text: fmt.Sprintf("Novel not Found for id %s", id)})
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
	id := vars["id"]

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
	n.Id = id
	if !n.IsValid() || n.Id == "" {
		write(w, http.StatusPreconditionFailed, JsonErr{
			Code: http.StatusPreconditionFailed, Text: "The given novel has incorrect attributes",
		})
		return
	}
	log.Printf("[-] Updating novel id %s", id)
	n.Update()
	write(w, http.StatusOK, n)
}

func DeleteNovel(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	id := vars["id"]
	n := novel.New()
	n.Id = id
	log.Printf("[-] Deleting novel id %s", id)
	n.Delete()
	write(w, http.StatusNoContent, nil)
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

func getToken(r *http.Request) string {
	token := oauth2.GetToken(r)
	if token == nil || !token.Valid() {
		log.Fatal("[x] The user is not authenticated yet!")
	}
	return token.Access()
}
