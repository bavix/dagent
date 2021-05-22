package store

type Store interface {
	Set(key string, val string)
	ReadAll() []string
}
