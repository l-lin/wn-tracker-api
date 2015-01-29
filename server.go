package main

import (
	"github.com/codegangsta/negroni"
	"github.com/l-lin/wn-tracker-api/novels"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"net/http"
	"os"
	"log"
	"strconv"
)

func main() {
	port := GetPort()
	log.Println("Listening to..." + port)

	app := negroni.Classic()
	app.UseHandler(NewRouter())
	app.Run(port)
}

func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "4747"
		log.Println("[INFO] No PORT environment variable detected. Setting to ", port)
	}
	return ":" + port
}

func NewRouter() *mux.Router {
	r := render.New()
	router := mux.NewRouter()
	router.HandleFunc("/novels/{id}", func (w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		novel := novels.New()
		if vars["id"] != "" {
			novel.Id, _ = strconv.ParseInt(vars["id"], 10, 64)
		}
		r.JSON(w, http.StatusOK, novel)
	})
	return router
}
