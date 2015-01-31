package main

import (
	"github.com/codegangsta/negroni"
	"github.com/l-lin/wn-tracker-api/web"
	"os"
	"log"
)

func main() {
	port := GetPort()
	log.Println("[-] Listening on...", port)

	app := negroni.Classic()
	app.UseHandler(web.NewRouter())
	app.Run(port)
}

func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "4747"
		log.Println("[-] No PORT environment variable detected. Setting to ", port)
	}
	return ":" + port
}
