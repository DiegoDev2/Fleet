package repository

import "time"

type RepoInfo struct {
	Name     string    `json:"name"`
	URL      string    `json:"url"`
	Type     string    `json:"type"`
	Priority int       `json:"priority"`
	LastSync time.Time `json:"last_sync"`
}
