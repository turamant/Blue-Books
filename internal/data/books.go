package data

import (
	"time"
)

type Book struct {
	ID int64 `json:"id"`
	CreatedAt time.Time `json:"-"`
	Title string `json:"title"`
	Year int32 `json:"year,omitempty,string"`
	Pages Pages `json:"pages,omitempty"`
	Genres []string `json:"genres,omitempty"`
	Language string `json:"language"`
}