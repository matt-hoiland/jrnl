package data

import "time"

type Entry struct {
	MD      Metadata
	Content string
}

type Metadata struct {
	Title    string    `json:"title,omitempty"`
	Filename string    `json:"filename,omitempty"`
	Date     time.Time `json:"date,omitempty"`
	Host     string    `json:"host,omitempty"`
	Tags     []string  `json:"tags,omitempty"`
}
