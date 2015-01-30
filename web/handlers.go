package web

import (
	"github.com/l-lin/wn-tracker-api/novel"
	"github.com/gorilla/mux"
	"net/http"
	"encoding/json"
	"log"
	"strconv"
	"fmt"
	"io"
	"io/ioutil"
)

func Novels(w http.ResponseWriter, r *http.Request) {
	write(w, http.StatusOK, novel.GetList())
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
	n.Save()
	write(w, http.StatusCreated, n)
}

func Novel(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	n := novel.Get(id)
	if n != nil && n.Id > 0 {
		log.Printf("[+] Found the novel id %d", id)
		write(w, http.StatusOK, n)
		return
	}

	// If we didn't find it, 404
	log.Printf("[-] Could not find the novel id %d", id)
	write(w, http.StatusNotFound, JsonErr{Code: http.StatusNotFound, Text: fmt.Sprintf("Novel not Found for id %d", id)})
}

func write(w http.ResponseWriter, status int, n interface {}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(n); err != nil {
		panic(err)
	}
}
