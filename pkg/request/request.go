package request

import (
	"encoding/json"
	"net/http"

	"github.com/dewzzjr/ais/pkg/collection"
	"github.com/gorilla/schema"
)

type Options string

const (
	JSON     Options = "application/json"
	Query    Options = "query"
	PostForm Options = "application/x-www-form-urlencoded"
)

func (o Options) Valid() bool {
	switch o {
	case JSON, Query, PostForm:
		return true
	default:
		return false
	}
}

func Read(r *http.Request, obj interface{}, opts ...Options) (err error) {
	opts = collection.AppendUnique(opts, Options(r.Header.Get("content-type")))
	var usedBody bool
	for _, opt := range opts {
		if !opt.Valid() {
			continue
		}
		switch opt {
		case Query:
			err = schema.NewDecoder().Decode(obj, r.URL.Query())
			if err != nil {
				return
			}
		case JSON:
			if usedBody {
				continue
			}
			err = json.NewDecoder(r.Body).Decode(obj)
			if err != nil {
				return
			}
			r.Body.Close()
			usedBody = true
		case PostForm:
			if usedBody {
				continue
			}
			err = r.ParseForm()
			if err != nil {
				return
			}
			err = schema.NewDecoder().Decode(obj, r.PostForm)
			if err != nil {
				return
			}
			r.Body.Close()
			usedBody = true
		// Add here to support another options
		default:
		}
	}
	return nil
}
