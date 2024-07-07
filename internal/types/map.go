package types

func HasKey[T any](key string, m map[string]T) bool {
	_, ok := m[key]
	return ok
}
