package delete

import (
	"errors"
)

var ErrKeyNotFound = errors.New("key not found for deletion")

func Delete(cache map[string]string, key string) error {
	_, ok := cache[key]
	if !ok {
		return ErrKeyNotFound
	}

	delete(cache, key)
	return nil
}
