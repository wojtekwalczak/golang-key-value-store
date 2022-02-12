package main

import (
	"kvstore/service"
	"kvstore/store"
)

func main() {
	store := store.StringStore{}
	service.RunKeyValueStoreService(store)
}
