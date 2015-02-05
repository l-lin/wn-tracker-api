package web

import (
	"github.com/gorilla/mux"
	"github.com/l-lin/wn-tracker-api/token"
	oauth2 "github.com/goincremental/negroni-oauth2"
	sessions "github.com/goincremental/negroni-sessions"
	"net/http"
	"log"
	"encoding/json"
	"fmt"
)

// Returns the routers for novels, feeds and notifications
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.Headers("Content-Type", "application/json", "X-Requested-With", "XMLHttpRequest")

	for _, route := range routes {
		router.
		Methods(route.Method).
		Path(route.Pattern).
		Name(route.Name).
		Handler(wrapWithCheckAuth(route.HandlerFunc))
	}

	return router
}

// Returns a router for signing in Google account
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

// Wrap the HandlerFunc by checking if the user is indeed authenticated
func wrapWithCheckAuth(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		oauthT := oauth2.GetToken(r)
		if oauthT == nil {
			w.WriteHeader(http.StatusUnauthorized)
			if err := json.NewEncoder(w).Encode(JsonErr{Code: http.StatusUnauthorized, Text: "You are not authenticated!"}); err != nil {
				log.Fatalf("[x] Error when encoding the json. Reason: %s", err.Error())
			}
		} else {
			if !oauthT.Valid() {
				s := sessions.GetSession(r)
				userId := s.Get(SESSION_USER_ID)
				t := token.Get(fmt.Sprintf("%v", userId))
				log.Printf("[-] Refreshing the token %s", t.RefreshToken)
				if t.Refresh() {
					handlerFunc.ServeHTTP(w, r)
				}
			} else {
				handlerFunc.ServeHTTP(w, r)
			}
		}
	}
}
