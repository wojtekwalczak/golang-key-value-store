package service

import (
	"io"
	"kvstore/store"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func VersionHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("KV Store, v0.0.1"))
}

func GetKeyHandler(store *store.KVStore) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		key := vars["key"]

		value, err := store.Get(key)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write([]byte(value))
	}
}

func PutKeyValueHandler(store *store.KVStore) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		key := vars["key"]

		value, err := io.ReadAll(r.Body)
		defer r.Body.Close()

		if err != nil {
			http.Error(
				w,
				err.Error(),
				http.StatusInternalServerError,
			)
			return
		}

		err = (*store).Put(key, string(value))

		if err != nil {
			http.Error(
				w,
				err.Error(),
				http.StatusInternalServerError,
			)
		}

		w.WriteHeader(http.StatusCreated)
	}
}

func RunKeyValueStoreService(kvstore *store.KVStore) {
	r := mux.NewRouter()

	r.HandleFunc("/", VersionHandler)
	r.HandleFunc("/{key}", GetKeyHandler(kvstore)).Methods("GET").Schemes("http")
	r.HandleFunc("/{key}", PutKeyValueHandler(kvstore)).Methods("PUT").Schemes("http")

	log.Fatal(http.ListenAndServe(":8080", r))
}
