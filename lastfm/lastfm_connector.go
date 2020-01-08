package lastfm

import (
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//CreateServer create local server for authentication API
func CreateServer() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/lastfm", LastfmAuthenticator)
	log.Fatal(http.ListenAndServe(":8080", router))
}

// Index manage root endpoint
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

// LastfmAuthenticator manage authentication from last.fm api
func LastfmAuthenticator(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "LastfmAuthenticator, %q", html.EscapeString(r.URL.Path))
}
