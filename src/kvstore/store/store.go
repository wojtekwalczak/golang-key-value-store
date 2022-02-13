package store

import (
	"errors"
)

type KVStore struct{}

var store map[string]string = make(map[string]string)

var ErrorNoSuchKey = errors.New("no such key")

func (kvs KVStore) Put(key string, value string) error {
	store[key] = value
	return nil
}

func (kvs KVStore) Get(key string) (string, error) {
	value, ok := store[key]

	if !ok {
		return "", ErrorNoSuchKey
	}

	return value, nil
}

func (kvs KVStore) Delete(key string) error {
	delete(store, key)

	return nil
}
