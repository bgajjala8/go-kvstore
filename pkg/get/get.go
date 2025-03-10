package get

import (
	"errors"
)

var ErrKeyNotFound = errors.New("key not found")

func Get(cache map[string]string, key string) (string, error) {
	value, ok := cache[key]
	if !ok {
		return "", ErrKeyNotFound
	}

	return value, nil
}
