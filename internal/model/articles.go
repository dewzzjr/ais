package model

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/dewzzjr/ais/pkg/errs"
	"gorm.io/gorm"
)

type Article struct {
	Author string `json:"author"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	gorm.Model
}

type Filter struct {
	Query  *string `schema:"query" json:"-"`
	Author *string `schema:"author" json:"-"`
}

func (m Article) Validate() error {
	if m.Author == "" {
		return errs.Wrap(http.StatusBadRequest, errors.New("author can't be empty"))
	}

	if m.Body == "" {
		return errs.Wrap(http.StatusBadRequest, errors.New("body can't be empty"))
	}

	if m.Title == "" {
		return errs.Wrap(http.StatusBadRequest, errors.New("title can't be empty"))
	}

	return nil
}

func (m *Article) MarshalBinary() ([]byte, error) {
	return json.Marshal(m)
}

func (p *Article) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, p)
}
