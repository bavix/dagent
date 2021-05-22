package store

import "time"

type Store interface {
	Set(key string, val string, duration *time.Duration)
	ReadAll() []string
}
