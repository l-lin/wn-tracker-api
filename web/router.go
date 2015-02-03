package web

import (
	"github.com/gorilla/mux"
	oauth2 "github.com/goincremental/negroni-oauth2"
	"net/http"
	"log"
	"encoding/json"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.Headers("Content-Type", "application/json", "X-Requested-With", "XMLHttpRequest")

	for _, route := range routes {
		router.
		Methods(route.Method).
		Path(route.Pattern).
		Name(route.Name).
		Handler(wrapWithCheckAuth(&route.HandlerFunc))
	}

	return router
}

func NewSignInRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.Headers("Content-Type", "application/json", "X-Requested-With", "XMLHttpRequest")
	route := Route{
		"SignIn",
		"GET",
		"/signin",
		SignIn,
	}
	router.Methods(route.Method).
	Path(route.Pattern).
	Name(route.Name).
	Handler(route.HandlerFunc)
	return router
}

func wrapWithCheckAuth(handlerFunc *http.HandlerFunc) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		token := oauth2.GetToken(r)
		if token == nil || !token.Valid() {
			w.WriteHeader(http.StatusUnauthorized)
			if err := json.NewEncoder(w).Encode(JsonErr{Code: http.StatusUnauthorized, Text: "You are not authenticated!"}); err != nil {
				log.Fatalf("[x] Error when encoding the json. Reason: %s", err.Error())
			}
		} else {
			handlerFunc.ServeHTTP(w, r)
		}
	}
}
