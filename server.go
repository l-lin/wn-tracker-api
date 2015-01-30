package main

import (
	"github.com/codegangsta/negroni"
	"github.com/l-lin/wn-tracker-api/novel"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"net/http"
	"os"
	"log"
	"strconv"
)

func main() {
	port := GetPort()
	db := os.Getenv("DATABASE_URL")
	log.Println(db)
	log.Println("Listening to..." + port)

	app := negroni.Classic()
	app.UseHandler(Router())
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

func Router() *mux.Router {
	r := render.New()
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/novels", func (w http.ResponseWriter, req *http.Request) {
		r.JSON(w, http.StatusOK, novel.FetchList())
	})
	muxRouter.HandleFunc("/novels/{id}", func (w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		n := novel.New()
		if vars["id"] != "" {
			n.Id, _ = strconv.ParseInt(vars["id"], 10, 64)
		}
		r.JSON(w, http.StatusOK, n)
	})
	return muxRouter
}
