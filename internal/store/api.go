package store

import (
	"net/http"

	"github.com/brnskn/kv-memory/pkg/response"
	"github.com/gorilla/mux"
)

// A variable that typed Repository interface
var repo Repository

// Inits new repository to use on all handlers
func Init() {
	repo = NewRepository()
}

// Inits handlers to router
func RegisterHandlers(router *mux.Router) {
	Init()
	router.HandleFunc("/", Get).Methods(http.MethodGet)
	router.HandleFunc("/", Set).Methods(http.MethodPost)
	router.HandleFunc("/", Flush).Methods(http.MethodDelete)
}

// Gets and writes to http response writer
func Get(w http.ResponseWriter, r *http.Request) {
	store, err := repo.Get(r.FormValue("key"))
	if err != nil {
		response.JsonError(w, err)
		return
	}
	response.Json(w, store)
}

// Sets and writes result to http response writer
func Set(w http.ResponseWriter, r *http.Request) {
	response.Json(w, repo.Set(r.FormValue("key"), r.FormValue("value")))
}

// Flushes and writes result to http response writer
func Flush(w http.ResponseWriter, r *http.Request) {
	repo.Flush()
	response.JsonSuccess(w, "store successfully flushed")
}
