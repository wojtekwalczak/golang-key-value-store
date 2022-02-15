package main

import (
	"kvstore/service"
	"kvstore/store"
)

func main() {
	service.RunKeyValueStoreService(store.GetKvStore())
}
