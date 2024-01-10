package models

import (
	"time"
)

type Note struct {
	Id        string    `json:"id,omitempty"`
	Author    string    `json:"author"`
	Title     string    `json:"title"`
	Desc      string    `json:"desc"`
	Tags      string    `json:"tags"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewNote(author, title, desc, tags string) *Note {
	return &Note{
		Author:    author,
		Title:     title,
		Desc:      desc,
		Tags:      tags,
		CreatedAt: time.Now().UTC(),
	}
}
