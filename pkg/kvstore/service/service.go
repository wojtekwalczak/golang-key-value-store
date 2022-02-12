package service

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type KVStore interface {
	Put(string, string) error
	Get(string) (string, error)
	Delete(string) error
}

func VersionHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("KV Store, v0.0.1"))
}

func KeyHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("key" + r.Method))
}

func KeyValueHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("key - value"))
}

func RunKeyValueStoreService(store KVStore) {
	r := mux.NewRouter()

	r.HandleFunc("/", VersionHandler)
	r.HandleFunc("/{key}", KeyHandler).Methods("GET", "DELETE").Schemes("http")
	r.HandleFunc("/{key}/{value}", KeyValueHandler).Methods("PUT").Schemes("http")

	log.Fatal(http.ListenAndServe(":8080", r))
}
