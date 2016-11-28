package main

import (
	"net/http"
	"sync"
	"encoding/json"
	"strings"
	"flag"
	"log"
	"fmt"

	"github.com/satori/go.uuid"
)

var (
	port int
)

func handler() http.Handler {
	var lock sync.RWMutex
	sessions := make(map[string]struct{})
	mux := http.NewServeMux()
	mux.HandleFunc("/session", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		u := uuid.NewV4().String()
		lock.Lock()
		sessions[u] = struct{}{}
		lock.Unlock()
		json.NewEncoder(w).Encode(struct {
			S string `json:"sessionId"`
		}{u})
	})
	mux.HandleFunc("/session/", func(w http.ResponseWriter, r *http.Request) {
		u := strings.Split(r.URL.Path, "/")[2]
		lock.RLock()
		_, ok := sessions[u]
		lock.RUnlock()
		if !ok {
			http.Error(w, "Session not found", http.StatusNotFound)
			return
		}
		if r.Method != http.MethodDelete {
			return
		}
		lock.Lock()
		delete(sessions, u)
		lock.Unlock()
	})
	return mux
}

func init() {
	flag.IntVar(&port, "port", 4444, "port to bind to")
	flag.Parse()
}

func main() {
	listen := fmt.Sprintf(":%d", port)
	log.Println("listening on", listen)
	log.Print(http.ListenAndServe(listen, handler()))
}