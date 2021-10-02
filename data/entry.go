package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

var (
	TimeLayout = time.RFC3339
	TimeZone   = time.Local
)

type Entry struct {
	Metadata
	Content string
}

func ReadEntry(r io.Reader) (*Entry, error) {
	return nil, nil
}

func WriteEntry(w io.Writer, entry *Entry) error {
	md, err := json.MarshalIndent(entry.Metadata, "", "  ")
	if err != nil {
		return err
	}
	_, err = fmt.Fprintf(w, "# %s\n\n```json\n%s\n```\n\n%s", entry.Title, md, entry.Content)
	return err
}

type Metadata struct {
	Title    string   `json:"title,omitempty"`
	Date     string   `json:"date,omitempty"`
	Filename string   `json:"filename,omitempty"`
	Host     string   `json:"host,omitempty"`
	Tags     []string `json:"tags,omitempty"`
}

func (md *Metadata) SetDate(date time.Time) {
	md.Date = date.In(TimeZone).Format(TimeLayout)
}

func (md *Metadata) GetDate() (time.Time, error) {
	return time.Parse(TimeLayout, md.Date)
}

func (md *Metadata) GetWeekDay() string {
	date, _ := md.GetDate()
	return date.Weekday().String()[0:2]
}
