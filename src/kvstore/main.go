package main

import (
	"kvstore/service"
	"kvstore/store"
)

func main() {
	var kvstore *store.KVStore = new(store.KVStore)
	service.RunKeyValueStoreService(kvstore)
}
