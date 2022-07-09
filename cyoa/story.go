package cyoa

import (
	"encoding/json"
	"io"
)

type Story map[string]Chapter

type Chapter struct {
	Title      string    `json:"title"`
	Paragraphs []string  `json:"story"`
	Options    []Options `json:"options"`
}

type Options struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}

func ParseJSON(i io.Reader) (Story, error) {
	d := json.NewDecoder(i)

	var stories Story
	err := d.Decode(&stories)

	if err != nil {
		return nil, err
	}

	return stories, nil
}
