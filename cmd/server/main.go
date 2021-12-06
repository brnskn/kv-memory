package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/brnskn/kv-memory/internal/store"
	"github.com/brnskn/kv-memory/pkg/config"
	s "github.com/brnskn/kv-memory/pkg/store"
	"github.com/gorilla/mux"
)

func main() {
	config.Load()
	interval := time.Minute *
		time.Duration(config.GetInt("AUTO_SAVE_INTERVAL", 5))
	s.Instance().StartAutoSaver(interval)

	router := mux.NewRouter()
	store.RegisterHandlers(router)
	addr := fmt.Sprintf("%s:%s", config.Get("BIND_HOST", "0.0.0.0"),
		config.Get("PORT", "8080"))
	log.Fatal(http.ListenAndServe(addr, router))
}
