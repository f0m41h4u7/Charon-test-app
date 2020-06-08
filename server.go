package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

func Index(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprintf(w, "Hello World!")
}

var route = Route{
	"Index",
	"GET",
	"/",
	Index,
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	handler := route.HandlerFunc

	router.
		Methods(route.Method).
		Path(route.Pattern).
		Name(route.Name).
		Handler(handler)

	return router
}

func main() {
	log.Printf("Server started")
	router := NewRouter()

	log.Fatal(http.ListenAndServe(":1984", router))
}
