package store

import (
	"net/http"

	"github.com/brnskn/kv-memory/pkg/response"
	"github.com/gorilla/mux"
)

var repo Repository

func RegisterHandlers(router *mux.Router) {
	repo = NewRepository()
	router.HandleFunc("/", Get).Methods(http.MethodGet)
	router.HandleFunc("/", Set).Methods(http.MethodPost)
	router.HandleFunc("/", Flush).Methods(http.MethodDelete)
}

func Get(w http.ResponseWriter, r *http.Request) {
	store, err := repo.Get(r.FormValue("key"))
	if err != nil {
		response.JsonError(w, err)
		return
	}
	response.Json(w, store)
}

func Set(w http.ResponseWriter, r *http.Request) {
	response.Json(w, repo.Set(r.FormValue("key"), r.FormValue("value")))
}

func Flush(w http.ResponseWriter, r *http.Request) {
	repo.Flush()
	response.JsonSuccess(w, "Store successfully flushed.")
}
