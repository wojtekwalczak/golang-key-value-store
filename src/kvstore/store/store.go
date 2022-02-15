package store

import (
	"errors"
)

type KVStore struct {
	store map[string]string
}

var kvstore *KVStore = &KVStore{
	store: make(map[string]string),
}

func GetKvStore() *KVStore {
	return kvstore
}

var ErrorNoSuchKey = errors.New("no such key")

func (kvs *KVStore) Put(key string, value string) error {
	kvs.store[key] = value
	return nil
}

func (kvs *KVStore) Get(key string) (string, error) {
	value, ok := kvs.store[key]

	if !ok {
		return "", ErrorNoSuchKey
	}

	return value, nil
}

func (kvs *KVStore) Delete(key string) error {
	delete(kvs.store, key)

	return nil
}
