package models

import (
	"time"
)

type Note struct {
	Id        int64     `json:"id,omitempty"`
	Author    string    `json:"author,omitempty"`
	Title     string    `json:"title,omitempty"`
	Desc      string    `json:"desc,omitempty"`
	Tags      string    `json:"tags,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
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
