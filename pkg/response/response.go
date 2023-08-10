package response

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/dewzzjr/ais/pkg/errs"
)

func Error(w http.ResponseWriter, err error) {
	var e errs.Err
	if !errors.As(err, &e) {
		e = errs.Wrap(http.StatusInternalServerError, err)
	}
	Send(w, e.Code, e)
	log.Printf("Status: %s\tError: %s\n", e.Status, e.Message)
}

func Send(w http.ResponseWriter, header int, obj interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(header)
	err := json.NewEncoder(w).Encode(obj)
	if err != nil {
		log.Println(err)
	}
}

func Success(w http.ResponseWriter, header int) {
	Send(w, header, map[string]interface{}{
		"status":  http.StatusText(header),
		"message": "success",
	})
}
