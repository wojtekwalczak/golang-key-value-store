package store

import "errors"

type StringStore struct {
	store map[string]string
}

// var store = make(map[string]string)

var ErrorNoSuchKey = errors.New("no such key")

func (ss StringStore) Put(key string, value string) error {
	ss.store[key] = value

	return nil
}

func (ss StringStore) Get(key string) (string, error) {
	value, ok := ss.store[key]

	if !ok {
		return "", ErrorNoSuchKey
	}

	return value, nil
}

func (ss StringStore) Delete(key string) error {
	delete(ss.store, key)

	return nil
}
