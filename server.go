package main

import (
	"github.com/codegangsta/negroni"
	"net/http"
	"fmt"
	"os"
	"log"
)

func main() {
	port := GetPort()
	log.Println("Listening..." + port)

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
			fmt.Fprintf(w, "Welcome to the home page!")
		})

	n := negroni.Classic()
	n.UseHandler(mux)
	n.Run(port)
}

func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "4747"
		log.Println("[INFO] No PORT environment variable detected. Setting to ", port)
	}
	return ":" + port
}
