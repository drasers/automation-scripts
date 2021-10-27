package storage

import (
	"time"
)

type Service interface {
	Save(string, time.Time) (string, error)
	Load(string) (string, error)
	LoadInfo(string) (*Item, error)
	Close() error
}

type Item struct {
	Id      uint64 `json:"id" redis:"id"`
	URL     uint64 `json:"url" redis:"url"`
	Expires uint64 `json:"expires" redis:"expires"`
	Visits  uint64 `json:"visits" redis:"visits"`
}
