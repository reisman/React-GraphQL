package data

import (
	"time"
)

type Author struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Message struct {
	ID          string    `json:"id"`
	Text        string    `json:"text"`
	PublishedAt time.Time `json:"publishedat"`
	Author      *Author   `json:"author"`
}
