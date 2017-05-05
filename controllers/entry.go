package controllers

import (
	"encoding/json"
	"net/http"
)

// Bucket is the name of the bucket storing all the entries
const (
	Bucket = "entries"
	Type   = "entry"
)

// Entry is the main struct
type Entry struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Markdown string `json:"markdown"`
}

// Validate validates that all the required files are not empty.
func (e Entry) Validate() Errors {
	var errors Errors
	if e.Title == "" {
		errors = append(errors, Error{
			Status: http.StatusBadRequest,
			Title:  "title field is required",
		})
	}
	if e.Markdown == "" {
		errors = append(errors, Error{
			Status: http.StatusBadRequest,
			Title:  "markdown field is required",
		})
	}
	return errors
}

// Data contains the Type of the request and the Attributes
type Data struct {
	Type       string `json:"type,omitempty"`
	Attributes *Entry `json:"attributes,omitempty"`
	Links      *Links `json:"links,omitempty"`
}

// Links represent a list of links
type Links map[string]string

// Wrapper is the HATEOAS wrapper
type Wrapper struct {
	Data   *Data   `json:"data,omitempty"`
	Errors *Errors `json:"errors,omitempty"`
}

func (e Entry) Encode() ([]byte, error) {
	return json.Marshal(e)
}

// Decode loads an Entry from json
func (e *Entry) Decode(data []byte) error {
	return json.Unmarshal(data, e)
}
