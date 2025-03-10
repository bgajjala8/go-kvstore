package set

func Set(cache map[string]string, key, value string) string {
	cache[key] = value

	return cache[key]
}
